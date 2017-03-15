package api

import (
        "encoding/json"
        "encoding/base64"
        "fmt"
        "sync"
        "time"
        log "github.com/Sirupsen/logrus"
        "github.com/eclipse/paho.mqtt.golang"
        "github.com/loraAdapter/models"
		"github.com/loraAdapter/config"
)

// Backend implements a MQTT pub-sub backend.
type Backend struct {
        conn         mqtt.Client
        mutex        sync.RWMutex
}

var (
        Pubsub *Backend
)

// NewBackend creates a new Backend.
func NewBackend(server, username, password string) (*Backend, error) {
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
        SubLora(&b)
        Pubsub = &b
        return &b, nil
}

func NewTargetBackend(server, username, password string) (*Backend, error) {
        t := Backend{}
        opts := mqtt.NewClientOptions()
        opts.AddBroker(server)
        opts.SetUsername(username)
        opts.SetPassword(password)
        opts.SetOnConnectHandler(t.onConnected)
        opts.SetConnectionLostHandler(t.onConnectionLost)

        log.WithField("server", server).Info("target backend: connecting to mqtt broker")
        t.conn = mqtt.NewClient(opts)
        if token := t.conn.Connect(); token.Wait() && token.Error() != nil {
                return nil, token.Error()
        }
        return &t, nil
}

//callback to publish on target server mainFlux
func (t *Backend) SendMQTTTargetMsg(topic string, data []byte) error {
        log.WithField("topic", topic).Info("backend publishing packet: ", string(data))
        if token := t.conn.Publish(topic, 0, false, data); token.Wait() && token.Error() != nil {
                return token.Error()
        }
        return nil
}

//callback to Listen on source server loraserver
func (b *Backend) SendMQTTMsg(topic string, data []byte) error {
        log.WithField("topic", topic).Info("backend publishing packet: ", string(data))
        if token := b.conn.Publish(topic, 0, false, data); token.Wait() && token.Error() != nil {
                return token.Error()
        }
        return nil
}

// Close closes the backend.
func (b *Backend) Close() {
        b.conn.Disconnect(250)  // wait 250 milisec to complete pending actions
        log.Info("-- DISCONNECTING\n")
}

func (t *Backend) LoraHandler(client *mqtt.Client, msg mqtt.Message) {
        u := models.Message{}
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
		topic = cfg.LORAChannel
        t.SendMQTTMsg(topic, data)
        log.Info(" --> PUSH DATA: %s to %s\n", topic, data)
}

// Subscribe to lora server gateways messages 
func SubLora(b *Backend) {
        if receipt := b.conn.Subscribe("application/+/node/+/rx", 0, b.MessageHandler); receipt.Wait() && receipt.Error() != nil {
                log.Info("Failed to subscribe, err: %v\n", receipt.Error())
        } else {
                log.Info("Subscription Done ...")
                for b.conn.IsConnected() {
                        time.Sleep(60*time.Second)
                        //fmt.Println(time.Now().String())
                }
        }
}

// Handler for received messages from loraserver
func (b *Backend) MessageHandler(c mqtt.Client, msg mqtt.Message) {
        log.WithField("topic", msg.Topic()).Info("backend: packet received")
        log.Info("TOPIC: %s\n", msg.Topic())
        log.Info("MSG: %s\n", msg.Payload())
        b.LoraHandler(&c, msg)
}

func (b *Backend) onConnected(c mqtt.Client) {
        defer b.mutex.RUnlock()
        b.mutex.RLock()
        log.Info("backend: connected to mqtt broker")
}

func (b *Backend) onConnectionLost(c mqtt.Client, reason error) {
        log.Errorf("backend: mqtt connection error: %s", reason)
}

