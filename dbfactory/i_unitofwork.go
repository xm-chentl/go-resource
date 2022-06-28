package dbfactory

type IUnitOfWork interface {
	Commit() error
}
