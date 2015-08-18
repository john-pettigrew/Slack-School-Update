package main

import (
	"net/http"
	"encoding/json"
	"math/rand"
	"net/url"
	"io/ioutil"
	"time"
	"log"
)

type Settings struct{
	Token string `json:"token"`
	Channel string
	GoodMessages []string
	BadMessages []string
	DaysInClass []string
	DaysAtWork []string
}

func main(){
	
	//Read and convert settings file
	fileData, err := ioutil.ReadFile("settings.json")
	if(err != nil){
		log.Fatal("Error reading file")			
	}
	var settings Settings
	
	err = json.Unmarshal(fileData, &settings)
	if(err != nil){
		log.Fatal("Error unmarshaling JSON")
	}

	//See what type of day it is

	now := time.Now()

	todayIsClass := false
	todayIsWork := false

	for _, day := range settings.DaysInClass{
		if now.Weekday().String() == day {
			todayIsClass = true
		}
	}

	for _, day := range settings.DaysAtWork{
		if now.Weekday().String() == day {
			todayIsWork = true
		}
	}

	//Get random message

	todaysMessage := ""
	rand.Seed(now.Unix())

	if todayIsWork {

		todaysMessage = settings.GoodMessages[rand.Intn(len(settings.GoodMessages))]
	}else if todayIsClass{

		todaysMessage = settings.BadMessages[rand.Intn(len(settings.BadMessages))]

	}


	if todaysMessage != "" {

		//Send message	
		_, err := http.Post("https://slack.com/api/chat.postMessage?token=" + settings.Token + "&channel="+settings.Channel + "&text=" + url.QueryEscape(todaysMessage), "application/json", nil)
		if(err != nil){
			log.Fatal("Error sending message")	
		}
	}
	
}
