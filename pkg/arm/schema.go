package arm

import (
	"fmt"

	packr "github.com/gobuffalo/packr/v2"
)

func (m *Mixin) PrintSchema() error {
	schema, err := m.GetSchema()
	if err != nil {
		return err
	}

	fmt.Fprint(m.Out, schema)

	return nil
}

func (m *Mixin) GetSchema() (string, error) {
	t := packr.New("schema", "./schema")

	b, err := t.Find("arm.json")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
