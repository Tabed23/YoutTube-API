package main

import (
	"YouTube-API/youtube"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	args()
}
func usage() {
	fmt.Printf("plz enter channel id \n -h  help \n")
}
func help() {
	fmt.Println("enter the id of the channel \n{\ngo to the channel from you want to get data \nclick on the url\nyou will see a URL like this https://www.youtube.com/channel/ some id \nclick on the id after the slash and past it in the terminal\n}")
}
func args() {
	cmd := os.Args[1]
	if cmd == "-h" {
		help()
	} else if cmd == "-id" {
		getData()
	} else {
		usage()
		os.Exit(1)
	}
}
func getData() {
	channelID := os.Args[2]
	item, err := youtube.GetChannelDetail(channelID)
	if err != nil {
		log.Fatalf("error %v", err)
	}
	prettyJSON, _ := json.MarshalIndent(item, "", "\t")
	ioutil.WriteFile("test.json", prettyJSON, 0644)
}
