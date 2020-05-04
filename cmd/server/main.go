package main

import (
	"fmt"
	cmd "github.com/vitush/go-grpc-dg-poc/pkg/cmd/server"
	"os"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}