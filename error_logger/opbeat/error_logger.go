package opbeat

import (
	opbeatClient "github.com/upfluence/goutils/Godeps/_workspace/src/github.com/roncohen/opbeat-go"
	"github.com/upfluence/goutils/error_logger"
	"runtime"
)

type Logger struct {
	client *opbeatClient.Opbeat
}

func NewErrorLogger() *Logger {
	return &Logger{opbeatClient.NewFromEnvironment()}
}

func (l *Logger) Capture(err error, opts *error_logger.Options) error {
	var options *opbeatClient.Options

	extra := make(opbeatClient.Extra)
	trace := make([]byte, 16384)
	runtime.Stack(trace, true)

	extra["stacktrace"] = string(trace)

	if opts != nil {
		for k, v := range *opts {
			extra[k] = v
		}
	}

	options = &opbeatClient.Options{Extra: &extra}

	return l.client.CaptureError(err, options)
}

func (l *Logger) Close() {
	l.client.Close()
}
