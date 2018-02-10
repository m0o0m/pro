package render

import (
	"fmt"
	"global"
	"io/ioutil"

	"html/template"

	"path/filepath"
)

func parseFiles(filenames ...string) (*template.Template, error) {
	if len(filenames) == 0 {
		// Not really a problem, but be consistent.
		return nil, fmt.Errorf("html/template: no files named in call to ParseFiles")
	}
	var t *template.Template
	for _, filename := range filenames {
		var b []byte
		var err error
		if MongoCacheSwitch {
			b, err = global.ReadHTML(filename)
		} else {
			b, err = ioutil.ReadFile(filename)
		}
		if err != nil {
			global.GlobalLogger.Error("err:%s <%s>", err.Error(), filename)
			return nil, err
		}
		s := string(b)
		name := filepath.Base(filename)
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
