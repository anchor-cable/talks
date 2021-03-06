package main

import (
	"go/build"
	"os"
	"strings"
	"text/template"
)

var tmpl = template.Must(template.New("").Parse(`// Don't edit! generated by gen.go

package {{.Package}}

{{range .Params}}func Handle{{title .Name}}(w http.ResponseWriter, r *http.Request) {
	p := r.FormValue("{{.Name}}")
	if p == "" {
		http.Error(w, "{{.Name}} not specified", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%s is god.", p)
}
{{end}}
`))

func Generate(name string...) error {
	pkg, err := build.Default.Import(".", 0)
	if err != nil {
		return err
	}
	f, err := os.Create(fmt.Sprintf("%s_gen.go"), strings.ToLower(name))
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.Execute(f, map[string]string{})
}
