package db

import (
	"strconv"

	gc "github.com/GreptimeTeam/greptimedb-client-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Greptime struct {
	Host     string // 127.0.0.1
	Port     string // 4001
	User     string
	Password string

	Client gc.Client
}

func (g *Greptime) Setup() error {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	cfg := gc.NewCfg(g.Host).
		WithAuth(g.User, g.Password).
		WithDialOptions(options...)

	if len(g.Port) > 0 {
		port, err := strconv.Atoi(g.Port)
		if err != nil {
			return err
		}
		cfg.WithPort(port)
	}

	cli, err := gc.NewClient(cfg)
	if err != nil {
		return err
	}

	g.Client = *cli
	return nil
}
