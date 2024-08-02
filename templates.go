package webutil

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func LoadTemplates(root string) (*template.Template, error) {
	var tmpl *template.Template
	err := filepath.WalkDir(root, func (path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			p, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			p = strings.Replace(p, "\\", "/", -1)
			if tmpl == nil {
				tmpl = template.New(p)
			} else {
				tmpl = tmpl.New(p)
			}
			bytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			_, err = tmpl.Parse(string(bytes))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
