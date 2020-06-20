package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("HTTPRESTAPI_DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=test password=123 dbname=test_test sslmode=disable"
	}

	os.Exit(m.Run())
}
