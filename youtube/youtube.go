package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	Kind  string `json:"kind"`
	Items []Item `json:"items"`
}

type Item struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}

func GetSubscribers() (Item, error) {
	var response Response
	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		fmt.Println(err)
		return Item{}, err
	}

	// Query parameters
	q := req.URL.Query()

	q.Add("key", os.Getenv("YOUTUBE_KEY")) // Environment Variable
	q.Add("id", os.Getenv("CHANNEL_ID"))   // Environment Variable
	q.Add("part", "statistics")
	req.URL.RawQuery = q.Encode()

	// Request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Item{}, err
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)

	// read the JSON response
	body, _ := ioutil.ReadAll(resp.Body)
	// Umarshal into an Response Struct
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Item{}, err
	}

	return response.Items[0], nil
}
