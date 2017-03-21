/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 * Fork done for smartAgri project G.SALEH
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

var usageStr = `
Usage: maiflux-lora-adapter [options]
Logging Options:
    -l, --log <file>                 File to redirect log output
Common Options:
    -h, --help                       Show this message
    -v, --version                    Show version
`

const (
	Version string = "0.1"
)

// usage will print out the flag options for the server.
func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

// PrintServerAndExit will print our version and exit.
func PrintServerAndExit() {
	fmt.Printf("mainflux-lora-adapter version %s\n", Version)
	os.Exit(0)
}

func main() {

	// Config
	cfg = Config{
		MQTTMainfluxHost: "0.0.0.0",
		MQTTMainfluxPort: 1883,
		MQTTLoraHost:     "0.0.0.0",
		MQTTLoraPort:     1883,
	}

	var showVersion bool

	// Parse flags
	flag.StringVar(&cfg.LogFile, "l", "", "File to store logging output.")
	flag.StringVar(&cfg.LogFile, "log", "", "File to store logging output.")
	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")

	flag.Usage = usage

	flag.Parse()

	// Show version and exit
	if showVersion {
		PrintServerAndExit()
	}

	// Process args looking for non-flag options,
	// 'version' and 'help' only for now
	for _, arg := range flag.Args() {
		switch strings.ToLower(arg) {
		case "version":
			PrintServerAndExit()
		case "help":
			usage()
		}
	}

	// Print banner
	color.Cyan(banner)
	color.Cyan(fmt.Sprintf("MainFlux Lora Server Adapter is running %s:%d-%s:%d", cfg.MQTTMainfluxHost, cfg.MQTTMainfluxPort, cfg.MQTTLoraHost, cfg.MQTTLoraPort))

	// mqttClient
	mainfluxMqttAddr := fmt.Sprintf("tcp://%s:%d", cfg.MQTTMainfluxHost, cfg.MQTTMainfluxPort)
	loraMqttAddr := fmt.Sprintf("tcp://%s:%d", cfg.MQTTLoraHost, cfg.MQTTLoraPort)

	// Create backends that connect as MQTT clients to brokers of Mainflux and LoRa Server
	var e error
	if mainfluxBackend, e = NewBackend(mainfluxMqttAddr, "", "", false); e != nil {
		println("Cannot create the Mainflux backend")
	}

	if loraBackend, e = NewBackend(loraMqttAddr, "", "", true); e != nil {
		println("Cannot create LoRa Server backend")
	}

	// Subscribe LoRa backend to LoRa Network Server topic
	if err := loraBackend.Sub(); err != nil {
		println("Cannot subsribe to LoRa Network Server")
	}

	runtime.Goexit()

	defer mainfluxBackend.Close()
	defer loraBackend.Close()
}

var banner = `MAINFLUX LORASERVER ADAPTER  `
