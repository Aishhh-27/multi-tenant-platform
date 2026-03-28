package terraform

import (
	"fmt"
	"os/exec"
)

func Apply(tenant string) {
	fmt.Println("Running Terraform for:", tenant)

	exec.Command("terraform", "workspace", "new", tenant).Run()

	cmd := exec.Command("terraform", "apply", "-auto-approve",
		"-var=tenant_name="+tenant)

	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Run()
	if err != nil {
		fmt.Println("Terraform apply failed:", err)
	}
}
