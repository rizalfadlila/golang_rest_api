package logger

import (
	"context"
	"fmt"

	"github.com/rest_api/pkg/util"

	"github.com/sirupsen/logrus"
)

type logField map[string]interface{}

// Field :nodoc:
type Field struct {
	Key   string
	Value interface{}
}

// FieldFunc :nodoc:
type FieldFunc func(key string, value interface{}) *Field

// SetField :nodoc:
func SetField(k string, v interface{}) Field {
	return Field{
		Key:   k,
		Value: v,
	}

}

// Any :nodoc:
func Any(k string, v interface{}) Field {
	return Field{
		Key:   k,
		Value: v,
	}

}

// EventName :nodoc:
func EventName(v interface{}) Field {
	return Field{
		Key:   "name",
		Value: v,
	}
}

// EventID :nodoc:
func EventID(v interface{}) Field {
	return Field{
		Key:   "id",
		Value: v,
	}
}

// SetMessageFormat :nodoc:
func SetMessageFormat(format string, args ...interface{}) interface{} {
	return fmt.Sprintf(format, args...)
}

func extract(args ...Field) map[string]interface{} {
	data := map[string]interface{}{}

	if len(args) == 0 {
		return data
	}

	for _, fl := range args {
		data[fl.Key] = fl.Value
	}

	return data
}

// Error :nodoc:
func Error(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Error(arg)
}

// Info :nodoc:
func Info(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Info(arg)
}

// Debug :nodoc:
func Debug(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Debug(arg)
}

// Fatal :nodoc:
func Fatal(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Fatal(arg)
}

// Panic :nodoc:
func Panic(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Panic(arg)
}

// Warn :nodoc:
func Warn(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Warn(arg)
}

// Trace :nodoc:
func Trace(arg interface{}, fl ...Field) {
	logrus.WithFields(map[string]interface{}{
		"event": extract(fl...),
	}).Trace(arg)
}

// AccessLog :nodoc:
func AccessLog(arg interface{}, fl ...Field) {
	logrus.WithFields(extract(fl...)).Info(arg)
}

// InfoWithContext :nodoc:
func InfoWithContext(ctx context.Context, arg interface{}, fl ...Field) {
	logrus.WithFields(extractContext(ctx.Value("access"), map[string]interface{}{
		"event": extract(fl...),
	})).WithContext(ctx).Info(arg)
}

func extractContext(i interface{}, logField map[string]interface{}) map[string]interface{} {
	if util.IsSameType(i, logField) {
		x := i.(map[string]interface{})
		for k, v := range x {
			logField[k] = v
		}
	}

	return logField
}
