package main

import (
	"testing"
	"os"
	"log"
)

func TestAnalyser(t *testing.T) {
    log.Println("TestOne running")

	os.Args = []string{"/analyse", "12345"}

	main()
}