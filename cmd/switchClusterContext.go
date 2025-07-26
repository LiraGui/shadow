/*
Copyright © 2025 LiraGui <liraguilherme252@gmail.com>

*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"
	"github.com/spf13/cobra"
)

// switchClusterContextCmd represents the switchClusterContext command
var switchClusterContextCmd = &cobra.Command{
	Use:     "switch-cluster-context",
	Aliases: []string{"sc	c"},
	Short:   "Switch Kubernetes cluster context",
	Long:    "Switch Kubernetes cluster context",
	Run: func(cmd *cobra.Command, args []string) {
		ignite()
	},
}

func init() {
	rootCmd.AddCommand(switchClusterContextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// switchClusterContextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// switchClusterContextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// NOTE: ignite is the function that is called when the switch-cluster-context command is executed
func ignite() {
	fmt.Println("Select the cluster context:\n")
	SwitchClusterContext()
}

// NOTE: SwitchClusterContext is the function that is called when the switch-cluster-context command is executed
func SwitchClusterContext() {
	// NOTE: Get the list of available cluster contexts
	clusterContexts := ListClusterContext()
	actualClusterContext := GetActualClusterContext()

	for index, clusterContext := range clusterContexts {
		if strings.TrimSpace(clusterContext) == strings.TrimSpace(actualClusterContext) {
			fmt.Printf("%d. \033[32m%s\033[0m\n", index+1, clusterContext)
		} else {
			fmt.Printf("%d. %s\n", index+1, clusterContext)
		}
	}
	
	// NOTE: Prompt user to select a context
	fmt.Print("\nEnter the number of the context to switch to: ")
	var selection int
	fmt.Scanln(&selection)
	
	// NOTE: Validate selection
	if selection < 1 || selection > len(clusterContexts) {
		fmt.Println("Invalid selection. Please try again.")
		return
	}
	
	// NOTE: Get the selected context
	selectedContext := clusterContexts[selection-1]
	
	// NOTE: Execute kubectl command to switch context
	cmd := exec.Command("kubectl", "config", "use-context", selectedContext)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to switch context: %s\n", err)
		return
	}
	
	fmt.Printf("Switched to context: %s\n", selectedContext)
}	
