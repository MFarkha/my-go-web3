package render

import (
	"fmt"
	"net/http"
	"text/template"
	"web3/models"
)

var tmplCache = make(map[string]*template.Template)

func makeTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tmplCache[t] = tmpl
	return nil
}

func RenderTemplate(w http.ResponseWriter, t string, pd *models.PageData) {
	var tmpl *template.Template
	var err error
	_, inMap := tmplCache[t]
	if !inMap {
		err = makeTemplateCache(t)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Template is in cache")
	}
	tmpl = tmplCache[t]
	err = tmpl.Execute(w, pd)
	if err != nil {
		fmt.Println(err)
	}
}
