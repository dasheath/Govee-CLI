/*
Copyright © 2024 Heath McCabe
*/
package cmd

import (
	"fmt"
	"govee/api"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your devices",
	Long:  `List all the devices attached to your Govee account.`,
	Run: func(cmd *cobra.Command, args []string) {

		devices := api.GetDevices()
		for i, device := range devices.Data {
			fmt.Printf("Device Number %d\n", i)
			fmt.Println("\tSKU:", device.SKU)
			fmt.Println("\tDevice:", device.DeviceId)
			fmt.Println("\tDevice Name:", device.DeviceName)
			fmt.Println("\tType:", device.Type)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
