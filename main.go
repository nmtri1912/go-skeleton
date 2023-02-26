package main

import "go-skeleton/cmd"

func main() {
	s := cmd.NewServer()
	s.Start()
}
