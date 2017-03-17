
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
	"strings"

	"github.com/fatih/color"
	"github.com/loraAdapter/api"
	"github.com/loraAdapter/config"
)


var usageStr = `
Usage: mainflux [options]
Adapter Options:
    -c, --config <file>              Configuration file
Logging Options:
    -l, --log <file>                 File to redirect log output
    -T, --logtime                    Timestamp log entries (default: true)
    -D, --debug                      Enable debugging output
    -V, --trace                      Trace the raw protocol
    -DV                              Debug and trace
Common Options:
    -h, --help                       Show this message
    -v, --version                    Show version
`

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
	// Server Options
	opts := Options{}

	var showVersion bool
	var debugAndTrace bool
	var configFile string

	// Parse flags

	flag.BoolVar(&opts.Debug, "D", false, "Enable Debug logging.")
	flag.BoolVar(&opts.Debug, "debug", false, "Enable Debug logging.")
	flag.BoolVar(&opts.Trace, "V", false, "Enable Trace logging.")
	flag.BoolVar(&opts.Trace, "trace", false, "Enable Trace logging.")
	flag.BoolVar(&debugAndTrace, "DV", false, "Enable Debug and Trace logging.")
	flag.BoolVar(&opts.Logtime, "T", true, "Timestamp log entries.")
	flag.BoolVar(&opts.Logtime, "logtime", true, "Timestamp log entries.")
	flag.StringVar(&configFile, "c", "", "Configuration file.")
	flag.StringVar(&configFile, "config", "", "Configuration file.")
	flag.StringVar(&opts.LogFile, "l", "", "File to store logging output.")
	flag.StringVar(&opts.LogFile, "log", "", "File to store logging output.")
	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")

	flag.Usage = usage

	flag.Parse()

	// Show version and exit
	if showVersion {
		PrintServerAndExit()
	}

	// One flag can set multiple options.
	if debugAndTrace {
		opts.Trace, opts.Debug = true, true
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

	// Parse config
	var cfg config.Config
	cfg.Parse()

	// Print banner
	color.Cyan(banner)
	color.Cyan(fmt.Sprintf("MainFlux Lora Server Adapter is running %s:%d-%s:d", cfg.MQTTTargetHost, cfg.MQTTTargetPort,cfg.MQTTHost, cfg.MQTTPort ))

	// mqttClient 
	mqttTargetHost := fmt.Sprintf("tcp://%s:%d", cfg.MQTTTargetHost, cfg.MQTTTargetPort)
	mqttHost := fmt.Sprintf("tcp://%s:%d", cfg.MQTTHost, cfg.MQTTPort)
	api.NewTargetBackend(mqttTargetHost, "", "")
	api.NewBackend(mqttHost, "", "")
	//if err != nil {
	//	log.Fatalf("could not setup mqtt backend: %s", err)
	//}
	//defer Pubsub.Close()

}

var banner = `MAINFLUX LORASERVER ADAPTER  `

