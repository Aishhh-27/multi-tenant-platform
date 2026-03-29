package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"gitlab-multi-tenant-platform/internal/healing"
)

var healCmd = &cobra.Command{
	Use:   "heal-tenant",
	Short: "Auto-heal tenant pods",
	Run: func(cmd *cobra.Command, args []string) {
		tenantName, _ := cmd.Flags().GetString("name")

		fmt.Println("Starting auto-healing for:", tenantName)

		for {
			healing.Heal(tenantName)
			time.Sleep(30 * time.Second)
		}
	},
}

func init() {
	healCmd.Flags().StringP("name", "n", "", "Tenant name")
	healCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(healCmd)
}
