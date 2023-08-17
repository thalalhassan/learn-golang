/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Test command",
	Long:  `Test  command`,
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetInt("port")
		code := viper.GetInt("code")
		fmt.Println("serve called")
		fmt.Printf("port: %v\n", port)
		fmt.Printf("code: %v\n", code)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().Int("port", 8080, "Port to run the server on")
}
