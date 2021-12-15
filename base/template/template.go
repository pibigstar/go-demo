package template

import (
	"bytes"
	"github.com/pkg/errors"
	"go/format"
	"text/template"
)

// tpl 生成代码需要用到模板
const tpl = `
package {{.pkg}}

// messages get msg from const comment
var messages = map[string]string{
	{{range $key, $value := .comments}}
	{{$key}}: "{{$value}}",{{end}}
}

// GetErrMsg get error msg
func GetErrMsg(code string) string {
	if msg, ok := messages[code]; ok {
		return msg
	}
	return ""
}
`

// gen 生成代码
func gen(comments map[string]string) ([]byte, error) {
	var buf = bytes.NewBufferString("")

	data := map[string]interface{}{
		"pkg":      "main",
		"comments": comments,
	}

	t, err := template.New("").Parse(tpl)
	if err != nil {
		return nil, errors.Wrapf(err, "template init err")
	}

	err = t.Execute(buf, data)
	if err != nil {
		return nil, errors.Wrapf(err, "template data err")
	}

	return format.Source(buf.Bytes())
}
