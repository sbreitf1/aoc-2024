package helper

import (
	"fmt"
	"reflect"
)

func Clone[T any](obj T) T {
	return clone(reflect.ValueOf(obj)).Interface().(T)
}

func clone(src reflect.Value) reflect.Value {
	ordinaryTypes := map[reflect.Kind]bool{
		reflect.Int:        true,
		reflect.Int8:       true,
		reflect.Int16:      true,
		reflect.Int32:      true,
		reflect.Int64:      true,
		reflect.Uint:       true,
		reflect.Uint8:      true,
		reflect.Uint16:     true,
		reflect.Uint32:     true,
		reflect.Uint64:     true,
		reflect.Float32:    true,
		reflect.Float64:    true,
		reflect.Bool:       true,
		reflect.Complex64:  true,
		reflect.Complex128: true,
		reflect.String:     true,
	}

	if _, ok := ordinaryTypes[src.Type().Kind()]; ok {
		return src
	}

	switch src.Type().Kind() {
	case reflect.Struct:
		return cloneStruct(src)

	case reflect.Slice:
		return cloneSlice(src)

	case reflect.Map:
		return cloneMap(src)

	case reflect.Pointer:
		return clonePointer(src)

	default:
		panic(fmt.Sprintf("clone for kind '%v' not defined", src.Type().Kind()))
	}
}

func cloneStruct(src reflect.Value) reflect.Value {
	dst := reflect.New(src.Type()).Elem()
	for i := 0; i < src.NumField(); i++ {
		if dst.Field(i).CanSet() {
			dst.Field(i).Set(clone(src.Field(i)))
		} else {
			//TODO assign unexported field with unsafe (https://stackoverflow.com/questions/42664837/how-to-access-unexported-struct-fields)
		}
	}
	return dst
}

func cloneSlice(src reflect.Value) reflect.Value {
	dst := reflect.MakeSlice(src.Type(), src.Len(), src.Len())
	for i := 0; i < src.Len(); i++ {
		dst.Index(i).Set(reflect.ValueOf(Clone(src.Index(i).Interface())))
	}
	return dst
}

func cloneMap(src reflect.Value) reflect.Value {
	dst := reflect.MakeMapWithSize(src.Type(), src.Len())
	it := src.MapRange()
	for it.Next() {
		dst.SetMapIndex(it.Key(), clone(it.Value()))
	}
	return dst
}

func clonePointer(src reflect.Value) reflect.Value {
	obj := reflect.New(src.Elem().Type())
	obj.Elem().Set(clone(src.Elem()))
	return obj
}
