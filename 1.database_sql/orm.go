package orm

import (
	"database/sql"
	"orm/log"
	"orm/session"
)

// Engine 管理所数据库的session和transactions
type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// 发送一个ping,确保数据库连接存在
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{db: db}
	log.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

//
// NewSession
//  @Description: 连接数据库，返回 *sql.DB。 调用 db.Ping()，检查数据库是否能够正常连接。
//  @receiver engine
//  @return *session.Session
//
func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
