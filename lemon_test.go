package lemon

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	client *Client
}

func TestSuite(t *testing.T) {
	client := New(os.Getenv("LEMON_KEY"), WithDebug(true))
	if client.key == "" {
		t.SkipNow()
	}

	suite.Run(t, &Suite{client: client})
}
