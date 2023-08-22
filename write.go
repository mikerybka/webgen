package webgen

import (
	"bytes"
	"embed"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	"github.com/mikerybka/apps/pkg/english"
	"github.com/mikerybka/apps/pkg/web/util"
)

func Write(dir string, data any) error {
	// Write index.json
	err := writeJSON(dir, data)
	if err != nil {
		return err
	}

	// Write index.html
	err = writeHTML(dir, data)
	if err != nil {
		return err
	}

	// Write index.css
	err = writeCSS(dir, data)
	if err != nil {
		return err
	}

	// Write index.js
	err = writeJS(dir, data)
	if err != nil {
		return err
	}

	// // Write children
	// t := reflect.TypeOf(data)
	// switch t.Kind() {
	// case reflect.Struct:
	// 	err = writeStruct(dir, data)
	// case reflect.Slice:
	// 	err = writeSlice(dir, data)
	// case reflect.Map:
	// 	err = writeMap(dir, data)
	// }
	// if err != nil {
	// 	return err
	// }

	// Write update handler (PUT)
	err = writeUpdateHandler(dir, data)
	if err != nil {
		return err
	}

	// Write method handlers (POST)
	err = writeMethodHandlers(dir, data)
	if err != nil {
		return err
	}

	return nil
}

func writeJSON(dir string, data any) error {
	path := filepath.Join(dir, "index.json")
	return util.WriteJSON(path, data)
}

func writeHTML(dir string, data any) error {
	b, err := embeded.ReadFile("embed/index.html")
	if err != nil {
		return err
	}
	path := filepath.Join(dir, "index.html")
	return os.WriteFile(path, b, os.ModePerm)
}

func writeCSS(dir string, data any) error {
	b, err := embeded.ReadFile("embed/index.css")
	if err != nil {
		return err
	}
	path := filepath.Join(dir, "index.css")
	return os.WriteFile(path, b, os.ModePerm)
}

func writeJS(dir string, data any) error {
	b, err := embeded.ReadFile("embed/index.js")
	if err != nil {
		return err
	}
	path := filepath.Join(dir, "index.js")
	return os.WriteFile(path, b, os.ModePerm)
}

//go:embed embed/*
var embeded embed.FS

func writeUpdateHandler(dir string, data any) error {
	b, err := embeded.ReadFile("embed/PUT/main.go.tmpl")
	if err != nil {
		return err
	}
	tmpl := string(b)
	t := template.Must(template.New("main.go.tmpl").Parse(tmpl))
	t.Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	})
	buf := bytes.NewBuffer(nil)
	err = t.Execute(buf, data)
	if err != nil {
		return err
	}
	path := filepath.Join(dir, "PUT", "main.go")
	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}
	return os.WriteFile(path, buf.Bytes(), os.ModePerm)
}

func writeMethodHandlers(dir string, data any) error {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		err := writeMethodHandler(dir, method)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeMethodHandler(dir string, method reflect.Method) error {
	name := english.ParsePascalCaseWithAcronyms(method.Name)
	path := filepath.Join(dir, name.KebabCase(), "POST", "main.go")
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}
	b, err := embeded.ReadFile("embed/POST/main.go.tmpl")
	if err != nil {
		return err
	}
	tmpl := string(b)
	t := template.Must(template.New("main.go.tmpl").Parse(tmpl))
	t.Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	})
	buf := bytes.NewBuffer(nil)
	err = t.Execute(buf, method)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, buf.Bytes(), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
