package util

import (
	"os"

	"github.com/sirkon/message"

	"github.com/sirkon/ch-encode/internal/util"
)

// EnvCHParams extracts clickhouse connection parameters from CLICKHOUSE environment variable
func EnvCHParams() util.CHParams {
	CH := os.Getenv("CLICKHOUSE")
	if len(CH) == 0 {
		CH = "default@localhost:8123/default"
		message.Warningf("No CLICKHOUSE environment variable found, will use `%s` to connect to clickhouse", CH)
	}
	return util.ExtractCHParams(CH)
}
