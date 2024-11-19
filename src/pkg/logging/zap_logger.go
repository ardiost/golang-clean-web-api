package logging

import (
	"github.com/ardiost/golang-clean-web-api/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

type zapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}
	logger.Init()
	return logger
}
func (l *zapLogger) Init() {

}

func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
func (l *zapLogger) Debugf(template string, args ...interface{})
