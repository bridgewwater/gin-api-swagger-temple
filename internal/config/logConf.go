package config

import (
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog"
)

// Initialization log
func (c *Config) initLog() error {

	if err := zlog.ZapLoggerInitByViper(); err != nil {
		return err
	}

	return nil
}
