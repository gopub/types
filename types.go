package types

import (
	"github.com/google/uuid"
	"reflect"
	"runtime"
	"strings"
)

func FuncNameOf(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func NameOf(i interface{}) string {
	return reflect.TypeOf(i).Name()
}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

// AllocValue allocate value: ppObj should be the address of a pointer to a value
func AllocValue(ppObj interface{}) {
	v := reflect.ValueOf(ppObj)
	if v.Kind() != reflect.Ptr {
		panic("pointer required")
	}

	v = v.Elem()

	//v is a pointer to a non-pointer value
	if v.Kind() != reflect.Ptr {
		return
	}

	//v is a pointer to a non-nil pointer
	if !v.IsNil() {
		return
	}

	v.Set(reflect.New(v.Type().Elem()))
}

func MakeZero(ptr interface{}) {
	v := reflect.ValueOf(ptr).Elem()
	v.Set(reflect.Zero(v.Type()))
}

func Renew(ptrDst interface{}, src interface{}) {
	pdv := reflect.ValueOf(ptrDst)
	sv := reflect.ValueOf(src)
	if sv.Kind() == reflect.Ptr {
		pdv.Elem().Set(reflect.New(sv.Type().Elem()))
	} else {
		pdv.Elem().Set(reflect.Zero(sv.Type()))
	}
}

func DeepNew(t reflect.Type) reflect.Value {
	v := reflect.New(t)
	e := v.Elem()
	for e.Kind() == reflect.Ptr {
		e.Set(reflect.New(e.Type().Elem()))
		e = e.Elem()
	}

	if e.Kind() != reflect.Struct {
		return v
	}

	for i := 0; i < e.NumField(); i++ {
		ft := e.Type().Field(i)
		if ft.Type.Kind() == reflect.Ptr {
			e.Field(i).Set(DeepNew(ft.Type.Elem()))
		}
	}
	return v
}

func NewUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
