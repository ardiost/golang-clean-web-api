package logging

type Logger interface {
	Init()

	Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Infoof(template string, args ...interface{})

	Warn(err error, cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Warnf(err error, template string, args ...interface{})

	Fatal(err error, cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Fatalf(err error, template string, args ...interface{})
}
