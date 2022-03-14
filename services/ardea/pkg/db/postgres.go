package db

import (
	"context"
	"github.com/mattn/go-nulltype"

	"ardea/pkg/errors"

	"github.com/go-kit/kit/log"
	"github.com/satori/go.uuid"
	"golang.org/x/xerrors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"ardea/pkg/config"
	"ardea/pkg/entities"
)

type PostgresModelDB struct {
	db     *sqlx.DB
	logger log.Logger
}

func NewPostgresModelDB(cfg config.DBConfig, logger log.Logger) (*PostgresModelDB, error) {
	modelDB, err := sqlx.Connect("postgres", cfg.Target)
	if err != nil {
		return nil, err
	}
	return &PostgresModelDB{
		db:     modelDB,
		logger: logger,
	}, mapDBError(logger, err)
}

func (PostgresModelDB) fieldName(t modelUpdateType) (res string, err error) {
	switch t {
	case modelUpdateTypePath:
		res = "path"
	case modelUpdateTypeState:
		res = "state"
	case modelUpdateTypeErrStr:
		res = "err"
	default:
		err = xerrors.New("invalid model update type")
	}
	return
}

const createModelQuery = `INSERT
	INTO models (id, owner_id, name, state, input_shape, output_shape, value_type)
	VALUES (:id, :owner_id, :name, :state, :input_shape, :output_shape, :value_type)
	RETURNING id, owner_id, name, state, input_shape, output_shape, path, err, value_type`

func (p *PostgresModelDB) CreateModel(
	ctx context.Context, modelID, ownerID uuid.UUID, name string,
	inputShape, outputShape [][]nulltype.NullInt64, valueType entities.ValueType,
) (model entities.Model, err error) {
	var (
		query   string
		args    []interface{}
		pgModel ModelResult
	)
	query, args, err = p.db.BindNamed(createModelQuery, map[string]interface{}{
		"id":           modelID,
		"owner_id":     ownerID,
		"name":         name,
		"state":        serializeModelState(entities.ModelStateInit),
		"input_shape":  ioShape(inputShape),
		"output_shape": ioShape(outputShape),
		"value_type":   serializeValueType(valueType),
	})
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}

	err = p.db.GetContext(ctx, &pgModel, query, args...)
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}

	return makeModelEntity(pgModel), nil
}

const getModelQuery = `SELECT id, owner_id, name, state, input_shape, output_shape, path, err, value_type
	    	 	 	   FROM models WHERE id = :id`

func (p *PostgresModelDB) GetModel(ctx context.Context, modelID uuid.UUID) (model entities.Model, err error) {
	var (
		query   string
		args    []interface{}
		pgModel ModelResult
	)

	query, args, err = p.db.BindNamed(getModelQuery, map[string]interface{}{
		"id": modelID,
	})
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	err = p.db.GetContext(ctx, &pgModel, query, args...)
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}

	if pgModel.State == "DELETED" {
		err = errors.ErrModelDeleted
		return
	}

	return makeModelEntity(pgModel), nil
}

const getModelByNameQuery = `SELECT id, owner_id, name, state, input_shape, output_shape, path, err, value_type
	FROM models
	WHERE owner_id = :owner_id AND name = :name AND state != :state`

func (p *PostgresModelDB) GetModelByName(
	ctx context.Context, ownerID uuid.UUID, name string,
) (model entities.Model, err error) {
	var (
		query   string
		args    []interface{}
		pgModel ModelResult
	)

	query, args, err = p.db.BindNamed(getModelByNameQuery, map[string]interface{}{
		"owner_id": ownerID,
		"name":     name,
		"state":    serializeModelState(entities.ModelStateDeleted),
	})
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}

	err = p.db.GetContext(ctx, &pgModel, query, args...)
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	return makeModelEntity(pgModel), nil
}

const updateModelQuery = `UPDATE models
	SET 
		state = coalesce(:state, state),
		path = coalesce(:path, path),
		err = coalesce(:err, err)
	WHERE id = :id
	RETURNING id, owner_id, name, state, input_shape, output_shape, path, err, value_type`

func (p *PostgresModelDB) UpdateModel(
	ctx context.Context, modelID uuid.UUID, params ...ModelUpdateParam,
) (model entities.Model, err error) {
	var (
		query   string
		args    []interface{}
		pgModel ModelResult
	)

	paramsMap := map[string]interface{}{
		"id":    modelID,
		"state": nil,
		"path":  nil,
		"err":   nil,
	}
	for _, param := range params {
		var name string
		name, err = p.fieldName(param.field)
		if err != nil {
			err = mapDBError(p.logger, err)
			return
		}
		paramsMap[name] = param.value
	}
	query, args, err = p.db.BindNamed(updateModelQuery, paramsMap)
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	err = p.db.GetContext(ctx, &pgModel, query, args...)
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	return makeModelEntity(pgModel), nil
}

const getModelsByOwnerQuery = `SELECT id, owner_id, name, state, input_shape, output_shape, path, err, value_type
	FROM models
	WHERE owner_id = :owner_id AND state != :state`

func (p *PostgresModelDB) ListModels(ctx context.Context, ownerID uuid.UUID) (models []entities.Model, err error) {
	var (
		query    string
		args     []interface{}
		pgModels []ModelResult
	)

	query, args, err = p.db.BindNamed(getModelsByOwnerQuery, map[string]interface{}{
		"owner_id": ownerID,
		"state":    serializeModelState(entities.ModelStateDeleted),
	})
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	err = p.db.SelectContext(ctx, &pgModels, query, args...)
	if err != nil {
		err = mapDBError(p.logger, err)
		return
	}
	result := make([]entities.Model, 0, len(pgModels))
	for i := 0; i < len(pgModels); i++ {
		result = append(result, makeModelEntity(pgModels[i]))
	}
	return result, nil
}
