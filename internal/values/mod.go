package values

import (
	"fmt"
	"os"

	goYaml "gopkg.in/yaml.v3"
)

type Values map[interface{}]interface{}

// Parse unmarshals yaml string into a generic map data structure.
// For nil values the user is prompted to provide a value (can be left empty).
func Parse(yaml string) (Values, error) {
	values := make(Values)
	err := goYaml.Unmarshal([]byte(yaml), &values)
	for key, val := range values {
		if val == nil {
			var tmp string
			fmt.Printf("Provide value for [%s]: ", key)
			fmt.Scanf("%s", &tmp)
			values[key] = tmp
		}
	}
	return values, err
}

// ParseFile unmarshals yaml file into a generic map data structure.
// For nil values the user is prompted to provide a value (can be left empty).
func ParseFile(valuesFilePath string) (Values, error) {
	valBytes, err := os.ReadFile(valuesFilePath)
	if err != nil {
		return nil, err
	}
	return Parse(string(valBytes))
}
