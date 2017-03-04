# Mainflux CLI

[![License](https://img.shields.io/badge/license-Apache%20v2.0-blue.svg)](LICENSE)
[![Build Status](https://travis-ci.org/mainflux/mainflux-cli.svg?branch=master)](https://travis-ci.org/mainflux/mainflux-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/Mainflux/mainflux-cli)](https://goreportcard.com/report/github.com/Mainflux/mainflux-cli)
[![Join the chat at https://gitter.im/Mainflux/mainflux](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/Mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Adapter between Mainflux IoT system and [LoRa Server](https://github.com/brocaar/loraserver).

This adapter sits between Mainflux and LoRa server and just forwards the messages form one system to another via MQTT protocol, using the adequate MQTT topics and in the good message format (JSON and SenML), i.e. respecting the APIs of both systems. 

LoRa Server is used for connectivity layer and data is pushed via this adapter service to Mainflux, where it is persisted and routed to other protocols via Mainflux multi-protocol message broker. Mainflux adds user accounts, application management and security in order to obtain the overall end-to-end LoRa solution.

### Installation
#### Prerequisite
If not set already, please set your `GOPATH` and `GOBIN` environment variables. For example:
```bash
mkdir -p ~/go
export GOPATH=~/go
export GOBIN=$GOPATH/bin
# It's often useful to add $GOBIN to $PATH
export PATH=$PATH:$GOBIN
```

#### Get the code
Use [`go`](https://golang.org/cmd/go/) tool to "get" (i.e. fetch and build) `mainflux-lora-adapter` package:
```bash
go get github.com/mainflux/mainflux-lora-adapter
```

This will download the code to `$GOPATH/src/github.com/mainflux/mainflux-lora-adapter` directory,
and then compile it and install the binary in `$GOBIN` directory.

Now you can run the program with:
```
mainflux-cli
```
if `$GOBIN` is in `$PATH` (otherwise use `$GOBIN/mainflux-lora-adapter`)

### Documentation
Development documentation can be found [here](http://mainflux.io/).

### Community
#### Mailing lists
[mainflux](https://groups.google.com/forum/#!forum/mainflux) Google group.

#### IRC
[Mainflux Gitter](https://gitter.im/Mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

#### Twitter
[@mainflux](https://twitter.com/mainflux)

### License
[Apache License, version 2.0](LICENSE)
