package infrastructure

import (
	"fmt"
	"time"

	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/config"
	entities "github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/entities"
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

func OpenMysqlConn(config config.Config) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DBUser, config.DBPw, config.DBHost, config.DBPort, config.DBDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		helpers.HandleErr(err)
	}

	db.Use(
		dbresolver.Register(dbresolver.Config{}).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(10 * time.Minute).
			SetMaxIdleConns(2).
			SetMaxOpenConns(5),
	)

	DB = db
}

func Migrate() {
	DB.AutoMigrate(&entities.User{})
}
