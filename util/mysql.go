package util

import (
	"github.com/fatih/structs"
	"reflect"
)

func IsDuplicateError(err error) bool {

	isStruct := false

	vType := reflect.ValueOf(err)

endfor:
	for {
		switch vType.Kind() {
		case reflect.Ptr:
			vType = vType.Elem()

		case reflect.Struct:
			isStruct = true
			break endfor

		default:
			break endfor
		}
	}
	if !isStruct {
		return false
	}
	v := vType.FieldByName("Number")
	if v.IsValid() {
		return v.Uint() == 1062
	}

	//
	//
	//
	//	err_struct := structs.New(err)
	//	f, ok := err_struct.FieldOk("Number")
	//	if ok {
	//		v := f.Value()
	//		switch v.(type) {
	//		case uint16:
	//			return v.(uint16) == 1026
	//		}
	//	}

	return false
}

func IsDuplicateError2(err error) bool {

	isStruct := false

	vType := reflect.ValueOf(err)

endfor:
	for {
		switch vType.Kind() {
		case reflect.Ptr:
			vType = vType.Elem()

		case reflect.Struct:
			isStruct = true
			break endfor

		default:
			break endfor
		}
	}
	if !isStruct {
		return false
	}
	vType.FieldByName("Number")

	err_struct := structs.New(err)
	f, ok := err_struct.FieldOk("Number")
	if ok {
		v := f.Value()
		switch v.(type) {
		case uint16:
			return v.(uint16) == 1062
		}
	}

	return false
}
