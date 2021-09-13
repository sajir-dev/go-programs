package main

import "fmt"

type serverInterface interface {
	Add()
	Start()
}

type Server struct {
	port int
}

func (s *Server) Add() {
	fmt.Println("Add Ran")
}

func (s *Server) Start() {
	fmt.Println("Started")
}

type mockServer struct{}

func (m *mockServer) Add() {
	fmt.Println("Mock Add ran")
}

func (m *mockServer) Start() {
	fmt.Println("Mock started")
}

func main() {
	s := Server{
		port: 3000,
	}

	sm := mockServer{}

	sm.Start()
	sm.Add()

	s.Start()
	s.Add()
}
