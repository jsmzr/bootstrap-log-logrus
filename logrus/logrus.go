package logrus

import (
	"github.com/jsmzr/bootstrap-log/log"
	"github.com/sirupsen/logrus"
)

type LogrusConfig struct{}

type LogrusContainer struct {
	logger *logrus.Logger
}

func (c *LogrusConfig) Load() (log.Logger, error) {
	// TODO 一些配置支持
	return &LogrusContainer{
		logger: logrus.New(),
	}, nil
}

func (c *LogrusContainer) Debug(msg string, params ...interface{}) {
	c.logger.Debug(append([]interface{}{msg}, params...)...)
}
func (c *LogrusContainer) Info(msg string, params ...interface{}) {
	c.logger.Info(append([]interface{}{msg}, params...)...)
}
func (c *LogrusContainer) Warn(msg string, params ...interface{}) {

	c.logger.Warn(append([]interface{}{msg}, params...)...)
}
func (c *LogrusContainer) Error(msg string, params ...interface{}) {
	c.logger.Error(append([]interface{}{msg}, params...)...)
}

func init() {
	log.Register("logrus", &LogrusConfig{})
}
