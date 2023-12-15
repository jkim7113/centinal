package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

func DecodeFormData(r *http.Request, result interface{}) {
	rv := reflect.ValueOf(result).Elem()
	r.ParseForm()
	for key, value := range r.Form {
		field := rv.FieldByName(key)
		if field == (reflect.Value{}) {
			err := errors.New(fmt.Sprintf("Couldn't find field %s in the given struct", key))
			PanicIfError(err)
		}
		switch field.Kind() {
			case reflect.Int:
				temp, err := strconv.Atoi(value[0])
				PanicIfError(err)
				field.Set(reflect.ValueOf(temp))
			case reflect.String:
				field.SetString(value[0])
		}
	}
}

func DecodeRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func EncodeResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}
