package main

import (
        "github.com/ghodss/yaml"
        "github.com/pkg/errors"
        "golang.org/x/crypto/bcrypt"
        "github.com/mkideal/cli"
        "github.com/go-openapi/jsonreference"
        "github.com/go-openapi/spec"
        "github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
        gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
        "github.com/spf13/viper"
        "github.com/spf13/pflag"
        log "github.com/sirupsen/logrus"
        "github.com/swaggo/files"
        "github.com/swaggo/gin-swagger"
        
)

type cliArgs struct {
	cli.Helper
	ConnectionStr string `cli:"*c,*conn" usage:"mysql connection str" dft:"$GATEWAY_CONN_STR"`
	ListenAddr    string `cli:"*l,*listen" usage:"gateway listen host and port" dft:"$GATEWAY_LS"`
	ResourceURL   string `cli:"*r,*resource" usage:"gateway resource url" dft:"$GATEWAY_RESOURCE_URL"`
}

func main() {
	cli.Run(new(cliArgs), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*cliArgs)
		server.NewGatewayServer(argv.ConnectionStr, argv.ResourceURL).Start(argv.ListenAddr)
		return nil
	})
}


