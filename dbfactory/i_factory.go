package dbfactory

// IFactory 数据库实现
type IFactory interface {
	Db(...interface{}) IRepository
	Uow() IUnitOfWork
}
