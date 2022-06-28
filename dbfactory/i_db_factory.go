package dbfactory

import "github.com/xm-chentl/go-resource/dbfactory/dbtype"

// IDbFactory 数据库实例工厂
type IDbFactory interface {
	BuildByType(dbtype.Value) (IFactory, error) // 类型与数据库 一对一
	BuildByName(string) (IFactory, error)       // 名称与数据库 一对一
}
