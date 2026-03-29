package kubernetes

import (
	"fmt"
	"os/exec"
)

func CreateNamespace(name string) {
	fmt.Println("Creating namespace:", name)

	err := exec.Command("kubectl", "create", "namespace", name).Run()
	if err != nil {
		fmt.Println("Namespace may already exist:", err)
	}
}

func DeleteNamespace(name string) {
	fmt.Println("Deleting namespace:", name)

	exec.Command("kubectl", "delete", "namespace", name).Run()
}
