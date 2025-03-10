package initializers

import "context"

func SyncDatabase() error {
	db := Client.Database("jwt_database")
	db.CreateCollection(context.TODO(), "users")
	return nil
}
