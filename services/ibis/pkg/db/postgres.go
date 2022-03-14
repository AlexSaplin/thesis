package db

import (
	"context"
	"github.com/mattn/go-nulltype"
	"ibis/pkg/errors"

	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"

	"ibis/pkg/config"
	"ibis/pkg/entities"
)

type PostgresFunctionDB struct {
	db     *sqlx.DB
	logger log.Logger
}

func NewPostgresFunctionDB(cfg config.DBConfig, logger log.Logger) (*PostgresFunctionDB, error) {
	FunctionDB, err := sqlx.Connect("postgres", cfg.Target)
	if err != nil {
		return nil, err
	}
	return &PostgresFunctionDB{
		db:     FunctionDB,
		logger: logger,
	}, mapDBError(logger, err)
}



const createFunctionQuery = `INSERT
	INTO Functions (id, owner_id, name, state)
	VALUES (:id, :owner_id, :name, :state)
	RETURNING id, owner_id, name, state, code_path, image, err, meta`

func (p *PostgresFunctionDB) CreateFunction(
	ctx context.Context, functionID, ownerID uuid.UUID, name string,
) (Function entities.Function, err error) {
	var (
		query   string
		args    []interface{}
		pgFunction functionResult
	)
	query, args, err = p.db.BindNamed(createFunctionQuery, map[string]interface{}{
		"id":           functionID,
		"owner_id":     ownerID,
		"name":         name,
		"state":        serializeFunctionState(entities.FunctionStateInit),
	})
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}

	err = p.db.GetContext(ctx, &pgFunction, query, args...)
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}

	return makeFunctionEntity(pgFunction), nil
}

const getFunctionQuery = `SELECT id, owner_id, name, state, code_path, image, err, meta
	    	 	 	   FROM Functions WHERE id = :id`

func (p *PostgresFunctionDB) GetFunction(ctx context.Context, functionID uuid.UUID) (Function entities.Function, err error) {
	var (
		query   string
		args    []interface{}
		pgFunction functionResult
	)

	query, args, err = p.db.BindNamed(getFunctionQuery, map[string]interface{}{
		"id": functionID,
	})
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	err = p.db.GetContext(ctx, &pgFunction, query, args...)
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}

	if pgFunction.State == "DELETED" {
		err = errors.ErrFunctionDeleted
		return
	}

	return makeFunctionEntity(pgFunction), nil
}

const getFunctionByNameQuery = `SELECT id, owner_id, name, state, code_path, image, err, meta
	FROM functions
	WHERE owner_id = :owner_id AND name = :name AND state != :state`

func (p *PostgresFunctionDB) GetFunctionByName(
	ctx context.Context, ownerID uuid.UUID, name string,
) (Function entities.Function, err error) {
	var (
		query   string
		args    []interface{}
		pgFunction functionResult
	)

	query, args, err = p.db.BindNamed(getFunctionByNameQuery, map[string]interface{}{
		"owner_id": ownerID,
		"name":     name,
		"state":    serializeFunctionState(entities.FunctionStateDeleted),
	})
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}

	err = p.db.GetContext(ctx, &pgFunction, query, args...)
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	return makeFunctionEntity(pgFunction), nil
}

const updateFunctionQuery = `UPDATE Functions
	SET 
		state = coalesce(:state, state),
		err = coalesce(:err, err),
		image = coalesce(:image, image),
		code_path = coalesce(:code_path, code_path),
		meta = coalesce(:meta, meta),
		name = coalesce(:name, name)
	WHERE id = :id
	RETURNING id, owner_id, name, state, code_path, image, err, meta`

func (p *PostgresFunctionDB) UpdateFunction(
	ctx context.Context, functionID uuid.UUID, param entities.UpdateFunctionParam,
) (Function entities.Function, err error) {
	var (
		query   string
		args    []interface{}
		pgFunction functionResult
	)
	var state nulltype.NullString
	if param.State != nil {
		state = nulltype.NullStringOf(string(serializeFunctionState(*param.State)))
	}

	var deleteName nulltype.NullString
	if param.State != nil && *param.State == entities.FunctionStateDeleted {
		deleteName.Set(functionID.String())
	}

	paramsMap := map[string]interface{}{
		"id":    functionID,
		"state":     state,
		"err":       param.ErrStr,
		"image":     param.ImageURL,
		"code_path": param.CodePath,
		"meta":      param.Metadata,
		"name":      deleteName,
	}

	query, args, err = p.db.BindNamed(updateFunctionQuery, paramsMap)
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	err = p.db.GetContext(ctx, &pgFunction, query, args...)
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	return makeFunctionEntity(pgFunction), nil
}

const getFunctionsByOwnerQuery = `SELECT id, owner_id, name, state, code_path, image, err, meta
	FROM functions
	WHERE owner_id = :owner_id AND state != :state`

func (p *PostgresFunctionDB) ListFunctions(ctx context.Context, ownerID uuid.UUID) (Functions []entities.Function, err error) {
	var (
		query    string
		args     []interface{}
		pgFunctions []functionResult
	)

	query, args, err = p.db.BindNamed(getFunctionsByOwnerQuery, map[string]interface{}{
		"owner_id": ownerID,
		"state":    serializeFunctionState(entities.FunctionStateDeleted),
	})
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	err = p.db.SelectContext(ctx, &pgFunctions, query, args...)
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	result := make([]entities.Function, 0, len(pgFunctions))
	for i := 0; i < len(pgFunctions); i++ {
		result = append(result, makeFunctionEntity(pgFunctions[i]))
	}
	return result, nil
}
