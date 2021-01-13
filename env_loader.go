package env

/* env_loader.go provides utility function Load to load structure attributes values */

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

var (
	ErrNotStruct = errors.New("required struct")
	ErrUnmarshal = errors.New("unmarshal")
)

func Load(target interface{}) (err error) {
	var (
		val string // temp storage for values loaded from env
		ok  bool   // value set indicator
		env string // env key of current field
	)

	configType := reflect.TypeOf(target).Elem()
	targetValue := reflect.ValueOf(target).Elem()

	if configType.Kind() != reflect.Struct {
		return fmt.Errorf("%w: %T", ErrNotStruct, target)
	}

	for i := 0; i < configType.NumField(); i++ {
		currentField := configType.Field(i)
		if env, ok = currentField.Tag.Lookup("env"); !ok {
			// no env tag, skip field
			continue
		}
		// variable defined, deal with field
		configField := targetValue.Field(i)
		if !configField.CanSet() {
			// cant set value, do nothing
			continue
		}
		// access field value, will use it further
		fieldValue := configField.Interface()
		loader, isLoader := fieldValue.(Loader)
		// may be current field is Loader?
		if isLoader {
			// call field LoadEnv method and check its loading error
			if err = loader.LoadEnv(); err != nil {
				return fmt.Errorf("%w: %s: %s", ErrUnmarshal, currentField.Name, err)
			}
			// current field done, do next
			continue
		}
		// assume current field assignable, has tag and not nil, take env value from operation system.
		if val, ok = os.LookupEnv(env); !ok {
			// no env variable, skip field
			continue
		}
		// value received, set value, detect way how to set value
		targetKind := currentField.Type.Kind()
		unmarshaler, isUnmarshaler := fieldValue.(StringScanner)
		// check field type and implemented interfaces
		switch {
		case isUnmarshaler:
			// field type is StringScanner, pass value into
			if err = unmarshaler.ScanString(val); err != nil {
				return fmt.Errorf("%w: %s: %s", ErrUnmarshal, currentField.Name, err)
			}
		case targetKind == reflect.String:
			// field type string, just set value as is
			configField.SetString(val)
		case targetKind == reflect.Bool:
			// field type bool, translate with boolean helper
			configField.SetBool(toBool(val))
		case targetKind == reflect.Int:
			int64Value, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return fmt.Errorf("%w: %s: %v", ErrUnmarshal, currentField.Name, err.Error())
			}
			// field type bool, translate with boolean helper
			configField.SetInt(int64Value)
		}
	}

	return nil
}
