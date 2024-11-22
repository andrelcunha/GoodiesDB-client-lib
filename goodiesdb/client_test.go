package goodiesdb

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"
)

// Mock server to test the client
func startMockServer(t *testing.T) net.Listener {
	ln, err := net.Listen("tcp", "localhost:6379")
	if err != nil {
		t.Fatalf("Failed to start mock server: %v", err)
	}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}
			go handleMockConnection(conn)
		}
	}()
	return ln
}

// Handle mock server connections
func handleMockConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if strings.HasPrefix(line, "AUTH ") {
			if strings.TrimSpace(line) == "AUTH mypassword" {
				fmt.Fprintln(conn, "OK")
			} else {
				fmt.Fprintln(conn, "ERR invalid password")
			}
		} else if strings.HasPrefix(line, "SET ") {
			fmt.Fprintln(conn, "OK")
		} else if strings.HasPrefix(line, "GET ") {
			fmt.Fprintln(conn, "myvalue")
		}
	}
}

// Test client connection
func TestClientConnection(t *testing.T) {
	// ln := startMockServer(t)
	// defer ln.Close()

	_, err := NewClient("localhost:6379", "")
	if err != nil {
		t.Errorf("Failed to connect to server: %v", err)
	}
}

// Test client authentication
func TestClientAuthentication(t *testing.T) {
	// ln := startMockServer(t)
	// defer ln.Close()

	client, err := NewClient("localhost:6379", "guest")
	if err != nil {
		t.Errorf("Failed to connect to server: %v", err)
	}

	err = client.Auth("guest")
	if err != nil {
		t.Errorf("Failed to authenticate: %v", err)
	}
}

// Test SET command
func TestClientSet(t *testing.T) {
	// ln := startMockServer(t)
	// defer ln.Close()

	client, err := NewClient("localhost:6379", "guest")
	if err != nil {
		t.Errorf("Failed to connect to server: %v", err)
	}

	err = client.Set("mykey", "myvalue")
	if err != nil {
		t.Errorf("Failed to set key-value: %v", err)
	}
}

// Test GET command
func TestClientGet(t *testing.T) {
	//ln := startMockServer(t)
	//defer ln.Close()

	client, err := NewClient("localhost:6379", "guest")
	if err != nil {
		t.Errorf("Failed to connect to server: %v", err)
	}

	value, err := client.Get("mykey")
	if err != nil {
		t.Errorf("Failed to get value: %v", err)
	}

	if value != "myvalue" {
		t.Errorf("Expected myvalue, got %s", value)
	}
}
