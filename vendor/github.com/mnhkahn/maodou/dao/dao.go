package dao

import (
	"fmt"
	. "github.com/mnhkahn/maodou/models"
)

type DaoContainer interface {
	AddResult(p *Result)
	AddResults(p []Result)
	DelResult(id interface{})
	DelResults(source string)
	UpdateResult(p *Result)
	AddOrUpdate(p *Result)
	GetResultById(id int) *Result
	GetResultByLink(url string) *Result
	GetResult(author, sort string, limit, start int) []Result
	IsResultUpdate(p *Result) bool
	Search(q string, limit, start int) (int, float64, []Result)
	Debug(is_debug bool)
}

type Dao interface {
	NewDaoImpl(dsn string) (DaoContainer, error)
}

var daos = make(map[string]Dao)

// Register makes a config adapter available by the adapter name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, dao Dao) {
	if dao == nil {
		panic("dao: Register dao is nil")
	}
	if _, ok := daos[name]; ok {
		panic("dao: Register called twice for adapter " + name)
	}
	daos[name] = dao
}

func NewDao(dao_name, dsn string) (DaoContainer, error) {
	dao, ok := daos[dao_name]
	if !ok {
		return nil, fmt.Errorf("parser: unknown dao_name %q", dao_name)
	}
	return dao.NewDaoImpl(dsn)
}
