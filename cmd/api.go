package cmd

import (
	"context"

	"github.com/byhammert/lj-go/api"
	"github.com/byhammert/lj-go/infra/config"
	"github.com/spf13/cobra"
)

var Api = &cobra.Command{
	Use:   "api",
	Short: "API lj-go register",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ctx := context.TODO()

		err = config.StartConfig()
		FatalError(err)

		db := GetDatabase(ctx)

		err = api.NewService(db).Start()
		FatalError(err)
	},
}
