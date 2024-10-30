package main

import (
	pb "github.com/michaelkleyn/grpc-job-queue/proto/"
	// pb "../proto/"
)

type jobStatus struct {
	status string
}

type server struct {
	pb.UnimplementedJobQueueServer
}
