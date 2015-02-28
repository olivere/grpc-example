package main

import (
	"flag"
	"log"
	"os"

	"github.com/olivere/grpc-example/tasks"
)

var (
	addr = flag.String("addr", ":8000", "server address, e.g. :8000")
	cert = flag.String("cert", "", "cert file")
	key  = flag.String("key", "", "key file")
)

func main() {
	flag.Parse()

	s, err := tasks.NewServer(
		tasks.SetAddr(*addr),
		tasks.SetTLS(*cert, *key),
		tasks.SetLogger(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		log.Fatal(err)
	}

	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
