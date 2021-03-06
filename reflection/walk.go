package reflection

import "reflect"

func Walk(x any, fn func(input string)) {
	value := reflect.ValueOf(x)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
