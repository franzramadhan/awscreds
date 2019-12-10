package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = []string{"stscreds", "-u", "https://dummyserver.com", "-r", "arn:aws:iam::dummy", "-e", "123123123123123123"}
	main()
}
