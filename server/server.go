package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/michaelkleyn/grpc-job-queue/proto"
	// pb "../proto/"
	"google.golang.org/grpc"
)

type jobStatus struct {
	status string
}

type server struct {
	pb.UnimplementedJobQueueServer
	jobStore map[string]*jobStatus
	mu       sync.Mutex
}

func newServer() *server {
	return &server{
		jobStore: make(map[string]*jobStatus),
	}
}

func (s *server) SubmitJob(ctx context.Context, req *pb.JobRequest) (*pb.JobResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Simulate job submission by storing it with "queued" status.
	s.jobStore[req.JobId] = &jobStatus{status: "queued"}
	log.Printf("Job %s submitted with payload: %s", req.JobId, req.Payload)

	return &pb.JobResponse{
		JobId:   req.JobId,
		Message: "Job submitted successfully",
	}, nil
}

func (s *server) GetJobStatus(ctx context.Context, req *pb.JobStatusRequest) (*pb.JobStatusResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	status, exists := s.jobStore[req.JobId]
	if !exists {
		return &pb.JobStatusResponse{
			JobId:  req.JobId,
			Status: "not found",
		}, nil
	}

	// Simulate job processing by changing status after a query.
	if status.status == "queued" {
		status.status = "completed"
	}

	log.Printf("Job %s status requested: %s", req.JobId, status.status)

	return &pb.JobStatusResponse{
		JobId:  req.JobId,
		Status: status.status,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterJobQueueServer(grpcServer, newServer())

	log.Println("Server is running on port 50052...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
