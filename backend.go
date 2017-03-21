package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/eclipse/paho.mqtt.golang"
	"sync"
)

// Backend implements a MQTT pub-sub backend.
type Backend struct {
	conn mqtt.Client
	// Are we connecting to LoRa Server or to Mainflux
	isLora bool
	mutex  sync.RWMutex
}

var (
	loraBackend     *Backend
	mainfluxBackend *Backend
)

// NewBackend creates a new Backend.
func NewBackend(server, username, password string, isLora bool) (*Backend, error) {
	b := Backend{}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(server)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetOnConnectHandler(b.onConnected)
	opts.SetConnectionLostHandler(b.onConnectionLost)

	log.WithField("server", server).Info("backend: connecting to mqtt broker")
	b.conn = mqtt.NewClient(opts)
	if token := b.conn.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	b.isLora = isLora

	return &b, nil
}

// Send MQTT message
func (b *Backend) SendMQTTMsg(topic string, data []byte) error {
	log.WithField("topic", topic).Info("backend publishing packet: ", string(data))
	if token := b.conn.Publish(topic, 0, false, data); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Close closes the backend.
func (b *Backend) Close() {
	b.conn.Disconnect(250) // wait 250 milisec to complete pending actions
	log.Info("-- DISCONNECTING\n")
}

// Subscribe to lora server messages
func (b *Backend) Sub() error {
	switch b.isLora {
	case true:
		if s := b.conn.Subscribe("application/+/node/+/rx", 0, b.MessageHandler); s.Wait() && s.Error() != nil {
			log.Info("Failed to subscribe, err: %v\n", s.Error())
			return s.Error()
		}
	case false:
		// For now we deo not SUB to Mainflux
		break
	}

	return nil
}

// Handler for received messages from loraserver
func (b *Backend) MessageHandler(c mqtt.Client, msg mqtt.Message) {
	log.WithField("topic", msg.Topic()).Info("backend: packet received")
	log.Info("TOPIC: %s\n", msg.Topic())
	log.Info("MSG: %s\n", msg.Payload())

	switch b.isLora {
	case true:
		// Mainflux backend is subscribed to LoRa Network Server and recieves LoRa messages
		u := LoraMessage{}
		errStatus := json.Unmarshal(msg.Payload(), &u)
		if errStatus != nil {
			log.Errorf("\nerror: decode json failed")
			log.Errorf(errStatus.Error())
			return
		}

		fmt.Printf("\n <-- RCVD DATA: %s\n", u.Data)
		data, err := base64.StdEncoding.DecodeString(u.Data)
		if err != nil {
			log.Errorf("\nerror: decode base64 failed")
		}

		topic := cfg.LORAChannel
		mainfluxBackend.SendMQTTMsg(topic, data)
		log.Info(" --> PUSH DATA: %s to %s\n", topic, data)
	case false:
		// LoRa backend is not currently subsctibed to Mainflux MQTT broker
		break
	}

}

func (b *Backend) onConnected(c mqtt.Client) {
	defer b.mutex.RUnlock()
	b.mutex.RLock()
	log.Info("backend: connected to mqtt broker")
}

func (b *Backend) onConnectionLost(c mqtt.Client, reason error) {
	log.Errorf("backend: mqtt connection error: %s", reason)
}
