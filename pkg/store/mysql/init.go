package mysql

import (
	"database/sql"
	"fmt"
	"github.com/dembygenesis/blog_exam/pkg/config"
	"time"
)

type Article struct {
	conn *sql.DB
}

const DbConnectTimeoutSecs int = 15
const DbExecTimeoutSecs int = 15

// NewStore returns a new Article MYSQL store instance
func NewStore(cfg *config.Database) (*Article, error) {
	// Setup connection
	if cfg == nil {
		return nil, fmt.Errorf("missing database cfg parameter")
	}
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&timeout=%ds",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Schema,
		DbConnectTimeoutSecs)
	db, err := sql.Open("mysql", str)
	if err != nil {
		return nil, fmt.Errorf("error establishing a connection: %v", err.Error())
	}
	db.SetConnMaxLifetime(time.Second * time.Duration(DbExecTimeoutSecs))

	return &Article{conn: db}, nil
}