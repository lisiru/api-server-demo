package options

import (
	"api-server-demo/pkg/logger"
	genericoptions "api-server-demo/pkg/options"
)

type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions `json:"server" mapstructure:"server"`
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	InsecuresServing *genericoptions.InsecureServerOptions `json:"insecure" mapstructure:"insecure"`
	Log *logger.Options `json:"log" mapstructure:"log"`
	RedisOptions *genericoptions.RedisOptions `json:"redis" mapstructure:"redis"`
}

func NewOptions() *Options  {
	o:=Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		MySQLOptions: genericoptions.NewMySQLOptions(),
		InsecuresServing: genericoptions.NewInsecureServerOptions(),
		RedisOptions: genericoptions.NewRedisOptions(),
		Log: logger.NewOptions(),

	}
	return &o
}
