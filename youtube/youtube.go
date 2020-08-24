package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}
type Item struct {
	Kind       string     `json:"kind"`
	Id         string     `json:"id"`
	Statistics statistics `json:"statistics"`
	Snippet    snippet    `json:"snippet"`
	Status     status     `json:"status"`
}
type statistics struct {
	ViewCount       string `json:"viewcount"`
	SubscriberCount string `json:"subscribercount"`
	VideoCount      string `json:"videocount"`
}
type snippet struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"publishedAt"`
	Country     string    `json:"country"`
}
type status struct {
	PrivacyStatus string `json:"privacyStatus"`
	MadeForKids   bool   `json:"madeForKids"`
}

func GetChannelDetail(channel_Id string) (Item, error) {
	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		fmt.Println(" request Error ", err)
		return Item{}, err
	}
	query := req.URL.Query()
	query.Add("key", os.Getenv("YOUTUBE_KEY"))
	query.Add("id", channel_Id)
	query.Add("part", "statistics")
	query.Add("part", "snippet")
	query.Add("part", "status")
	req.URL.RawQuery = query.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client error")
		log.Fatal("client Error ", err.Error())
		return Item{}, err
	}
	fmt.Println("Status ", resp.Status)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("body reader Error", err.Error())
		return Item{}, err
	}
	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		log.Fatal("Unmarshal Error ", err.Error())
		return Item{}, err
	}

	return response.Items[0], nil
}
