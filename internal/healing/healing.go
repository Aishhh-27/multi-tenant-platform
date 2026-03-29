package healing

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func Heal(namespace string) {
	fmt.Println("Checking pods in:", namespace)

	cmd := exec.Command("kubectl", "get", "pods", "-n", namespace)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error fetching pods:", err)
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))

	var failingPods []string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "CrashLoopBackOff") ||
			strings.Contains(line, "Error") ||
			strings.Contains(line, "ImagePullBackOff") {

			fields := strings.Fields(line)
			if len(fields) > 0 {
				podName := fields[0]
				failingPods = append(failingPods, podName)
			}
		}
	}

	if len(failingPods) == 0 {
		fmt.Println("No issues detected")
		return
	}

	fmt.Println("Failing pods:", failingPods)

	for _, pod := range failingPods {
		fmt.Println("Deleting pod:", pod)
		exec.Command("kubectl", "delete", "pod", pod, "-n", namespace).Run()
	}

	fmt.Println("Healing complete for:", namespace)
}
