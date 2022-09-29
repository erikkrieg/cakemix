package values

import (
	goYaml "gopkg.in/yaml.v3"
)

type Values map[interface{}]interface{}

// Unmarshals values into a generic map data structure.
func Parse(yaml string) (Values, error) {
	values := make(Values)
	err := goYaml.Unmarshal([]byte(yaml), &values)
	return values, err
}
