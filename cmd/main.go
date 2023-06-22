package main

import (
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/config"
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/infrastructure"
	"github.com/adssenacgit/aprendizagem_ctx_controle_acesso/pkg/routes"
)

func main() {

	config := config.GetConfig()
	infrastructure.OpenMysqlConn(config)
	infrastructure.Migrate()
	routes.CreateUrlMappings(infrastructure.DB)
	routes.Router.Run(config.Port)
}
