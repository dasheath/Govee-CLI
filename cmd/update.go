/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"govee/api"
	"strings"
	"time"

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

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	updateCmd.Flags().StringP("device", "d", "office", "Alias for which device to update")
	updateCmd.Flags().IntP("power", "p", 0, "Whether to power on or off")
	updateCmd.Flags().IntP("brightness", "b", 25, "Brightness on scale of 0 to 100")
}

func updateDeviceState(cmd *cobra.Command, args []string) {

	device, _ := cmd.Flags().GetString("device")
	power, _ := cmd.Flags().GetInt("power")
	if power != 0 && power != 1 {
		fmt.Println("Power selection is invalid! Please input either 0 (off) or 1 (on).")
		return
	}
	brightness, _ := cmd.Flags().GetInt("brightness")
	if brightness < 0 || brightness > 100 {
		fmt.Printf("Brightness input of %d ", brightness)
		brightness = min(max(brightness, 0), 100)
		fmt.Printf("has been clipped to the value: %d\n", brightness)
	}

	devices := api.GetDevices()
	var deviceId string
	for _, searchDevice := range devices.Data {
		if strings.Contains(strings.ToLower(searchDevice.DeviceName), strings.ToLower(device)) {
			deviceId = searchDevice.DeviceId
		}
	}

	time.Sleep(time.Second)
	api.SetDevicePower(deviceId, power)

}
