package casting

import (
	"os"

	tpl "text/template"

	"github.com/erikkrieg/cast/internal/values"
)

func Render(values values.Values, template string, destination string) error {
	t, err := tpl.New("cast").Parse(template)
	if err != nil {
		return err
	}
	newFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	return t.Execute(newFile, values)
}
