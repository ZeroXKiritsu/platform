package services

import (
	"context"
	"errors"
	"reflect"
)

func Call(target any, otherArgs ...any) ([]any, error) {
	return CallForContext(context.Background(), target, otherArgs...)
}

func CallForContext(c context.Context, target any, otherArgs ...any) (results []any, err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Func {
		resultVals := invokeFunction(c, targetValue, otherArgs...)
		results = make([]any, len(resultVals))
		for i := 0; i < len(resultVals); i++ {
			results[i] = resultVals[i].Interface()
		}
	} else {
		err = errors.New("Only function can be invoked")
	}
	return
}
