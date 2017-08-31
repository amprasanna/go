package main

import (
	"encoding/json"
	"fmt"
	"code.cloudfoundry.org/cli/plugin"
	"log"
	"net/http"
	//"net/url"
	"io/ioutil"
	"time"
)

//User defines model for storing account details in database
type User struct {
	Username string
	Password string `json:"-"`
	IsAdmin bool
	CreatedAt time.Time
}

type BasicPlugin struct{}

func (c *BasicPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	user := User{} //initialize empty user
	
	if args[0] == "wiotp-rest-call" {
	fmt.Println("Connected to Watson IoT Platform ")		
	url := fmt.Sprintf("https://a-f71n0g-umjfbv8lea:sOwgIrKcCmZGHoY9c6@f71n0g.internetofthings.ibmcloud.com/api/v0002/device/types/")

	fmt.Println("--------------Printing the URL---------------")
	fmt.Println(url)
		
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	
	client := &http.Client{}
		
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()
	
	if resp.StatusCode == 200 { // OK
    bodyBytes, err  := ioutil.ReadAll(resp.Body)
	if err != nil {
	fmt.Println("--------------Printing the ERR---------------")
	fmt.Println(err)
	}
    bodyString := string(bodyBytes)
	fmt.Println("About to print the response bodyString")
	fmt.Println("--------------Printing the Response Body---------------")
	fmt.Println(bodyString)
	fmt.Println("--------------Parsing the Response Body---------------")
	fmt.Println("--------------Work in Progress---------------")
}
	
	} else {
        fmt.Println("Execute 'cf wiotp-rest' to continue with the execution")
    }
}

func (c *BasicPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "WIoTP--REST",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "wiotp-rest-call",
				HelpText: "Basic plugin command's help text",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "wiotp-rest-call\n   cf wiotp-rest-call",
				},
			},
		},
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request){
	user := User{} //initialize empty user

	//Parse json request body and use it to set fields on user
	//Note that user is passed as a pointer variable so that it's fields can be modified
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		panic(err)
	}

	//Set CreatedAt field on user to current local time
	user.CreatedAt = time.Now().Local()

	//Marshal or convert user object back to json and write to response 
	userJson, err := json.Marshal(user)
	if err != nil{
		panic(err)
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response 
	w.Write(userJson)
	fmt.Println(userJson)
}

func main() {
	plugin.Start(new(BasicPlugin))
}
