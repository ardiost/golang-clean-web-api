package logging

type Logger interface {
	Init()

	Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Infoof(template string, args ...interface{})

	Warn( cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Warnf( template string, args ...interface{})

	Fatal( cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Fatalf( template string, args ...interface{})
}

func NewLogger(cfg *config.Config) Logger {
	return NewZapLogger(cfg)
}
