/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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
			fmt.Println("\tDevice:", device.Device)
			fmt.Println("\tDevice Name:", device.DeviceName)
			fmt.Println("\tType:", device.Type)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
