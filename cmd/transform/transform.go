// Copyright (c) 2021, SailPoint Technologies, Inc. All rights reserved.
package transform

import (
	"fmt"

	"github.com/sailpoint-oss/sailpoint-cli/internal/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	transformsEndpoint = "/v3/transforms"
	//previewEndpoint         = "/cc/api/user/preview"
	previewEndpoint         = "/beta/identity-profiles/identity-preview"
	identityProfileEndpoint = "/v3/identity-profiles"
	searchEndpoint          = "/v3/search"
	//userEndpoint            = "/cc/api/identity/list"
)

func NewTransformCmd(client client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "transform",
		Short:   "Manage transforms",
		Aliases: []string{"tran"},
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = fmt.Fprint(cmd.OutOrStdout(), cmd.UsageString())
		},
	}

	cmd.PersistentFlags().StringP("transforms-endpoint", "e", viper.GetString("baseurl")+transformsEndpoint, "Override transforms endpoint")

	cmd.AddCommand(
		newListCmd(client),
		newDownloadCmd(client),
		newCreateCmd(client),
		newUpdateCmd(client),
		newDeleteCmd(client),
		newPreviewCmd(client),
	)

	return cmd
}
