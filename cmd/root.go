package cmd

import (
	"context"
	"log"

	"github.com/byhammert/lj-go/infra/database"
	"github.com/byhammert/lj-go/infra/database/mongo"
	"github.com/byhammert/lj-go/infra/database/mongo/repositories"
	"github.com/spf13/cobra"
)

func Execute() {
	root := &cobra.Command{
		Short:   "LJ-GO",
		Version: "1.0.0",
	}

	root.AddCommand(
		Api,
	)

	root.Execute()
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetDatabase(ctx context.Context) *database.Database {
	client, err := mongo.GetConnection(ctx)
	FatalError(err)

	studentRepository := repositories.NewCategoryRepository(ctx, client)

	return database.NewDatabase(client, studentRepository)
}
