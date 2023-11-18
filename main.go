package main

import (
	"context"
	"log"
	"net"

	Invoicer "github.com/abhinandpn/gRPC-Demo/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	Invoicer.UnimplementedInvoicerServer
}

func (s *myInvoicerServer) Create(ctx context.Context, req *Invoicer.CreateRequest) (*Invoicer.CreateResponse, error) {
	return &Invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegister := grpc.NewServer()
	service := &myInvoicerServer{}
	Invoicer.RegisterInvoicerServer(serverRegister, service)
	err = serverRegister.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
