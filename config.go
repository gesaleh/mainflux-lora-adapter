package main

// Config struct
type Config struct {
	// HTTP
	HTTPHost string
	HTTPPort int

	// MQTT
	MQTTMainfluxHost string
	MQTTMainfluxPort int

	MQTTLoraHost string
	MQTTLoraPort int

	// LoRa
	LORAChannel string

	// Log
	LogFile string
}

var cfg Config
