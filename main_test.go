package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = []string{"awscreds-sts", "-u", "https://dummyserver.com", "-r", "arn:aws:iam::dummy"}
	main()
}
