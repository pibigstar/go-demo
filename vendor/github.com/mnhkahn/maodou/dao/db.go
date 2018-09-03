package dao

// import (
// 	. "cyeam_Result/models"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/go-xorm/xorm"
// )

// type DbDao struct {
// 	Dsn string
// }

// func (this *DbDao) NewDaoImpl(dsn string) (DaoContainer, error) {
// 	db := new(DbDaoContainer)
// 	var err error
// 	db.Engine, err = xorm.NewEngine("mysql", dsn)
// 	if err != nil {
// 		return db, err
// 	}
// 	db.Engine.SetMaxConns(5)
// 	return db, nil
// }

// type DbDaoContainer struct {
// 	Engine *xorm.Engine
// }

// func (this *DbDaoContainer) Debug(is_debug bool) {
// 	this.Engine.ShowDebug = is_debug
// }

// func (this *DbDaoContainer) AddResult(p *Result) {
// 	this.Engine.Table("Result").Insert(p)
// }

// func (this *DbDaoContainer) AddResults(p []Result) {

// }

// func (this *DbDaoContainer) DelResult(id int) {

// }

// func (this *DbDaoContainer) DelResults(source string) {

// }

// func (this *DbDaoContainer) UpdateResult(p *Result) {
// 	this.Engine.Table("Result").Where("link=?", p.Link).Update(p)
// }

// func (this *DbDaoContainer) GetResultById(id int) *Result {
// 	p := new(Result)
// 	p.Id = id
// 	this.Engine.Table("Result").Get(p)
// 	return p
// }

// func (this *DbDaoContainer) GetResultByLink(url string) *Result {
// 	p := new(Result)
// 	p.Link = url
// 	_, err := this.Engine.Table("Result").Get(p)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return p
// }

// func (this *DbDaoContainer) GetResult(author, sort string, limit, start int) []Result {
// 	p := make([]Result, 0)
// 	if author == "" {
// 		author = "TRUE"
// 	}
// 	if sort == "" {
// 		sort = "create_time desc"
// 	}
// 	this.Engine.Table("Result").Where("author=?", author).OrderBy(sort).Limit(limit, start).Find(&p)
// 	return p
// }

// func (this *DbDaoContainer) IsResultUpdate(p *Result) bool {
// 	is_update := false
// 	temp := this.GetResultByLink(p.Link)
// 	if temp.Title != "" {
// 		if temp.Title != p.Title || temp.Author != p.Author || temp.Detail != p.Detail || temp.Figure != p.Figure {
// 			is_update = true
// 		}
// 	}
// 	return is_update
// }

// func (this *DbDaoContainer) Search(q string, limit, start int) []Result {
// 	res := make([]Result, 0)
// 	if q == "" {
// 		q = "%"
// 	} else {
// 		q = "%" + q + "%"
// 	}
// 	this.Engine.ShowDebug = true
// 	err := this.Engine.Table("Result").Decr("create_time").Limit(limit, start).Where("title like ? OR detail like ?", q, q).Find(&res)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return res
// }

// func init() {
// 	Register("db", &DbDao{})
// }
