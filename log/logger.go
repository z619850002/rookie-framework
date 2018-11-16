package mylog

import (
	"go.uber.org/zap"
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
	"log"
)

const (
	InfoLevel = 0
	WarnLevel = 1
	ErrorLevel = 2
	FatalError = 3
)


var (
	atom  zap.AtomicLevel
	logger *zap.Logger
	l *zap.SugaredLogger
)

func init() {
	atom = zap.NewAtomicLevel()
	// default : print all level log
	atom.SetLevel(zapcore.InfoLevel)

	cfg := zap.NewProductionConfig()
	cfg.Level = atom

	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	l = logger.Sugar()
}

func Info(format string, a ...interface{}) {
	l.Info(format,a)
}

func Warn(format string, a ...interface{}){
	l.Warn(format,a)
}

func Error(format string, a ...interface{}) {
	l.Error(format,a)
}

func Fatal(format string, a ...interface{}) {
	l.Fatal(format,a)
}

func SetLevel(level int) error {
	switch level {
	case InfoLevel:
		setLevel(zap.InfoLevel)
	case WarnLevel:
		setLevel(zap.WarnLevel)
	case ErrorLevel:
		setLevel(zap.ErrorLevel)
	case FatalError:
		setLevel(zap.FatalLevel)
	default:
		return errors.New("incorrect log level")
	}
	return nil
}

func setLevel(level zapcore.Level){
	atom.SetLevel(level)
}
