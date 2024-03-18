/*
Copyright © 2024 Heath McCabe
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "govee",
	Short: "Get and Update info about Govee devices.",
	Long: `This program allows you to see the devices on your Govee account,
	set the power state of those devices, and set the brightness of those devices.
	
	This was implemented with the express purpose of affecting H6008 bulbs, 
	but could easily be extended to interact with other devices.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
