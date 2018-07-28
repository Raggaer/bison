package main

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func loadTemplates() (*template.Template, error) {
	tpl := template.New("bison")
	err := filepath.Walk("views", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), ".html") {
			if _, err := tpl.ParseFiles(path); err != nil {
				return err
			}
		}
		return nil
	})
	return tpl, err
}
