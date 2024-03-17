package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func makeRequest(method string, endpoint string, jsonBody io.Reader) *http.Request {
	// Would be handy to send a generic request to the Govee api
	// with one function
	//
	// OR at least return a

	req, err := http.NewRequest(
		method,
		GoveeBaseUrl+endpoint,
		jsonBody,
	)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Govee-API-Key", os.Getenv("Govee_API_Key"))

	return req

}

func GetDevices() GetDeviceResponse {
	req := makeRequest(http.MethodGet, "/user/devices", nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("API call for reading device info failed")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var response GetDeviceResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	return response
}
