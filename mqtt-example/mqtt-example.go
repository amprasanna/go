package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"code.cloudfoundry.org/cli/plugin"
	"time"
)

type BasicPlugin struct{}

func (c *BasicPlugin) Run(cliConnection plugin.CliConnection, args []string) {

	if args[0] == "wiotp" {
		fmt.Println("Connected to Watson IoT Platform ")
  opts := MQTT.NewClientOptions().AddBroker("tcp://quickstart.messaging.internetofthings.ibmcloud.com:1883")
  opts.SetClientID("d:quickstart:device:112233445566")

  c := MQTT.NewClient(opts)
  if token := c.Connect(); token.Wait() && token.Error() != nil {
    panic(token.Error())
  }


    text := fmt.Sprintf("{ \"d\" : { \"temp\" : 24 }}")
    token := c.Publish("iot-2/evt/stat/fmt/json", 0, false, text)
    token.Wait()


  time.Sleep(3 * time.Second)


  c.Disconnect(250)
	}
}

func (c *BasicPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "MyBasicPlugin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "wiotp",
				HelpText: "Basic plugin command's help text",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "basic-plugin-command\n   cf basic-plugin-command",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(BasicPlugin))
}
