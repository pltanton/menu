package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/yosssi/ace"
)

const aceRoot = "templates"
const aceBase = "layout"

// RespondAce just wrapper for ace.Load and template.Execute commands query
func RespondAce(w http.ResponseWriter,
	inner string, data interface{}) (tmpl *template.Template, err error) {

	aceOptions := &ace.Options{DynamicReload: true}
	basePath := filepath.Join(aceRoot, aceBase)
	innerPath := filepath.Join(aceRoot, inner)
	tmpl, err = ace.Load(basePath, innerPath, aceOptions)
	if err != nil {
		return
	}

	err = tmpl.Execute(w, data)
	return
}
