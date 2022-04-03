package migrations

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo/bson"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	fmt.Println("migration start")
	ctx := context.Background()
	migrate.Register(func(db *mongo.Database) error { //Up
		var err error

		 _, err = db.Collection("characters").InsertOne(ctx,
		 	bson.D{
		 	{"name","Poppy"},
		 	{"description","Poppy is always a cheerful princess. She is the leader of a group of 10 BFFS (Best Friends) called Snack Pack and loves to sing and hug."},
		 	{"Image","poppy.jpg"},
		 	})
		if err != nil {
			fmt.Println(err)
			return err
		}

		_, err = db.Collection("characters").InsertOne(ctx,
			bson.D{
				{"name","Branch"},
				{"description","Branch is not the same as the other characters. He's not the biggest fan of the usual activities that involve singing, dancing, and hugging. In addition, he is too cautious about everything, and the others do not understand him well. But when problems arise, Branch is the only one who is ready for them."},
				{"Image","branch.jpg"},
			})
		if err != nil {
			fmt.Println(err)
			return err
		}

		_, err = db.Collection("characters").InsertOne(ctx,
			bson.D{
				{"name","Biggie & Mr. Dinkles"},
				{"description","Biggie is the largest, but the softest of all. Of course, he has a big stature, but his heart is even bigger, and he tends to cry happy tears often. The big man and his pet, a worm named Mr. Dinkles, are practically inseparable."},
				{"Image","b&m.jpg"},
			})
		if err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	}, func(db *mongo.Database) error {
		_, err := db.Collection("characters").DeleteMany(ctx, "{}")
		return err
	})
}
