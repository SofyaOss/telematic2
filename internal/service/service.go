package service

import (
	"context"

	"practice/storage"
	"practice/storage/postgres"
)

// CreateTelematicTable calls the method for creating telematic table in db
func CreateTelematicTable(ctx context.Context, db *postgres.TelematicDB) error {
	return db.CreateTable(ctx)
}

// DropTelematicTable calls the method for dropping telematic table in db
func DropTelematicTable(ctx context.Context, db *postgres.TelematicDB) error {
	return db.DropTable(ctx)
}

// AddTelematic calls the method for adding telematic in db
func AddTelematic(ctx context.Context, db *postgres.TelematicDB, data *storage.Car) error {
	return db.AddData(ctx, data)
}
