package mysql

import (
	"database/sql"
	"fmt"
	"github.com/dembygenesis/blog_exam/pkg/config"
	"github.com/dembygenesis/blog_exam/pkg/utils/string_utils"
	"github.com/friendsofgo/errors"
	"time"
)

type Article struct {
	conn *sql.DB
}

const DbConnectTimeoutSecs int = 15
const DbExecTimeoutSecs int = 15

func getConnection(schema string, cfg *config.Database) (*sql.DB, error) {
	// Setup connection
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&timeout=%ds",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		schema,
		DbConnectTimeoutSecs)
	db, err := sql.Open("mysql", str)
	if err != nil {
		return nil, fmt.Errorf("error establishing a connection: %v, conn details: %v", err.Error(), str)
	}
	db.SetConnMaxLifetime(time.Second * time.Duration(DbExecTimeoutSecs))
	return db, nil
}

// NewStore returns a new Article MYSQL store instance
func NewStore(cfg *config.Database) (*Article, error) {
	fmt.Println("cfg:", string_utils.ToJSON(cfg))
	if cfg == nil {
		return nil, fmt.Errorf("missing database cfg parameter")
	}

	db, err := getConnection(cfg.Schema,cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error trying to establish a connection")
	}

	return &Article{conn: db}, nil
}