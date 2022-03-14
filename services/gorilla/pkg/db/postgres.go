package db

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"gorilla/pkg/config"
	"gorilla/pkg/entities"
	"gorilla/pkg/errors"
	"time"
)

type PostgresDeltaDB struct {
	db     *sqlx.DB
	logger log.Logger
}

func NewPostgresDeltaDB(cfg config.DBConfig, logger log.Logger) (*PostgresDeltaDB, error) {
	deltaDB, err := sqlx.Connect("postgres", cfg.Target)
	if err != nil {
		return nil, err
	}
	return &PostgresDeltaDB{
		db:     deltaDB,
		logger: logger,
	}, mapDBError(logger, err)
}

const addDeltaQuery = `INSERT INTO deltas (owner_id, object_id, object_type, category, date, balance) 
VALUES (:owner_id, :object_id, :object_type, :category, :date, :balance) 
ON CONFLICT (date, owner_id, object_id, object_type, category) DO
UPDATE SET balance = deltas.balance + :balance;`

func (p *PostgresDeltaDB) AddDeltas(ctx context.Context, deltas []entities.Delta) (err error) {
	var (
		query string
		args  []interface{}
	)
	for i := 0; i < len(deltas); i++ {
		query, args, err = p.db.BindNamed(addDeltaQuery, map[string]interface{}{
			"owner_id": deltas[i].OwnerID,
			"object_id": deltas[i].ObjectID,
			"object_type": deltas[i].ObjectType,
			"category": deltas[i].Category,
			"date":     deltas[i].Date,
			"balance":  deltas[i].Balance,
		})
		if err != nil {
			err = mapDBError(p.logger, err)
			return
		}
		_, err = p.db.ExecContext(ctx, query, args...)
		if err != nil {
			err = mapDBError(p.logger, err)
			return
		}
	}
	return nil
}

const getModelDeltasQuery = `SELECT owner_id, object_id, object_type, category, date, balance FROM deltas
WHERE owner_id = :owner_id AND object_id = :object_id AND object_type = :object_type AND date >= :first_date AND date <= :last_date;`

const getDeltasQuery = `SELECT owner_id, object_id, object_type, category, DATE(date) as date, sum(balance) as balance FROM deltas 
WHERE owner_id = :owner_id AND date >= :first_date AND date <= :last_date
GROUP BY owner_id, object_id, object_type, category, DATE(date);`

func (p *PostgresDeltaDB) GetDeltas(
	ctx context.Context, ownerID uuid.UUID, objectID, objectType string, firstDate, lastDate time.Time, useCategories bool,
) (deltas []entities.Delta, err error) {
	var (
		query    string
		args     []interface{}
		pgDeltas []DeltaResult
	)
	if objectID != "" {
		objectID, err := uuid.FromString(objectID)
		if err != nil {
			return nil, err
		}
		query, args, err = p.db.BindNamed(getModelDeltasQuery, map[string]interface{}{
			"owner_id":   ownerID,
			"object_id":  objectID,
			"object_type": objectType,
			"first_date": firstDate,
			"last_date":  lastDate,
		})
	} else {
		query, args, err = p.db.BindNamed(getDeltasQuery, map[string]interface{}{
			"owner_id":   ownerID,
			"first_date": firstDate,
			"last_date":  lastDate,
		})
	}
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	err = p.db.SelectContext(ctx, &pgDeltas, query, args...)
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	result := make([]entities.Delta, 0, len(pgDeltas))
	for i := 0; i < len(pgDeltas); i++ {
		result = append(result, makeDeltaEntity(pgDeltas[i]))
	}
	return result, nil
}

const getBalanceQuery = `SELECT COALESCE(SUM(balance), 0) as balance
FROM deltas
WHERE owner_id = :owner_id;`

func (p *PostgresDeltaDB) GetBalance(ctx context.Context, ownerID uuid.UUID) (balance float64, err error) {
	var (
		query     string
		args      []interface{}
		pgBalance BalanceResult
	)
	query, args, err = p.db.BindNamed(getBalanceQuery, map[string]interface{}{
		"owner_id": ownerID,
	})
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	err = p.db.GetContext(ctx, &pgBalance, query, args...)
	if err != nil {

		err = mapDBError(p.logger, err)
		return
	}
	return pgBalance.Balance, nil
}

func mapDBError(logger log.Logger, err error) error {
	if err == nil {
		return nil
	}

	pqErr, ok := err.(*pq.Error)
	if !ok {
		logger.Log("msg", "unknown error format", "err", err)
		return errors.ErrInternal
	}

	switch pqErr.Code {
	default:
		logger.Log("msg", "unknown postgres error code", "code", pqErr.Code, "err", err)
		return errors.ErrInternal
	}
}
