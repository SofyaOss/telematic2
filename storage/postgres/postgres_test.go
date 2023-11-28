package postgres

import (
	"context"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"practice/internal/generator"
	"practice/storage"
	"testing"
	"time"
)

func TestTelematicDB_AddData(t *testing.T) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %v", err)
	}

	resourceDB, err := pool.RunWithOptions(&dockertest.RunOptions{
		Name:         "postgres-test",
		Repository:   "postgres",
		Tag:          "latest",
		Env:          []string{"POSTGRES_PASSWORD=postgres", "POSTGRES_USER=postgres", "POSTGRES_DB=testdb"},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: "5432"},
			},
		},
	})
	if err != nil {
		log.Fatalf("could not start resourse: %s", err)
	}
	defer pool.Purge(resourceDB)

	var db *TelematicDB
	err = pool.Retry(func() error {
		dbConnStr := fmt.Sprintf("host=localhost port=5432 user=postgres dbname=testdb password=postgres sslmode=disable")
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		db, err = New(ctx, dbConnStr) // подключение к базе данных
		if err != nil {
			log.Println("db not ready yet")
			return err
		}

		err = db.DropTable(ctx)
		if err != nil {
			cancel()
			return err
		}

		err = db.CreateTable(ctx)
		if err != nil {
			cancel()
			return err
		}

		for i := 0; i < 7; i++ {
			c := &storage.Car{
				ID:          i,
				Number:      i,
				Speed:       100,
				Coordinates: storage.Coordinates{80, 80},
				Date:        generator.RandomTimestamp(),
			}
			err = db.AddData(ctx, c)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Postgres test failed: %s", err)
	}
	if err = pool.Purge(resourceDB); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

/*
func TestNew(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	_, err := New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
	}
}

func TestTelematicDB_DropTable(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	newDB, err := New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
	}
	err = newDB.DropTable(ctx)
	if err != nil {
		log.Println("failed to drop table with error:", err)
	}
}

func TestTelematicDB_CreateTable(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	newDB, err := New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
	}
	err = newDB.CreateTable(ctx)
	if err != nil {
		log.Println("failed to create table with error:", err)
	}
}

*/
