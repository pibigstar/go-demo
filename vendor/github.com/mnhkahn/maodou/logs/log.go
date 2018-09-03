package logs

import (
	"fmt"
)

type LogContainer interface {
	Log(format string, a ...interface{})
}

type Log interface {
	NewLogImpl(dsn string) (LogContainer, error)
}

var logss = make(map[string]Log)

func Register(name string, log Log) {
	if log == nil {
		panic("log: Register log is nil")
	}
	if _, ok := logss[name]; ok {
		panic("log: Register called twice for adapter " + name)
	}
	logss[name] = log
}

func NewLog(log_name, dsn string) (LogContainer, error) {
	log, ok := logss[log_name]
	if !ok {
		return nil, fmt.Errorf("parser: unknown log_name %q", log_name)
	}
	return log.NewLogImpl(dsn)
}
