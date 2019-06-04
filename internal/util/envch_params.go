package util

import (
	"os"

	"github.com/sirkon/message"
)

// EnvCHParams extracts clickhouse connection parameters from CLICKHOUSE environment variable
func EnvCHParams() CHParams {
	CH := os.Getenv("CLICKHOUSE")
	if len(CH) == 0 {
		CH = "default@localhost:8123/default"
		message.Warningf("No CLICKHOUSE environment variable found, will use `%s` to connect to clickhouse", CH)
	}
	return ExtractCHParams(CH)
}
