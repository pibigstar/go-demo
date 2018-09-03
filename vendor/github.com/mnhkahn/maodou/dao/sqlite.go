package dao

type SqliteDao struct {
}

func (this *SqliteDao) NewDaoImpl(dsn string) (DaoContainer, error) {
	d := new(DuoShuoDaoContainer)
	return d, nil
}

type SqliteDaoContainer struct {
	is_debug bool
}

func init() {
	Register("sqlite", &SqliteDao{})
}
