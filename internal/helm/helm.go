package helm

import (
	"fmt"
	"os/exec"
)

func Deploy(tenant string) {
	fmt.Println("Deploying Helm chart for:", tenant)

	cmd := exec.Command("helm", "upgrade", "--install",
		tenant,
		"./helm/gitlab-app",
		"-n", tenant)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Helm deploy failed:", err)
	}
}
