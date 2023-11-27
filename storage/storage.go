package storage

import (
	"context"
	"time"

	pb "practice/internal/grpc"
)

// Coordinates struct contains info about latitude
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// Car struct contains car info
type Car struct {
	ID     int
	Number int
	Speed  int
	Coordinates
	Date time.Time
}

// DBInterface describes db methods
type DBInterface interface {
	CreateTable(ctx context.Context) error
	DropTable(ctx context.Context) error
	AddData(ctx context.Context, c *Car) error
	GetByDate(ctx context.Context, d1s, d2s string, nums []int64) ([]*pb.Car, error)
	GetByCarNumber(ctx context.Context, carNums []int64) ([]*pb.Car, error)
}
