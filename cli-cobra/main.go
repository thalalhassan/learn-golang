/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cli-cobra/cmd"
	"cli-cobra/config"
)

func main() {
	config.LoadConfig()
	cmd.Execute()
}
