package util

// EnvCHParams extracts clickhouse connection parameters from CLICKHOUSE environment variable
func EnvCHParams(connParams string) CHParams {
	return ExtractCHParams(connParams)
}
