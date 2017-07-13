package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractCHParams(t *testing.T) {
	require.Equal(t,
		CHParams{
			Host:   "localhost",
			Port:   8123,
			DBName: "default",
			User:   "default",
		},
		ExtractCHParams("default@localhost:8123/default"),
	)
	require.Equal(t,
		CHParams{
			Host: "localhost",
			Port: 8123,
		},
		ExtractCHParams("localhost:8123"),
	)
	require.Equal(t,
		CHParams{
			Host:     "localhost",
			Port:     8123,
			DBName:   "default",
			User:     "default",
			Password: "bugaga",
		},
		ExtractCHParams("default:bugaga@localhost:8123/default"),
	)
}
