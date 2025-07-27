/*
Copyright © 2025 LiraGui <liraguilherme252@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shadow",
	Short: "Shadow CLI",
	Long: `
 ____  _   _    _    ____   _____        __   ____ _     ___ 
/ ___|| | | |  / \  |  _ \ / _ \ \      / /  / ___| |   |_ _|
\___ \| |_| | / _ \ | | | | | | \ \ /\ / /  | |   | |    | | 
 ___) |  _  |/ ___ \| |_| | |_| |\ V  V /   | |___| |___ | | 
|____/|_| |_/_/   \_|____/ \___/  \_/\_/     \____|_____|___|
	`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "~/.shadow.yaml", "config file (default is ~/.shadow.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


