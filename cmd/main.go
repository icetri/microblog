package main

import (
	"flag"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"microblog/pkg"
	"microblog/postgres"
	"microblog/server"
	"microblog/server/handlers"
	"microblog/server/service"
	"microblog/types"
	"os"
)

func main() {
	configPath := new(string)
	flag.StringVar(configPath, "configs-path", "configs/configs-host.yaml", "path to yaml configs file")
	flag.Parse()

	f, err := os.Open(*configPath)
	if err != nil {
		pkg.FatalError(errors.Wrap(err, "err with os.Open"))
	}

	cfg := &types.Config{}
	if err = yaml.NewDecoder(f).Decode(&cfg); err != nil {
		pkg.FatalError(errors.Wrap(err, "err with yaml.NewDecoder"))
	}

	pg, err := postgres.NewSQL(cfg.PsqlInfo)
	if err != nil {
		pkg.FatalError(errors.Wrap(err, "err with postgres.NewSQL"))
	}
	defer pg.Close()

	s, err := service.NewService(pg)
	if err != nil {
		pkg.FatalError(errors.Wrap(err, "err with service.NewService"))
	}

	endpoints, err := handlers.NewHandler(s)
	if err != nil {
		pkg.FatalError(errors.Wrap(err, "err with handlers.NewHandle"))
	}

	server.StartServer(cfg.ServerPort, endpoints)
}
