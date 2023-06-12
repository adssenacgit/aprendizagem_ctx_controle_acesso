package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)

func OpenMysqlConn(config config.Config) (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPw, config.DBHost, config.DBPort, config.DBDatabase)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)

	return db, nil

}
