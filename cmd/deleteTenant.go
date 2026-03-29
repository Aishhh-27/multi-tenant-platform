package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"gitlab-multi-tenant-platform/internal/terraform"
	"gitlab-multi-tenant-platform/internal/kubernetes"
	"gitlab-multi-tenant-platform/internal/helm"
	"gitlab-multi-tenant-platform/internal/tenant"
)

var deleteTenantCmd = &cobra.Command{
	Use:   "delete-tenant",
	Short: "Delete an existing tenant",
	Run: func(cmd *cobra.Command, args []string) {
		tenantName, _ := cmd.Flags().GetString("name")

		fmt.Println("Deleting tenant:", tenantName)

		helm.Delete(tenantName)
		kubernetes.DeleteNamespace(tenantName)
		terraform.Destroy(tenantName)
		tenant.Delete(tenantName)

		fmt.Println("Tenant deleted successfully!")
	},
}

func init() {
	deleteTenantCmd.Flags().StringP("name", "n", "", "Tenant name")
	deleteTenantCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(deleteTenantCmd)
}
