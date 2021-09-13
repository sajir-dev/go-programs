package main

import "fmt"

type mockServer struct{}

func (m *mockServer) Add() {
	fmt.Println("Mock Add ran")
}

func (m *mockServer) Start() {
	fmt.Println("Mock started")
}
