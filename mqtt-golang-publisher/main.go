package main

import (
	"fmt"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"time"
)

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

	i := 0

	for {
		timer := time.NewTimer(time.Second)
		<-timer.C

		i++

		fmt.Printf("Ticker executed. Run times: %d\n", i)

		message := MQTT.NewMessage([]byte(fmt.Sprintf("Hello paho, ticker:%d", i)))
		message.SetQoS(MQTT.QOS_ZERO)

		receipt := client.PublishMessage("example/topic", message)
		<-receipt
	}

}
