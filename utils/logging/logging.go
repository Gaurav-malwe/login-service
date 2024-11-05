package log

import (
	"context"

	"github.com/Gaurav-malwe/login-service/internal/constants"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func SetCustomLogger(l *logrus.Logger) {

	logger = l

}

// Fields wraps logrus.Fields, which is a map[string]interface{}

type Fields logrus.Fields

type LoggerFields struct {
	LoggerFieldsMap map[string]interface{}
}

func SetLevel(level logrus.Level) {

	logger.Level = level

}

func SetLogLevel(level string) {

	switch level {

	case logrus.DebugLevel.String():

		logger.Level = logrus.DebugLevel

	case logrus.InfoLevel.String():

		logger.Level = logrus.InfoLevel

	case logrus.ErrorLevel.String():

		logger.Level = logrus.ErrorLevel

	case logrus.WarnLevel.String():

		logger.Level = logrus.WarnLevel

	default:

		logger.Level = logrus.InfoLevel

	}

}

func SetFormatter(format string) {

	if format == "text" {

		logger.SetFormatter(&logrus.TextFormatter{

			FieldMap: logrus.FieldMap{

				logrus.FieldKeyLevel: "level",

				logrus.FieldKeyTime: "time",

				logrus.FieldKeyMsg: "msg",
			},
		})

	} else {

		logger.SetFormatter(&logrus.JSONFormatter{

			FieldMap: logrus.FieldMap{

				logrus.FieldKeyLevel: "level",

				logrus.FieldKeyTime: "time",

				logrus.FieldKeyMsg: "msg",
			},
		})

	}

}

// Debug logs a message at level Debug on the standard logger.

func Debug(args ...interface{}) {

	if logger.Level >= logrus.DebugLevel {

		entry := logger.WithFields(logrus.Fields{})

		entry.Debug(args...)

	}

}

// Debug logs a message with fields at level Debug on the standard logger.

func DebugWithFields(l interface{}, f Fields) {

	if logger.Level >= logrus.DebugLevel {

		entry := logger.WithFields(logrus.Fields(f))

		entry.Debug(l)

	}

}

// Info logs a message at level Info on the standard logger.

func Info(args ...interface{}) {

	if logger.Level >= logrus.InfoLevel {

		entry := logger.WithFields(logrus.Fields{})

		entry.Info(args...)

	}

}

// Debug logs a message with fields at level Debug on the standard logger.

func InfoWithFields(l interface{}, f Fields) {

	if logger.Level >= logrus.InfoLevel {

		entry := logger.WithFields(logrus.Fields(f))

		entry.Info(l)

	}

}

// Warn logs a message at level Warn on the standard logger.

func Warn(args ...interface{}) {

	if logger.Level >= logrus.WarnLevel {

		entry := logger.WithFields(logrus.Fields{})

		entry.Warn(args...)

	}

}

// Debug logs a message with fields at level Debug on the standard logger.

func WarnWithFields(l interface{}, f Fields) {

	if logger.Level >= logrus.WarnLevel {

		entry := logger.WithFields(logrus.Fields(f))

		entry.Warn(l)

	}

}

// Error logs a message at level Error on the standard logger.

func Error(args ...interface{}) {

	if logger.Level >= logrus.ErrorLevel {

		entry := logger.WithFields(logrus.Fields{})

		entry.Error(args...)

	}

}

// Debug logs a message with fields at level Debug on the standard logger.

func ErrorWithFields(l interface{}, f Fields) {

	if logger.Level >= logrus.ErrorLevel {

		entry := logger.WithFields(logrus.Fields(f))

		entry.Error(l)

	}

}

// Fatal logs a message at level Fatal on the standard logger.

func Fatal(args ...interface{}) {

	if logger.Level >= logrus.FatalLevel {

		entry := logger.WithFields(logrus.Fields{})

		entry.Fatal(args...)

	}

}

// Debug logs a message with fields at level Debug on the standard logger.

func FatalWithFields(l interface{}, f Fields) {

	if logger.Level >= logrus.FatalLevel {

		entry := logger.WithFields(logrus.Fields(f))

		entry.Fatal(l)

	}

}

// Panic logs a message at level Panic on the standard logger.

func Panic(args ...interface{}) {

	if logger.Level >= logrus.PanicLevel {

		entry := logger.WithFields(logrus.Fields{})

		entry.Panic(args...)

	}

}

// Debug logs a message with fields at level Debug on the standard logger.

func PanicWithFields(l interface{}, f Fields) {

	if logger.Level >= logrus.PanicLevel {

		entry := logger.WithFields(logrus.Fields(f))

		entry.Panic(l)

	}

}

// GetLogger retrieves the current logger from the context.
func GetLogger(ctx context.Context) *logrus.Entry {
	return logger.WithContext(ctx).WithField("CorrelationId", constants.CorrelationId)
}
