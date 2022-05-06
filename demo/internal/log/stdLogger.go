package log

import (
	"fmt"
	"strings"
)

var _ Logger = &StdLogger{}

type StdLogger struct {
}

func (l *StdLogger) Log(level Level, kv ...interface{}) error {
	sb := strings.Builder{}
	sb.WriteString(level.String() + ":")
	for i := 0; i < len(kv); i += 2 {
		sb.WriteString(fmt.Sprintf(" %s=%v&", kv[i], kv[i+1]))
	}
	str := sb.String()
	str = str[:len(str)-1]
	fmt.Printf(str)
	return nil
}
