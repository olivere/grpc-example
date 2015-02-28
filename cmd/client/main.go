package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/olivere/grpc-example/tasks"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	addr       = flag.String("addr", ":8000", "Server address, e.g. :8000")
	cert       = flag.String("cert", "", "CA certificate file")
	servername = flag.String("name", "", "Server name")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	if *cert != "" {
		cred, err := credentials.NewClientTLSFromFile(*cert, *servername)
		if err != nil {
			log.Fatal(err)
		}
		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := tasks.NewServiceClient(conn)

	req := &tasks.ListRequest{}
	res, err := c.List(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	if res != nil && res.Total > 0 {
		//fmt.Printf("%d task(s)\n", res.Total)
		for i, task := range res.Tasks {
			fmt.Printf("%2d. %s (#%d)\n", i+1, task.Name, task.Id)
		}
	} else {
		fmt.Print("No tasks\n")
	}
}
