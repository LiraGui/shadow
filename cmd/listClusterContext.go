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

// NOTE: listClusterContextCmd represents the list-cluster-context command
var listClusterContextCmd = &cobra.Command{
	Use:     "list-cluster-context",
	Aliases: []string{"lcc"},
	Short:   "List Kubernetes cluster context",
	Long:    "List Kubernetes cluster context",
	Run: func(cmd *cobra.Command, args []string) {
		igniteListClusterContext()
	},
}

func init() {
	rootCmd.AddCommand(listClusterContextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listClusterContextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listClusterContextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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