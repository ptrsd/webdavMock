package main

import (
	"github.com/ptrsd/webdavMock/pkg/repository"
	"github.com/ptrsd/webdavMock/pkg/server"
)

type Config struct {
	Server     server.Config     `yaml:server`
	Repository repository.Config `yaml:repository`
}

var defaultCfg = Config{
	Server: server.Config{
		Port: "7080",
	},
	Repository: repository.Config{
		Fallback: repository.FallbackConfig{
			Path: "./assets",
		},
	},
}

func main() {
	var (
		fallbackRepo repository.ReadRepository
		err          error
	)

	if fallbackRepo, err = repository.CreateFallback(defaultCfg.Repository.Fallback); err != nil {
		panic(err)
	}

	webDavSrv := server.Create(defaultCfg.Server, fallbackRepo)
	if err := webDavSrv.Start(); err != nil {
		panic(err)
	}
}
