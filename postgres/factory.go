package postgres

import (
	"context"

	"github.com/xm-chentl/go-resource/dbfactory"

	"github.com/jackc/pgx/v4/pgxpool"
)

type factory struct {
	pgxPool *pgxpool.Pool
	connStr string
}

// Todo: 联合事务有问题，重新
func (f factory) Db(args ...interface{}) dbfactory.IRepository {
	repo := &repository{
		pool: &pool{
			pgxPool: f.pgxPool,
		},
	}
	for index := range args {
		if ctx, ok := args[index].(context.Context); ok {
			repo.ctx = ctx
			continue
		}
		if uow, ok := args[index].(*unitOfWork); ok {
			repo.uow = uow
			continue
		}
		if uow, ok := args[index].(dbfactory.IUnitOfWork); ok {
			repo.uow = &unitOfWork{
				pool:          repo.pool,
				addOfQueue:    make([]commitQueueInfo, 0),
				deleteOfQueue: make([]commitQueueInfo, 0),
				updateOfQueue: make([]commitQueueInfo, 0),
			}
			repo.repositoryBase = dbfactory.NewRepository(uow)
		}
	}
	if repo.ctx == nil {
		repo.ctx = context.Background()
	}
	repo.pool.ctx = repo.ctx
	if repo.uow != nil {
		repo.uow.ctx = repo.ctx
	}

	return repo
}

func (f factory) Uow() dbfactory.IUnitOfWork {
	return &unitOfWork{
		pool: &pool{
			pgxPool: f.pgxPool,
		},
		addOfQueue:    make([]commitQueueInfo, 0),
		updateOfQueue: make([]commitQueueInfo, 0),
	}
}

func New(connStr string) dbfactory.IFactory {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		panic("connect to database config faild: " + err.Error())
	}

	ctx := context.Background()
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		panic("Unable to connect to database: " + connStr)
	}
	if err = pool.Ping(ctx); err != nil {
		panic("connect to database faild: " + err.Error())
	}

	return &factory{
		pgxPool: pool,
		connStr: connStr,
	}
}

func NewByGorm(connStr string) dbfactory.IFactory {

	return nil
}
