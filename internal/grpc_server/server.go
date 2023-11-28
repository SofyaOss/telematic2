package grpc_server

import (
	"context"
	"log"

	pb "practice/internal/grpc"
	"practice/storage/postgres"

	"google.golang.org/grpc"
)

// Server struct contains info about gRPC server
type Server struct {
	pb.UnimplementedGRPCServiceServer
	Grpc   *grpc.Server
	db     *postgres.TelematicDB
	logger *log.Logger
}

type grpcServer Server

// New creates and returns new gRPC server
func New(db *postgres.TelematicDB) *Server {
	s := Server{
		Grpc: grpc.NewServer(),
		db:   db,
	}
	pb.RegisterGRPCServiceServer(s.Grpc, (*grpcServer)(&s))
	return &s
}

// Close method stops gRPC server
func (g *grpcServer) Close() error {
	g.Grpc.GracefulStop()
	g.logger.Println("gRPC server stopped")
	return nil
}

// GetCarsByDate gRPC method for getting list of cars by its date
func (g *grpcServer) GetCarsByDate(ctx context.Context, req *pb.CarsByDateRequest) (*pb.CarsByDateResponse, error) {
	firstDate := req.GetFirstDate()
	lastDate := req.GetLastDate()
	nums := req.GetNums()
	res, err := g.db.GetByDate(ctx, firstDate, lastDate, nums)
	if err != nil {
		return nil, err
	}
	return &pb.CarsByDateResponse{
		Cars: res,
	}, nil
	//for
	//return &pb.GetByDateResponse{
	//	Cars: res,
	//}, nil
}

// GetLastCars is gRPC method for getting list of recent entries about cars by its number
func (g *grpcServer) GetLastCars(ctx context.Context, req *pb.LastCarsRequest) (*pb.LastCarsResponse, error) {
	nums := req.GetNums()
	res, err := g.db.GetByCarNumber(ctx, nums)
	if err != nil {
		return nil, err
	}
	return &pb.LastCarsResponse{
		Cars: res,
	}, nil
}
