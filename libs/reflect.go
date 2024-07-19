package libs

import "reflect"

// Assign 结构体赋值.
func Assign(origin, target interface{}, excludes ...string) {
	valOrigin := reflect.ValueOf(origin).Elem()
	valTarget := reflect.ValueOf(target).Elem()

	for i := 0; i < valOrigin.NumField(); i++ {
		if !valTarget.FieldByName(valOrigin.Type().Field(i).Name).IsValid() {
			continue
		}
		isExclude := false
		for _, col := range excludes {
			if valOrigin.Type().Field(i).Name == col {
				isExclude = true
				break
			}
		}
		if isExclude {
			continue
		}
		tmpOrigin := valOrigin.Field(i)
		tmpTarget := valTarget.FieldByName(valOrigin.Type().Field(i).Name)
		if reflect.TypeOf(tmpOrigin.Interface()) != reflect.TypeOf(tmpTarget.Interface()) {
			continue
		}
		tmpTarget.Set(tmpOrigin)
	}
}
