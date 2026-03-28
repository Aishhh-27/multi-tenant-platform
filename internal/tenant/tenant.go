package tenant

import (
	"fmt"
	"os"
)

func Create(name string) {
	fmt.Println("Initializing tenant:", name)

	err := os.MkdirAll("tenants/"+name, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating tenant folder:", err)
	}
}
