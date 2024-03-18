# Govee CLI

This simple CLI implemented in Golang allows the user to see and update information about their Govee devices.

### Local Setup

1. You will need to get an API key from Govee
1. Add that API key to a `.env` file with the key of `Govee_API_Key`
1. Run `go build`, `go install`
1. You can now use the CLI!

### Usage

1. Run `govee list` to see a list of your devices
1. Run `govee update -h` to see the options available for updating your devices
   - At time of writing: the `device` option is used to determine which of your devices to update by searching for any device whose `deviceName` field (as per the response from the `GetDevices` function) contains the value you provide.
   - This is suitable for my simple use case, but could do with being made more robust should the need arise.
