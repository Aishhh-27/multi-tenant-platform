package terraform

import (
	"fmt"
	"os"
	"os/exec"
)

func Apply(tenant string) {
	fmt.Println("Running Terraform for:", tenant)

	// Create workspace (ignore error if exists)
	cmd := exec.Command("terraform", "workspace", "new", tenant)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()

	// Apply Terraform
	cmd = exec.Command("terraform", "apply", "-auto-approve", "-var=tenant_name="+tenant)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Terraform apply failed:", err)
	}
}

func Destroy(tenant string) {
	fmt.Println("Destroying Terraform for:", tenant)

	exec.Command("terraform", "workspace", "select", tenant).Run()

	cmd := exec.Command("terraform", "destroy", "-auto-approve", "-var=tenant_name="+tenant)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
