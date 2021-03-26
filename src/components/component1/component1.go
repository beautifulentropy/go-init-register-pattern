package components

import (
	"fmt"

	"github.com/beautifulentropy/go-init-register-pattern/src/components"
	"gopkg.in/yaml.v2"
)

type Component1 struct {
	Input string `yaml:"input"`
}

func (c Component1) Greet() string {
	return fmt.Sprintf(
		"Hello, the input you provided to me was: %s", c.Input)
}

func (c Component1) UnmarshalSettings(settings []byte) components.Component {
	var c1 Component1
	yaml.Unmarshal(settings, &c1)
	return c1
}

func init() {
	components.Register("C1", Component1{})
}
