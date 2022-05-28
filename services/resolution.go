package services

import (
	"context"
	"errors"
	"reflect"
)

func GetService(target any) error {
	return GetServiceForContext(context.Background(), target)
}

func GetServiceForContext(c context.Context, target any) (err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr && targetValue.Elem().CanSet() {
		err = resolveServiceFromValue(c, targetValue)
	} else {
		err = errors.New("type cannot be used at target")
	}
	return
}
