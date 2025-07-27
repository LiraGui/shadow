/*
Copyright © 2025 LiraGui <liraguilherme252@gmail.com>

*/
package context

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// ListClusterContextCmd represents the list command
var ListClusterContextCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List Kubernetes cluster context",
	Long:    "List Kubernetes cluster context",
	Run: func(cmd *cobra.Command, args []string) {
		igniteListClusterContext()
	},
}

// NOTE: igniteListClusterContext is the function that is called when the list-cluster-context command is executed
func igniteListClusterContext() {
	fmt.Println("Listing Kubernetes cluster context:\n")
	clusterContexts := ListClusterContext()
	actualClusterContext := GetActualClusterContext()

	for index, clusterContext := range clusterContexts {
		if strings.TrimSpace(clusterContext) == strings.TrimSpace(actualClusterContext) {
			fmt.Printf("%d. \033[32m%s\033[0m\n", index+1, clusterContext)
		} else {
			fmt.Printf("%d. %s\n", index+1, clusterContext)
		}
	}
}

// NOTE: ListClusterContext returns the list of available cluster contexts
func ListClusterContext() []string {
	cmd := exec.Command("kubectl", "config", "get-contexts", "--output", "name")
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to execute command: %s\n", err)
		return nil
	}

	clusterContexts := strings.Split(string(output), "\n")
	// NOTE: Get off the last index
	clusterContexts = clusterContexts[:len(clusterContexts)-1]

	return clusterContexts
}

// NOTE: GetActualClusterContext returns the actual cluster context
func GetActualClusterContext() string {
	cmd := exec.Command("kubectl", "config", "current-context")
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to execute command: %s\n", err)
		return ""
	}
	
	return string(output)
}