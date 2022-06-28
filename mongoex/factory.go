package mongoex

import (
	"context"

	"github.com/xm-chentl/go-resource/dbfactory"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type factory struct {
	dbName   string
	database *mongo.Database
}

func (f factory) Db(args ...interface{}) dbfactory.IRepository {
	repo := &repository{
		database: f.database,
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
			repo.uow = newUnitOfWork(f.database)
			repo.repositoryBase = dbfactory.NewRepository(uow)
		}
	}
	if repo.ctx == nil {
		repo.ctx = context.Background()
	}
	if repo.uow != nil {
		repo.uow.ctx = repo.ctx
	}

	return repo
}

func (f factory) Uow() dbfactory.IUnitOfWork {
	return nil
}

func New(dbName, connStr string) dbfactory.IFactory {
	opt := options.Client().ApplyURI(connStr)
	client, err := mongo.NewClient(opt)
	if err != nil {
		panic("create connect to mongo faild err: " + err.Error())
	}
	if err = client.Connect(context.Background()); err != nil {
		panic("connect to mongo faild err: " + err.Error())
	}

	return &factory{
		dbName:   dbName,
		database: client.Database(dbName),
	}
}
