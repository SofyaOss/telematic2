package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	pb "practice/internal/grpc"
	"practice/storage"
	"time"
)

// TelematicDB is a struct for telematic db
type TelematicDB struct {
	db *pgxpool.Pool
}

// New creates connection to db
func New(ctx context.Context, databaseUrl string) (*TelematicDB, error) {
	for {
		_, err := pgxpool.Connect(ctx, databaseUrl)
		if err == nil {
			break
		}
	}
	db, err := pgxpool.Connect(ctx, databaseUrl)
	if err != nil {
		return nil, err
	}
	t := TelematicDB{
		db: db,
	}
	return &t, nil
}

// CreateTable creates telematic table
func (t *TelematicDB) CreateTable(ctx context.Context) error {
	_, err := t.db.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS telematic (
			id SERIAL PRIMARY KEY, 
			car_number INT NOT NULL DEFAULT 1,
			speed INT NOT NULL DEFAULT 0,
			latitude FLOAT NOT NULL DEFAULT 0,
			longitude FLOAT NOT NULL DEFAULT 0,
			date DATE NOT NULL DEFAULT CURRENT_DATE
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

// DropTable drops telematic table
func (t *TelematicDB) DropTable(ctx context.Context) error {
	_, err := t.db.Exec(ctx, `DROP TABLE IF EXISTS telematic;`)
	if err != nil {
		return err
	}
	return nil
}

// GetByDate get cars by date for gRPC
func (t *TelematicDB) GetByDate(ctx context.Context, d1s, d2s string, nums []int64) ([]*pb.Car, error) {
	d1, err1 := time.Parse("2006-01-02", d1s)
	if err1 != nil {
		return nil, err1
	}
	d2, err2 := time.Parse("2006-01-02", d2s)
	if err2 != nil {
		return nil, err2
	}
	rows, err := t.db.Query(ctx, `SELECT * FROM telematic WHERE (date BETWEEN $1 AND $2) AND (car_number = ANY($3));`, d1, d2, pq.Array(nums))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cars []*pb.Car
	for rows.Next() {
		var car storage.Car
		err = rows.Scan(
			&car.ID,
			&car.Number,
			&car.Speed,
			&car.Coordinates.Latitude,
			&car.Coordinates.Longitude,
			&car.Date,
		)
		if err != nil {
			return nil, err
		}
		cars = append(cars, ConvertCarsToProtoCars(&car))
	}
	return cars, nil
}

// GetByCarNumber get cars by its number for gRPC
func (t *TelematicDB) GetByCarNumber(ctx context.Context, carNums []int64) ([]*pb.Car, error) {
	var cars []*pb.Car
	for _, num := range carNums {
		rows, err := t.db.Query(ctx, `SELECT * FROM telematic WHERE car_number=$1 ORDER BY id desc LIMIT 1;`, num)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var car storage.Car
			err = rows.Scan(
				&car.ID,
				&car.Number,
				&car.Speed,
				&car.Coordinates.Latitude,
				&car.Coordinates.Longitude,
				&car.Date,
			)
			if err != nil {
				return nil, err
			}
			cars = append(cars, ConvertCarsToProtoCars(&car))
		}
	}
	return cars, nil
}

// AddData add telematic to db
func (t *TelematicDB) AddData(ctx context.Context, c *storage.Car) error {
	err := t.db.QueryRow(ctx,
		`INSERT INTO telematic (car_number, speed, latitude, longitude, date) VALUES ($1, $2, $3, $4, $5);`,
		c.Number, c.Speed, fmt.Sprint(c.Coordinates.Latitude), fmt.Sprint(c.Coordinates.Longitude), c.Date).Scan()
	return err
}
