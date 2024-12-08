package logging

import (
	"github.com/ardiost/golang-clean-web-api/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
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

func newZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}
	logger.Init()
	return logger
}
func (l *zapLogger) getLogLevel() zapcore.Level {
	level, exist := logLevelMap[l.cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}
	return level
}

func (l *zapLogger) Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   l.cfg.Logger.FilePath,
		MaxSize:    1,
		MaxAge:     5,
		LocalTime:  true,
		MaxBackups: 10,
		Compress:   true,
	})
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config), w, l.getLogLevel(),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()

	l.logger = logger
}

func (l *zapLogger) Debug(cat, Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.Debugw(msg.params...)
}
func (l *zapLogger) Debugf(template string, args ...interface{}){
	l.Logger.Debugf(template,args)
}

func (l *zapLogger) Info(cat, Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.Infofw(msg.params...)
}
func (l *zapLogger) Infof(template string, args ...interface{}){
	l.Logger.Infof(template,args)
}

func (l *zapLogger) Warn(cat, Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.Warnw(msg.params...)
}
func (l *zapLogger) Warnf(template string, args ...interface{}){
	l.Logger.Warnf(template,args)
}

func (l *zapLogger) Error(cat, Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.Errorw(msg.params...)
}
func (l *zapLogger) Errorf(template string, args ...interface{}){
	l.Logger.Errorf(template,args)
}

func (l *zapLogger) Fatal(cat, Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.Fatalw(msg.params...)
}
func (l *zapLogger) Fatalf(template string, args ...interface{}){
	l.Logger.Fatalf(template,args)
}

func prepareLogKeys (extra map[ExtraKey]interface{},cat Category ,sub SubCategory)[]interface{}{
	if extra == nil {
		extra= make(map[ExtraKey]interface{}, 0)
	}
	extra["Category"]=cat
	extra["SubCategory"]=sub
	params:= mapToZapParams(extra)
}