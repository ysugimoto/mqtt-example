package main

import (
	"fmt"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"time"
)

var f MQTT.MessageHandler = func(client *MQTT.MqttClient, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {

	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://192.168.33.10:1883")
	opts.SetUsername("username")
	opts.SetPassword("password")

	client := MQTT.NewClient(opts)
	if _, err := client.Start(); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	filter, _ := MQTT.NewTopicFilter("example/topic", 0)
	if _, err := client.StartSubscription(f, filter); err != nil {
		panic(err)
	} else {
		println("Wainting message...")
	}

	for {
		time.Sleep(1 * time.Nanosecond)
	}

}
