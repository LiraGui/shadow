package busybox

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// upBusyBoxContainerScriptCmd represents the up busybox container script command
var DeleteBusyBoxContainerScript = &cobra.Command{
	Use:     "delete-busybox-container",
	Aliases: []string{"del-busybox"},
	Short:   "Delete a busybox container",
	Long:    ``, // TODO: Add long description
	Run: func(cmd *cobra.Command, args []string) {
		igniteDeleteBusyBoxContainerScript()
	},
}

func igniteDeleteBusyBoxContainerScript() {
	fmt.Println("Delete busybox container called")
	cmd := exec.Command("kubectl", "delete", "pod", "busybox")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to execute command: %s\n", err)
		return
	}

	fmt.Println(strings.TrimSpace(string(output)))
}
