package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dembygenesis/blog_exam/pkg/config"
	"github.com/dembygenesis/blog_exam/pkg/models_autogenerated"
	"github.com/dembygenesis/blog_exam/pkg/utils/string_utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"
)

type Article struct {
	conn *sql.DB
}

const DbConnectTimeoutSecs int = 15
const DbExecTimeoutSecs int = 15

// NewMySQLStore returns a new Article MYSQL store instance
func NewMySQLStore(cfg *config.Database) (*Article, error) {
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

func (a *Article) Create() error {
	return nil
}

func (a *Article) Read(id int) error {
	boilCtx := boil.WithDebug(context.Background(), true)
	articles, err := models_autogenerated.Articles().All(boilCtx, a.conn)
	if err != nil {
		return err
	}
	fmt.Println( string_utils.ToJSON(articles)  )
	return nil
}
