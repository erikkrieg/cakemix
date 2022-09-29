package values

import (
	"os"

	goYaml "gopkg.in/yaml.v3"
)

type Values map[interface{}]interface{}

// Unmarshals yaml string into a generic map data structure.
func Parse(yaml string) (Values, error) {
	values := make(Values)
	err := goYaml.Unmarshal([]byte(yaml), &values)
	return values, err
}

// Unmarshals yaml file into a generic map data structure.
func ParseFile(valuesFilePath string) (Values, error) {
	valBytes, err := os.ReadFile(valuesFilePath)
	if err != nil {
		return nil, err
	}
	return Parse(string(valBytes))
}
