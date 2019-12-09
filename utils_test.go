package main

import (
	"net"
	"testing"
)

func TestGetPrivateIP(t *testing.T) {
	privateIP, err := getPrivateIP()

	if err != nil {
		t.Fatalf("Failed to get private IP")
	}

	result := net.ParseIP(privateIP)

	if result == nil {
		t.Fatalf("Invalid IP address")
	}
}
