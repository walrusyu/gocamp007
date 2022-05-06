package log

type Logger interface {
	Log(level Level, kv ...interface{}) error
}
