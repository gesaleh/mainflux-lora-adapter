
package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

// Config struct
type Config struct {
	// HTTP
	HTTPHost string
	HTTPPort int

	// MQTT
	MQTTHost string
	MQTTPort int


	MQTTTargetHost string
	MQTTTargetPort int

	LORAChannel string
}

// Parse TOML config
func (cfg *Config) Parse() {

	var confFile string

	testEnv := os.Getenv("TEST_ENV")
	if testEnv == "" && len(os.Args) > 1 {
		// We are not in the TEST_ENV (where different args are provided)
		// and provided config file as an argument
		confFile = os.Args[1]
	} else {
		// default cfg path to source dir, as we keep cfg.yml there
		confFile = os.Getenv("GOPATH") + "/src/github.com/loraAdapter/config/config.toml"
	}

	if _, err := toml.DecodeFile(confFile, &cfg); err != nil {
		// handle error
		fmt.Println("Error parsing Toml")
	}
}
