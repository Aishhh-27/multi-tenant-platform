package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"gitlab-multi-tenant-platform/internal/tenant"
	"gitlab-multi-tenant-platform/internal/terraform"
	"gitlab-multi-tenant-platform/internal/kubernetes"
	"gitlab-multi-tenant-platform/internal/helm"
)

var createTenantCmd = &cobra.Command{
	Use:   "create-tenant",
	Short: "Create a new tenant",
	Run: func(cmd *cobra.Command, args []string) {
		tenantName, _ := cmd.Flags().GetString("name")

		fmt.Println("Creating tenant:", tenantName)

		tenant.Create(tenantName)
		terraform.Apply(tenantName)
		kubernetes.CreateNamespace(tenantName)
		helm.Deploy(tenantName)

		fmt.Println("Tenant created successfully!")
	},
}

func init() {
	createTenantCmd.Flags().StringP("name", "n", "", "Tenant name")
	createTenantCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(createTenantCmd)
}
