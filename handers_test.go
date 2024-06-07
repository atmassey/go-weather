package main

import (
	"fmt"
	"html/template"
	"os"
	"testing"
)

func TestParseTemplates(t *testing.T) {
	tmpl, err := template.ParseFS(os.DirFS("./web"), "template.html")
	if err != nil {
		t.Error("Error parsing template")
	}
	fmt.Printf("%v\n", tmpl)
}
