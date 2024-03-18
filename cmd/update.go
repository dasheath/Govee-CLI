/*
Copyright © 2024 Heath McCabe
*/
package cmd

import (
	"fmt"
	"govee/api"
	"strings"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the state of one of your devices.",
	Long: `Update the power, brightness, or color of a bulb. 
	This could also update various fields of any of your devices.
	
	Example:
	> govee update -d bedroom -p 1 -b 100
	`,

	Run: updateDeviceState,
}

const (
	SentinelInt int = 1005
)

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("device", "d", "office", "Alias for which device to update")
	updateCmd.Flags().IntP("power", "p", SentinelInt, "Whether to power on or off")
	updateCmd.Flags().IntP("brightness", "b", SentinelInt, "Brightness on scale of 0 to 100")
}

func updateDeviceState(cmd *cobra.Command, args []string) {

	device, _ := cmd.Flags().GetString("device")
	devices := api.GetDevices()
	var deviceId string
	for _, searchDevice := range devices.Data {
		if strings.Contains(strings.ToLower(searchDevice.DeviceName), strings.ToLower(device)) {
			deviceId = searchDevice.DeviceId
		}
	}

	if deviceId == "" {
		fmt.Println("No device could be found with that nickname! Please try again.")
		return
	}

	power, _ := cmd.Flags().GetInt("power")
	if power != SentinelInt {
		if power != 0 && power != 1 {
			fmt.Println("Power selection is invalid! Please input either 0 (off) or 1 (on).")
			return
		}
		api.SetDevicePower(deviceId, power)
	}

	brightness, _ := cmd.Flags().GetInt("brightness")
	if brightness != SentinelInt {
		if brightness < 0 || brightness > 100 {
			fmt.Printf("Brightness input of %d ", brightness)
			brightness = min(max(brightness, 0), 100)
			fmt.Printf("has been clipped to the value: %d\n", brightness)
		}
		api.SetDeviceBrightness(deviceId, brightness)
	}
}
