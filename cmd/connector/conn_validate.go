package connector

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/olekukonko/tablewriter"
	connvalidate "github.com/sailpoint-oss/sailpoint-cli/cmd/connector/validate"
	"github.com/sailpoint-oss/sailpoint-cli/internal/client"
	"github.com/spf13/cobra"
)

func newConnValidateCmd(apiClient client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate connector behavior",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			//endpoint := cmd.Flags().Lookup("conn-endpoint").Value.String()

			// Check if we just need to list checks
			list, _ := strconv.ParseBool(cmd.Flags().Lookup("list").Value.String())
			if list {
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"ID", "Description"})
				for _, c := range connvalidate.Checks {
					table.Append([]string{
						c.ID,
						c.Description,
					})
				}
				table.Render()
				return nil
			}

			cc, err := connClient(cmd, apiClient)
			if err != nil {
				return err
			}
			//cc := connclient.NewConnClient(spClient, v, cfg, connectorRef, endpoint)

			check := cmd.PersistentFlags().Lookup("check").Value.String()
			log.Printf("CHECK %q", check)

			isReadOnly, _ := strconv.ParseBool(cmd.Flags().Lookup("read-only").Value.String())
			valid := connvalidate.NewValidator(connvalidate.Config{
				Check:    check,
				ReadOnly: isReadOnly,
			}, cc)

			results, err := valid.Run(ctx)
			if err != nil {
				return err
			}
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Result", "Errors", "Warnings", "Skipped"})
			hasFailedCheck := false
			for _, res := range results {
				var result = aurora.Green("PASS")
				if len(res.Errors) > 0 {
					hasFailedCheck = true
					result = aurora.Red("FAIL")
				}

				if len(res.Skipped) > 0 {
					result = aurora.Yellow("SKIPPED")
				}

				table.Append([]string{
					aurora.Blue(res.ID).String(),
					result.String(),
					aurora.Red(strings.Join(res.Errors, "\n\n")).String(),
					aurora.Yellow(strings.Join(res.Warnings, "\n\n")).String(),
					aurora.Yellow(strings.Join(res.Skipped, "\n\n")).String(),
				})
			}
			table.Render()

			if hasFailedCheck {
				return fmt.Errorf("at least one check has failed")
			}
			return nil
		},
	}

	cmd.PersistentFlags().StringP("check", "", "", "Run a specific check")
	cmd.PersistentFlags().BoolP("list", "l", false, "List checks; don't run checks")
	cmd.PersistentFlags().BoolP("read-only", "r", false, "Run all checks that don't modify connector's data")

	cmd.PersistentFlags().StringP("version", "v", "", "Run against a specific version")
	cmd.MarkFlagRequired("version")

	cmd.PersistentFlags().StringP("config-path", "p", "", "Path to config to use for test command")
	cmd.MarkFlagRequired("config-path")

	cmd.PersistentFlags().StringP("id", "c", "", "Connector ID or Alias")
	cmd.MarkFlagRequired("id")

	return cmd
}
