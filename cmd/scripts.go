package cmd

import (
	"fmt"

	"github.com/LiraGui/shadow/features/scripts/busybox"
	"github.com/spf13/cobra"
)

// upBusyBoxContainerScriptCmd represents the up busybox container script command
var scriptsCmd = &cobra.Command{
	Use:     "scripts",
	Aliases: []string{"s"},
	Short:   "Manage scripts",
	Long:    ``, // TODO: Add long description
	Run: func(cmd *cobra.Command, args []string) {
		help()
	},
}

func init() {
	scriptsCmd.AddCommand(busybox.UpBusyBoxContainerScript)
	scriptsCmd.AddCommand(busybox.DeleteBusyBoxContainerScript)

	rootCmd.AddCommand(scriptsCmd)
}

func help() {
	fmt.Println("Available scripts:")
	fmt.Println("")
	fmt.Println("Busybox:")
	fmt.Println("  up-busybox-container    Up a busybox container")
	fmt.Println("  delete-busybox-container Delete a busybox container")
}
