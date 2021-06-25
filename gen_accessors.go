// +build ignore

// gen-accessors generates accessor methods for structs with pointer fields.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
)

const (
	fileSuffix = "_accessors.go"
)

var (
	verbose = flag.Bool("v", false, "Print verbose log messages")

	sourceTmpl = template.Must(template.New("source").Parse(source))

	// blacklistStructMethod lists "struct.method" combos to skip.
	blacklistStructMethod = map[string]bool{}
	// blacklistStruct lists structs to skip.
	blacklistStruct = map[string]bool{
		"Client":        true,
		"ConnectClient": true,
	}
)

func logf(fmt string, args ...interface{}) {
	if *verbose {
		log.Printf(fmt, args...)
	}
}

func main() {
	flag.Parse()
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, ".", sourceFilter, 0)
	if err != nil {
		log.Fatal(err)
		return
	}

	for pkgName, pkg := range pkgs {
		t := &templateData{
			filename: pkgName + fileSuffix,
			Package:  pkgName,
			Imports:  map[string]string{},
		}
		for filename, f := range pkg.Files {
			logf("Processing %v...", filename)
			if err := t.processAST(f); err != nil {
				log.Fatal(err)
			}
		}
		if err := t.dump(); err != nil {
			log.Fatal(err)
		}
	}
	logf("Done.")
}

func (t *templateData) processAST(f *ast.File) error {
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			// Skip unexported identifiers.
			if !ts.Name.IsExported() {
				logf("Struct %v is unexported; skipping.", ts.Name)
				continue
			}
			// Check if the struct is blacklisted.
			if blacklistStruct[ts.Name.Name] {
				logf("Struct %v is blacklisted; skipping.", ts.Name)
				continue
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}
			for _, field := range st.Fields.List {
				if len(field.Names) == 0 {
					continue
				}
				fieldName := field.Names[0]
				// Skip unexported identifiers.
				if !fieldName.IsExported() {
					logf("Field %v is unexported; skipping.", fieldName)
					continue
				}
				// Check if "struct.method" is blacklisted.
				if key := fmt.Sprintf("%v.Get%v", ts.Name, fieldName); blacklistStructMethod[key] {
					logf("Method %v is blacklisted; skipping.", key)
					continue
				}
				switch typ := field.Type.(type) {
				case *ast.StarExpr:
					switch x := typ.X.(type) {
					case *ast.ArrayType:
						t.addArrayType(x, ts.Name.String(), fieldName.String(), true)
					case *ast.Ident:
						t.addIdent(x, ts.Name.String(), fieldName.String())
					case *ast.MapType:
						t.addMapType(x, ts.Name.String(), fieldName.String())
					case *ast.SelectorExpr:
						t.addSelectorExpr(x, ts.Name.String(), fieldName.String())
					default:
						logf("processAST: type %q, field %q, unknown %T: %+v", ts.Name, fieldName, x, x)
					}
				case *ast.ArrayType:
					t.addArrayType(typ, ts.Name.String(), fieldName.String(), false)
				}

			}
		}
	}
	return nil
}

func sourceFilter(fi os.FileInfo) bool {
	return !strings.HasSuffix(fi.Name(), "_test.go") && !strings.HasSuffix(fi.Name(), fileSuffix)
}

func (t *templateData) dump() error {
	if len(t.Getters) == 0 {
		logf("No getters for %v; skipping.", t.filename)
		return nil
	}

	// Sort getters by ReceiverType.FieldName.
	sort.Sort(byName(t.Getters))

	var buf bytes.Buffer
	if err := sourceTmpl.Execute(&buf, t); err != nil {
		return err
	}
	clean, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	logf("Writing %v...", t.filename)
	return ioutil.WriteFile(t.filename, clean, 0644)
}

func newGetter(receiverType, fieldName, fieldType, zeroValue string, namedStruct, slice bool) *getter {
	return &getter{
		sortVal:      strings.ToLower(receiverType) + "." + strings.ToLower(fieldName),
		ReceiverVar:  strings.ToLower(receiverType[:1]),
		ReceiverType: receiverType,
		FieldName:    fieldName,
		FieldType:    fieldType,
		ZeroValue:    zeroValue,
		NamedStruct:  namedStruct,
		Slice:        slice,
	}
}

func (t *templateData) addArrayType(x *ast.ArrayType, receiverType, fieldName string, fromPtr bool) {
	var eltType string
	switch elt := x.Elt.(type) {
	case *ast.Ident:
		eltType = elt.String()
	case *ast.StarExpr:
		eltType = "*" + elt.X.(*ast.Ident).String()
	default:
		logf("addArrayType: type %q, field %q: unknown elt type: %T %+v; skipping.", receiverType, fieldName, elt, elt)
		return
	}

	t.Getters = append(t.Getters, newGetter(receiverType, fieldName, "[]"+eltType, "nil", false, !fromPtr))
}

