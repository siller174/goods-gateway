package logger

import (
	"context"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

const RequestLogField = "x-request-id"

type Logger struct {
	Level    string
	FilePath string
}

func (logger *Logger) Init() error {
	level, err := logrus.ParseLevel(logger.Level)
	if err != nil {
		return err
	}
	logFile, err := os.OpenFile(logger.FilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(mw)
	logrus.SetLevel(level)
	return nil
}

func Debug(message string, args ...interface{}) {
	logrus.Debugf(message, args...)
}

func Info(message string, args ...interface{}) {
	logrus.Infof(message, args...)
}

func Warn(message string, args ...interface{}) {
	logrus.Warnf(message, args...)
}

func Error(message string, args ...interface{}) {
	logrus.Errorf(message, args...)
}

func Fatal(message string, args ...interface{}) {
	logrus.Fatalf(message, args...)
}

func DebugCtx(ctx context.Context, message string, args ...interface{}) {
	rqid := getRequestID(ctx)
	logrus.WithField(RequestLogField, rqid).Debugf(message, args...)
}

func InfoCtx(ctx context.Context, message string, args ...interface{}) {
	rqid := getRequestID(ctx)
	logrus.WithField(RequestLogField, rqid).Infof(message, args...)
}

func WarnCtx(ctx context.Context, message string, args ...interface{}) {
	rqid := getRequestID(ctx)
	logrus.WithField(RequestLogField, rqid).Warnf(message, args...)
}

func ErrorCtx(ctx context.Context, message string, args ...interface{}) {
	rqid := getRequestID(ctx)
	logrus.WithField(RequestLogField, rqid).Errorf(message, args...)
}

func FatalCtx(ctx context.Context, message string, args ...interface{}) {
	rqid := getRequestID(ctx)
	logrus.WithField(RequestLogField, rqid).Fatalf(message, args...)
}

func getRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	t := ctx.Value(RequestLogField)
	if t != nil {
		return t.(string)
	} else {
		return ""
	}
}
