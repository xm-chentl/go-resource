package dbfactory

import "github.com/xm-chentl/go-resource/dbfactory/dbtype"

type RepositoryBase struct {
	uow *unitOfWork
}

func (r RepositoryBase) SetUow(dbType dbtype.Value, uow IUnitOfWork) {
	r.uow.uowMap[dbType] = uow
}

func NewRepository(uow IUnitOfWork) *RepositoryBase {
	return &RepositoryBase{
		uow: uow.(*unitOfWork),
	}
}
