package api_server_demo

import (
	"api-server-demo/cache/redis"
	"api-server-demo/config"
	genericoptions "api-server-demo/pkg/options"
    genericapiserver "api-server-demo/pkg/server"
	"api-server-demo/store/mysql"
)

type apiServer struct {

	genericAPIServer *genericapiserver.GenericAPIServer
}

type preparedAPIServer struct {
	*apiServer
}

type ExtraConfig struct {
	Addr         string
	mysqlOptions *genericoptions.MySQLOptions
	redisOptions *genericoptions.RedisOptions
}

func buildExtraConfig(cfg *config.Config) *ExtraConfig {
	return &ExtraConfig{

		mysqlOptions: cfg.MySQLOptions,
		redisOptions: cfg.RedisOptions,
		// etcdOptions:      cfg.EtcdOptions,
	}
}

type completedExtraConfig struct {
	*ExtraConfig
}

func (c *ExtraConfig) complete() *completedExtraConfig {
	if c.Addr == "" {
		c.Addr = "127.0.0.1:8081"
	}

	return &completedExtraConfig{c}
}

func (c *completedExtraConfig) New() {

	_, _ = mysql.GetMySQLFactoryOr(c.mysqlOptions)
	_, _ = redis.NewRedisFactoryOr(c.redisOptions)
}
func createAPIServer(cfg *config.Config) (*apiServer, error) {

	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}


	genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}

	buildExtraConfig(cfg).complete().New()
	server := &apiServer{
		genericAPIServer: genericServer,
	}

	return server, nil
}
func (s *apiServer) PrepareRun() preparedAPIServer {
	initRouter(s.genericAPIServer.Engine)

	return preparedAPIServer{s}
}

func (s preparedAPIServer) Run(stopCh <-chan struct{}) error {


	return s.genericAPIServer.Run(stopCh)
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()
	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	if lastErr = cfg.InsecuresServing.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	return
}
