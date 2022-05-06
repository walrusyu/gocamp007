package log

type WrappedLogger struct {
	logger Logger
}

func (l *WrappedLogger) Debug(kv ...interface{}) {
	l.logger.Log(LevelDebug, kv)
}

func (l *WrappedLogger) Info(kv ...interface{}) {
	l.logger.Log(LevelInfo, kv)
}

func (l *WrappedLogger) Warn(kv ...interface{}) {
	l.logger.Log(LevelWarn, kv)
}

func (l *WrappedLogger) Error(kv ...interface{}) {
	l.logger.Log(LevelError, kv)
}

func (l *WrappedLogger) Fatal(kv ...interface{}) {
	l.logger.Log(LevelFatal, kv)
}

func NewLogger() *WrappedLogger {
	return &WrappedLogger{
		logger: &StdLogger{},
	}
}
