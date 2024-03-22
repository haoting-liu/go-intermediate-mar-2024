package main

import (
	"fmt"
	"runtime"

	_ "go.uber.org/automaxprocs"
)

/**
When containering application package your application with Minimal OS (alpine/slim)
Don't package your application with the SDK

Multi stage docker build
- stage 1
	- builds the application and produces the binary
- state 2
	- get the binary from stage 1 and package it with a Minimal OS
	- also the stage-1 output is discarded later

docker build -t myapp:1.0 .

docker run --rm --cpus=0.5 myapp:1.0
docker run --rm --cpus=1.5 myapp:1.0
docker run --rm --cpus=2 myapp:1.0
*/

func main() {
	fmt.Println("PROCS: ", runtime.GOMAXPROCS(0))
	fmt.Println("CPU: ", runtime.NumCPU())
}