func (t *templateData) addIdent(x *ast.Ident, receiverType, fieldName string) {
	var zeroValue string
	var namedStruct = false
	switch x.String() {
	case "int", "int16", "int32", "int64":
		zeroValue = "0"
	case "float32", "float64":
		zeroValue = "0.0"
	case "string":
		zeroValue = `""`
	case "bool":
		zeroValue = "false"
	case "Date":
		zeroValue = "Date{}"
	case "TimestampParis":
		zeroValue = "TimestampParis{}"
	case "TimestampLondon":
		zeroValue = "TimestampLondon{}"
	case "Currency":
		zeroValue = `Currency("")`
	case "Level":
		zeroValue = "LevelNone"
	case "Review":
		zeroValue = "ReviewNone"
	case "AdditionalDataOneOf":
		zeroValue = "AdditionalDataOneOf{}"
	default:
		zeroValue = "nil"
		namedStruct = true
	}

	t.Getters = append(t.Getters, newGetter(receiverType, fieldName, x.String(), zeroValue, namedStruct, false))
}

func (t *templateData) addMapType(x *ast.MapType, receiverType, fieldName string) {
	var keyType string
	switch key := x.Key.(type) {
	case *ast.Ident:
		keyType = key.String()
	default:
		logf("addMapType: type %q, field %q: unknown key type: %T %+v; skipping.", receiverType, fieldName, key, key)
		return
	}

	var valueType string
	switch value := x.Value.(type) {
	case *ast.Ident:
		valueType = value.String()
	default:
		logf("addMapType: type %q, field %q: unknown value type: %T %+v; skipping.", receiverType, fieldName, value, value)
		return
	}

	fieldType := fmt.Sprintf("map[%v]%v", keyType, valueType)
	zeroValue := fmt.Sprintf("map[%v]%v{}", keyType, valueType)
	t.Getters = append(t.Getters, newGetter(receiverType, fieldName, fieldType, zeroValue, false, false))
}

func (t *templateData) addSelectorExpr(x *ast.SelectorExpr, receiverType, fieldName string) {
	if strings.ToLower(fieldName[:1]) == fieldName[:1] { // Non-exported field.
		return
	}

	var xX string
	if xx, ok := x.X.(*ast.Ident); ok {
		xX = xx.String()
	}

	switch xX {
	case "time", "json", "http":
		switch xX {
		case "json":
			t.Imports["encoding/json"] = "encoding/json"
		case "http":
			t.Imports["net/http"] = "net/http"
		default:
			t.Imports[xX] = xX
		}
		fieldType := fmt.Sprintf("%v.%v", xX, x.Sel.Name)
		zeroValue := fmt.Sprintf("%v.%v{}", xX, x.Sel.Name)
		if xX == "time" && x.Sel.Name == "Duration" {
			zeroValue = "0"
		}
		if xX == "json" && x.Sel.Name == "Number" {
			zeroValue = `json.Number("")`
		}
		t.Getters = append(t.Getters, newGetter(receiverType, fieldName, fieldType, zeroValue, false, false))
	default:
		logf("addSelectorExpr: xX %q, type %q, field %q: unknown x=%+v; skipping.", xX, receiverType, fieldName, x)
	}
}

type templateData struct {
	filename string
	Package  string
	Imports  map[string]string
	Getters  []*getter
}

type getter struct {
	sortVal      string // Lower-case version of "ReceiverType.FieldName".
	ReceiverVar  string // The one-letter variable name to match the ReceiverType.
	ReceiverType string
	FieldName    string
	FieldType    string
	ZeroValue    string
	NamedStruct  bool // Getter for named struct.
	Slice        bool // Getter for slices.
}

type byName []*getter

func (b byName) Len() int           { return len(b) }
func (b byName) Less(i, j int) bool { return b[i].sortVal < b[j].sortVal }
func (b byName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

const source = `// Code generated by gen_accessors; DO NOT EDIT.

package {{.Package}}
{{with .Imports}}
import (
  {{- range . -}}
  "{{.}}"
  {{end -}}
)
{{end}}
{{range .Getters}}
{{if .NamedStruct}}
// Get{{.FieldName}} returns the {{.FieldName}} field.
func ({{.ReceiverVar}} *{{.ReceiverType}}) Get{{.FieldName}}() *{{.FieldType}} {
  if {{.ReceiverVar}} != nil {
    return {{.ReceiverVar}}.{{.FieldName}}
  }
  return {{.ZeroValue}}
}
{{else if .Slice}}
// Get{{.FieldName}} returns the {{.FieldName}} field.
func ({{.ReceiverVar}} *{{.ReceiverType}}) Get{{.FieldName}}() {{.FieldType}} {
  if {{.ReceiverVar}} != nil {
    return {{.ReceiverVar}}.{{.FieldName}}
  }
  return {{.ZeroValue}}
}
{{else}}
// Get{{.FieldName}} returns the {{.FieldName}} field if it's non-nil, zero value otherwise.
func ({{.ReceiverVar}} *{{.ReceiverType}}) Get{{.FieldName}}() {{.FieldType}} {
  if {{.ReceiverVar}} != nil && {{.ReceiverVar}}.{{.FieldName}} != nil {
    return *{{.ReceiverVar}}.{{.FieldName}}
  }
  return {{.ZeroValue}}
}
{{end}}
{{end}}
`
