package api

const (
	GoveeBaseUrl string = "https://openapi.api.govee.com/router/api/v1"
	BulbSKU      string = "H6008"
)

type GetDeviceResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []Device `json:"data"`
}

type Device struct {
	SKU        string `json:"sku"`        // The part number of the device
	DeviceId   string `json:"device"`     // The MAC Address of the device
	DeviceName string `json:"deviceName"` // The nickname of the device
	Type       string `json:"type"`
	// Omit the capabilities array for the time being.
}
