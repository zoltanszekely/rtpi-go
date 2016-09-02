package rtpi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/zoltanszekely/rtpi-go/rtpi/types"
)

type Client struct{}

func (client *Client) getURL(serverURL string) (u *url.URL) {
	u, err := url.Parse(serverURL)
	if err != nil {
		panic(fmt.Sprintf("Error parsing server URL: %s", serverURL))
	}
	return
}

func (client *Client) parse(response *http.Response) (data *interface{}) {
	if response.StatusCode != 200 {
		panic(fmt.Sprintf("Unexpected HTTP status: %d", response.StatusCode))
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(fmt.Sprintf("Error reading response body: %s", err))
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(fmt.Sprintf("Error parsing response body: %s", err))
	}

	return
}

func (client *Client) fetch(u *url.URL) *interface{} {
	urlString := u.String()
	fmt.Printf("downloading from %s\n", urlString)

	httpClient := http.Client{}
	response, err := httpClient.Get(urlString)
	if err != nil {
		panic(fmt.Sprintf("Error downloading from: %s", urlString))
	}
	defer response.Body.Close()

	return client.parse(response)
}

func (client *Client) GetBusStopInformation() (*types.BusStopInformation, error) {
	u := client.getURL("https://data.dublinked.ie/cgi-bin/rtpi/busstopinformation?format=json")
	q := u.Query()
	q.Set("stopid", "")
	u.RawQuery = q.Encode()

	data := client.fetch(u)

	return types.NewBusStopInformation(data)
}

func (client *Client) GetRealTimeBusInformation(stopID string) (*types.RealTimeBusInformation, error) {
	u := client.getURL("https://data.dublinked.ie/cgi-bin/rtpi/realtimebusinformation?format=json")
	q := u.Query()
	q.Set("stopid", stopID)
	u.RawQuery = q.Encode()

	data := client.fetch(u)

	return types.NewRealTimeBusInformation(data)
}
