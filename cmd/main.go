package main

import (
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/config"
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/infrastructure"
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/routes"
)

func main() {
	config := config.GetConfig()
	db, err := infrastructure.OpenMysqlConn(config)
	if err != nil {
		panic(err)
	}
	routes.CreateUrlMappings(db)
	routes.Router.Run(":8080")
}
