package main

import (
	"fmt"
	"sync"
)

const grpcPort = 50051

type inventoryService struct {
	mu    sync.RWMutex
	parts map[string]string
}

func main() {
	fmt.Println("Hello World")
}
