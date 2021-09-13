package main

import "testing"

func test_main(t *testing.T) {
	ts := mockServer{}
	ts.Add()
	ts.Start()
}
