package fronius

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	BaseUrl    string
	HttpClient *http.Client
}

func (c *Client) Get3PInverterData() (data InverterData, err error) {
	uri, err := url.Parse(fmt.Sprintf("%v/solar_api/v1/GetInverterRealtimeData.cgi", c.BaseUrl))
	if err != nil {
		return data, err
	}

	uri.RawQuery = url.Values{
		"Scope":          []string{"Device"},
		"DataCollection": []string{"3PInverterData"},
		"DeviceId":       []string{"1"},
	}.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return data, err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return data, err
	}

	var packet InverterDataPacket

	dec := json.NewDecoder(res.Body)
	if err = dec.Decode(&packet); err != nil {
		return data, err
	}

	return packet.Body.Data, nil
}

func (c *Client) GetCommonInverterData() (data CommonInverterData, err error) {
	uri, err := url.Parse(fmt.Sprintf("%v/solar_api/v1/GetInverterRealtimeData.cgi", c.BaseUrl))
	if err != nil {
		return data, err
	}

	uri.RawQuery = url.Values{
		"Scope":          []string{"Device"},
		"DataCollection": []string{"CommonInverterData"},
		"DeviceId":       []string{"1"},
	}.Encode()

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return data, err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return data, err
	}

	var packet CommonInverterDataPacket

	dec := json.NewDecoder(res.Body)
	if err = dec.Decode(&packet); err != nil {
		return data, err
	}

	return packet.Body.Data, nil
}
