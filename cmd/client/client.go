package main

import (
	"context"
	"fmt"
	"time"
	"log"
	"io"
	"github.com/iamoreira/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main () {
	connection , err := grpc.Dial("grpc-server:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUserStreamBoth(client)

}

func AddUser(client pb.UserServiceClient) {

	req := &pb.User {
		Id: "0",
		Name: "Joao",
		Email: "Joao@example.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User {
		Id: "0",
		Name: "Joao",
		Email: "Joao@example.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break;
		}
		
		if err != nil {
			log.Fatalf("Could not receive the msg %v", err)
		}

		fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User {
		&pb.User {
			Id: "w1",
			Name: "Name1",
			Email: "wes@example.com",
		},
		&pb.User {
			Id: "w2",
			Name: "Name2",
			Email: "wes@example.com",
		},
		&pb.User {
			Id: "w3",
			Name: "Name3",
			Email: "wes@example.com",
		},
		&pb.User {
			Id: "w4",
			Name: "Name4",
			Email: "wes@example.com",
		},
		&pb.User {
			Id: "w5",
			Name: "Name5",
			Email: "wes@example.com",
		},
		&pb.User {
			Id: "w6",
			Name: "Name6",
			Email: "wes@example.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating request %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())

	if err != nil  {
		log.Fatalf("Error creating request %v", err)
	}

	reqs := []*pb.User {
		&pb.User {
			Id: "w1",
			Name: "Name1",
			Email: "wes@example.com",
		},
		&pb.User {
			Id: "w2",
			Name: "Name2",
			Email: "wes@example.com",
		},
		&pb.User {
			Id: "w3",
			Name: "Name3",
			Email: "wes@example.com",
		},
		&pb.User {
			Id: "w4",
			Name: "Name4",
			Email: "wes@example.com",
		},
		&pb.User {
			Id: "w5",
			Name: "Name5",
			Email: "wes@example.com",
		},
		&pb.User {
			Id: "w6",
			Name: "Name6",
			Email: "wes@example.com",
		},
	}

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()
	
	wait := make(chan int) 

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
				break
			}

			fmt.Printf("Recebendo user %v com status: %v \n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait

}