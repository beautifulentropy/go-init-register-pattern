package main

import (
	"log"

	"github.com/beautifulentropy/go-init-register-pattern/src/components"
	_ "github.com/beautifulentropy/go-init-register-pattern/src/components/component1"
	_ "github.com/beautifulentropy/go-init-register-pattern/src/components/component2"
	"gopkg.in/yaml.v2"
)

type config struct {
	Kind     string                 `yaml:"kind"`
	Settings map[string]interface{} `yaml:"settings"`
}

func main() {
	// Component1
	// Our "fake" YAML input for Component 1 contains a key: `input`
	// with a string value.
	c1YAML := []byte(`
kind: C1
settings:
  input: Over 9000
`)
	// Unmarshal our YAML to a config struct so we can access `Kind` and
	// use it's value to fetch the right Component type for `Settings`
	// to get unmarshaled to.
	var c1data config
	yaml.Unmarshal(c1YAML, &c1data)

	// Our component was registered as "C1" by the init method in
	// `src/components/component1`. This was called when we performed
	// our blank import on line 9 above. So we just have to fetch C1
	// from the Registry
	c1, err := components.GetComponent(c1data.Kind)
	if err != nil {
		log.Fatal(err)
	}

	// Marshal just the `Settings` field of `config` back to bytes and
	// pass it to our `Component` type to be unmarshaled.
	c1settings, _ := yaml.Marshal(c1data.Settings)
	c1 = c1.UnmarshalSettings(c1settings)
	log.Printf("type: '%T' says: %q", c1, c1.Greet())

	// Component2
	c2YAML := []byte(`
kind: C2
settings:
  input: 9001
`)
	var c2data config
	yaml.Unmarshal(c2YAML, &c2data)

	c2, err := components.GetComponent(c2data.Kind)
	if err != nil {
		log.Fatal(err)
	}

	c2settings, _ := yaml.Marshal(c2data.Settings)
	c2 = c2.UnmarshalSettings(c2settings)
	log.Printf("type: '%T' says: %q", c2, c2.Greet())

	// Component3
	c3YAML := []byte(`
kind: C3
settings:
  input: Over 10000
`)
	var c3data config
	yaml.Unmarshal(c3YAML, &c3data)

	// There is no C3 so this should fail.
	c3, err := components.GetComponent(c3data.Kind)
	if err != nil {
		log.Fatalln(err)
	}

	c3settings, _ := yaml.Marshal(c3data.Settings)
	c3 = c3.UnmarshalSettings(c3settings)
	log.Printf("type: '%T' says: %q", c3, c3.Greet())
}
