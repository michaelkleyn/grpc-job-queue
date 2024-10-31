package main

import (
	"context"
	"log"
	"time"

	pb "github.com/michaelkleyn/grpc-job-queue/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewJobQueueClient(conn)

	// Submit a job
	jobID := "job123"
	payload := "Process data X"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	submitResp, err := client.SubmitJob(ctx, &pb.JobRequest{
		JobId:   jobID,
		Payload: payload,
	})
	if err != nil {
		log.Fatalf("Error submitting job: %v", err)
	}
	log.Printf("SubmitJob Response: %s", submitResp.Message)

	// Wait for a moment before checking status
	time.Sleep(500 * time.Millisecond)

	// Get job status
	statusResp, err := client.GetJobStatus(ctx, &pb.JobStatusRequest{
		JobId: jobID,
	})
	if err != nil {
		log.Fatalf("Error getting job status: %v", err)
	}
	log.Printf("GetJobStatus Response: Job %s is %s", statusResp.JobId, statusResp.Status)
}
