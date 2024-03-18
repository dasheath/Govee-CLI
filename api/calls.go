package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetDevices() GetDeviceResponse {
	req, err := http.NewRequest(
		http.MethodGet,
		GoveeBaseUrl+"/user/devices",
		nil,
	)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Govee-API-Key", os.Getenv("Govee_API_Key"))

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

func SetDevicePower(deviceId string, powerValue int) {

	powerDeviceBody := PowerDeviceBody{
		RequestId: "uuid",
		Payload: PowerDevicePayload{
			SKU:    BulbSKU,
			Device: deviceId,
			Capability: PowerDeviceCapability{
				CapabilityType: "devices.capabilities.on_off",
				Instance:       "powerSwitch",
				Value:          powerValue,
			},
		},
	}

	jsonBody, err := json.Marshal(powerDeviceBody)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		GoveeBaseUrl+"/device/control",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Govee-API-Key", os.Getenv("Govee_API_Key"))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode == 400 {
		panic("Bad Request!")
	} else if res.StatusCode != 200 {
		fmt.Printf("%v\n", res.StatusCode)
		panic("API call for power setting failed")
	} else {
		fmt.Printf("Call Succeeded for device: %v\n", deviceId)
	}
}
