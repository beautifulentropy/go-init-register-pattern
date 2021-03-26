package components

import (
	"fmt"

	"github.com/beautifulentropy/go-init-register-pattern/src/components"
	"gopkg.in/yaml.v2"
)

type Component2 struct {
	Input int `yaml:"input"`
}

func (c Component2) Greet() string {
	return fmt.Sprintf("Yeah I'm the other one and the input you provided to me was: %d", c.Input)
}

func (c Component2) UnmarshalSettings(settings []byte) components.Component {
	var c2 Component2
	yaml.Unmarshal(settings, &c2)
	return c2
}

func init() {
	components.Register("C2", Component2{})
}
