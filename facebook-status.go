package facebook-status

import (
	"fmt"
	"encoding/json"
	"github.com/heatxsink/go-httprequest"
)

type Current struct {
	Health int `json:"health"`
	Subject string `json:"subject"`
}

type Push struct {
	Status string `json:"status"`
	Updated string `json:"updated"`
	Id int `json:"id"`
}

type ApiStatusResponse struct {
	Push Push `json:"push"`
	Current Current `json:"current"`
}

func ApiStatus() (ApiStatusResponse, error) {
	url := "http://www.facebook.com/feeds/api_status.php"
	hr := httprequest.NewWithDefaults()
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11"
	body, status_code, err := hr.Get(url, headers)
	var response ApiStatusResponse
	if err == nil {
		if status_code == 200 {
			err := json.Unmarshal(body, &response)
			if err != nil {
				fmt.Errorf("ERROR: facebook.ApiStatus json.Unmarshal() %v", err)
				return response, err
			}
		}
	}
	return response, err
}
