package cmd

import (
	"context"
	"dysn/character/internal/config"
	_ "dysn/character/migrations"
	"fmt"
	"github.com/spf13/cobra"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)
const dbUri = "mongodb://%s:%s/"

var mgrCmd = &cobra.Command{
	Use: "migration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrations start")
		ctx := context.Background()
		cfg := config.NewConfig()
		uri := fmt.Sprintf(dbUri, cfg.GetHost(), cfg.GetPort())

		clientOptions := options.Client().ApplyURI(uri).
			SetAuth(options.Credential{
				AuthSource: cfg.GetDbName(),
				Username:   cfg.GetUser(),
				Password:   cfg.GetPassword(),
			})
		session, err := mongo.Connect(ctx, clientOptions)

		if err != nil {
			fmt.Println(err)
		}
		db := session.Database(cfg.GetDbName())
		migrate.SetDatabase(db)

		if (strings.ToLower(args[1]) == "up"){
			if err := migrate.Up(migrate.AllAvailable); err != nil {
				fmt.Println(err)
			}
		}

		if (strings.ToLower(args[1]) == "down"){
			if err := migrate.Down(migrate.AllAvailable); err != nil {
				fmt.Println(err)
			}
		}

		fmt.Println("OK")
	},
}

func init() {
	rootCmd.AddCommand(mgrCmd)
}
