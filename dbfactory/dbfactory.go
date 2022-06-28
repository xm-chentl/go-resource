package dbfactory

import (
	"fmt"

	"github.com/xm-chentl/go-resource/dbfactory/dbtype"
)

// Configure 配置 todo: 暂时不封装
type Configure struct {
	Alias   string
	Type    dbtype.Value
	Factory IFactory
}

type dbFactory struct {
	dbTypeOfFactory map[dbtype.Value]IFactory
	nameOfFacotry   map[string]IFactory
}

func (f dbFactory) BuildByType(dbType dbtype.Value) (dbFactory IFactory, err error) {
	dbFactory, ok := f.dbTypeOfFactory[dbType]
	if !ok {
		err = fmt.Errorf("dbtype: %s database implementation not configured", dbType.String())
		return
	}

	return
}

func (f dbFactory) BuildByName(name string) (dbFactory IFactory, err error) {
	dbFactory, ok := f.nameOfFacotry[name]
	if !ok {
		err = fmt.Errorf("name: %s database implementation not configured", name)
		return
	}

	return
}

func NewByType(dbTypeConfig map[dbtype.Value]IFactory) IDbFactory {
	return &dbFactory{
		dbTypeOfFactory: dbTypeConfig,
		nameOfFacotry:   make(map[string]IFactory),
	}
}

func NewByName(aliasConfig map[string]IFactory) IDbFactory {
	return &dbFactory{
		dbTypeOfFactory: map[dbtype.Value]IFactory{},
		nameOfFacotry:   aliasConfig,
	}
}
