package main

import (
	"github.com/obengwilliam/go-cms/cms"
	"os"
)

func main() {
	var data = cms.Page{
		Title:   "the boy who lost his Dad",
		Content: "Coming over age, is never an easy thing, but it takes time and pressure."}
	cms.Tmpl.ExecuteTemplate(os.Stdout, "index", data)
}
