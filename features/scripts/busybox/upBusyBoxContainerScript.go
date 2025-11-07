package busybox

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// upBusyBoxContainerScriptCmd represents the up busybox container script command
var UpBusyBoxContainerScript = &cobra.Command{
	Use:     "up-busybox-container",
	Aliases: []string{"up-busybox"},
	Short:   "Up a busybox container",
	Long:    ``, // TODO: Add long description
	Run: func(cmd *cobra.Command, args []string) {
		igniteUpBusyBoxContainerScript()
	},
}

func igniteUpBusyBoxContainerScript() {
	if verifyIfBusyBoxContainerIsRunning() {
		return
	}

	fmt.Println("Up busybox container called...")
	cmd := exec.Command("kubectl", "run", "busybox", "--image=busybox", "--restart=Never", "--command", "--", "sleep", "3600")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to execute command: %s\n", err)
		return
	}

	fmt.Println(strings.TrimSpace(string(output)))
}

func verifyIfBusyBoxContainerIsRunning() bool {
	fmt.Println("Verify if busybox container is running called...")
	cmd := exec.Command("kubectl", "get", "pod", "busybox")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(strings.TrimSpace(string(output)))
		return false
	}

	if strings.Contains(string(output), "busybox") {
		fmt.Println("Busybox container is already running")
		return true
	}

	return false
}
