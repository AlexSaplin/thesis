package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"k8s.io/client-go/util/homedir"

	"github.com/gofrs/uuid"

	"tesseract/pkg/k8s"
	pb "tesseract/pkg/service/pb"
)

var args = k8s.TemplateArgs{
	Name:     "example-nginx-xxx",
	DNS:      "mirror.example.com",
	Image:    "mirroraieuw.azurecr.io/app/api:latest",
	Port:     80,
	Scale:    3,
	CPU:      50,
	MemoryMB: 512,
	GPU:      "none",
	Env: []*pb.KV{
		{
			Key:   "key",
			Value: "value",
		},
		{
			Key:   "key2",
			Value: "value2",
		},
	},
	Auth: "",
}

func main() {
	namespace := uuid.Must(uuid.NewV4()).String()
	args.Namespace = namespace

	client, err  := k8s.NewK8SClient(filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		panic(err)
	}

	err = client.Apply(context.Background(), args)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 2; i++ {
		time.Sleep(time.Second)
		res, err := client.Get(context.Background(), args.Namespace)
		if err != nil {
			panic(err)
		}
		log.Println(res)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	_, _ = reader.ReadString('\n')
	err = client.Down(context.Background(), args.Namespace)
	if err != nil {
		panic(err)
	}
}


