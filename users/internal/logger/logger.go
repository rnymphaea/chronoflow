package logger

type Logger interface {
	Debug(msg string)
	Debugf(msg string, fields map[string]interface{})

	Info(msg string)
	Infof(msg string, fields map[string]interface{})

	Warn(msg string)
	Warnf(msg string, fields map[string]interface{})

	Error(err error, msg string)
	Errorf(err error, msg string, fields map[string]interface{})

	Fatal(err error, msg string)
	Fatalf(err error, msg string, fields map[string]interface{})

	With(fields map[string]interface{}) Logger

	Component(name string) Logger
}
