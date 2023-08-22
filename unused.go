package webgen

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strconv"

	"github.com/mikerybka/apps/pkg/english"
)

func writeStruct(dir string, data any) error {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		if field.Anonymous {
			return fmt.Errorf("anonymous fields not supported")
		} else {
			name := english.ParsePascalCaseWithAcronyms(field.Name)
			fmt.Println(field.Name, name, name.KebabCase())
			err := Write(filepath.Join(dir, name.KebabCase()), value.Interface())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func writeSlice(dir string, data any) error {
	s := reflect.ValueOf(data)
	for i := 0; i < s.Len(); i++ {
		err := Write(filepath.Join(dir, strconv.Itoa(i)), s.Index(i).Interface())
		if err != nil {
			return err
		}
	}
	return nil
}

func writeMap(dir string, data any) error {
	d := data.(map[string]any)
	for k, v := range d {
		err := Write(filepath.Join(dir, k), v)
		if err != nil {
			return err
		}
	}
	return nil
}
