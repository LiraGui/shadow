/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/LiraGui/shadow/features/context"
)

// contextCmd represents the context command
var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Manage Kubernetes cluster context",
	Long: `Manage Kubernetes cluster context`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("context called")
	},
}

func init() {
	// Add context commands as subcommands
	contextCmd.AddCommand(context.ListClusterContextCmd)
	contextCmd.AddCommand(context.SwitchClusterContextCmd)
	
	// Add context command to root command
	rootCmd.AddCommand(contextCmd)
}
