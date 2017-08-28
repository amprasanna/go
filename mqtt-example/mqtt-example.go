package main

import (
  "fmt"
  //import the Paho Go MQTT library
  MQTT "github.com/eclipse/paho.mqtt.golang"
  "time"
)

func main() {
  //create a ClientOptions struct setting the broker address, clientid, turn
  //off trace output and set the default message handler
  opts := MQTT.NewClientOptions().AddBroker("tcp://quickstart.messaging.internetofthings.ibmcloud.com:1883")
  opts.SetClientID("d:quickstart:device:112233445566")

  //create and start a client using the above ClientOptions
  c := MQTT.NewClient(opts)
  if token := c.Connect(); token.Wait() && token.Error() != nil {
    panic(token.Error())
  }


    text := fmt.Sprintf("{ \"d\" : { \"temp\" : 44 }}")
    token := c.Publish("iot-2/evt/stat/fmt/json", 0, false, text)
    token.Wait()


  time.Sleep(3 * time.Second)


  c.Disconnect(250)
}
