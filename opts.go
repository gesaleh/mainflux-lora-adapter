/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 * Fork done by G.SALEH for AgriSmart project
 */



package main

// Options block for gnatsd server.
type Options struct {
	Host          string
	Port          int
	Trace         bool
	Debug         bool
	MaxConn       int
	Logtime       bool
	Authorization string
	Username      string
	Password      string
	PidFile       string
	LogFile       string
}
