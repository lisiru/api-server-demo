package api_server_demo

import (
	"api-server-demo/config"
	"api-server-demo/options"
	"api-server-demo/pkg/app"
	"api-server-demo/pkg/logger"
	"api-server-demo/pkg/server"
	"os"
)

func NewApp() error  {
	opts:=options.NewOptions()
	if err:=app.AddConfigToOptions(opts);err!=nil {
		os.Exit(1)
	}

	logger.Init(opts.Log)
	defer logger.Flush()
	cfg,err:=config.CreateConfigFromOptions(opts)
	if err!=nil {
		return err
	}
	stopCh:=server.SetupSignalHandler()
	return Run(cfg,stopCh)
}
