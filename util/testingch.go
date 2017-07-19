package util

import (
	"fmt"
	"net/http"
	"os"

	"github.com/glossina/message"
)

// CHParams clickhouse connection coordinates
type CHParams struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// URL for clickhouse connection
func (ch CHParams) URL() string {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d", ch.Host, ch.Port), nil)
	if err != nil {
		panic(err)
	}
	q := req.URL.Query()
	if len(ch.User) > 0 {
		q.Set("user", ch.User)
	}
	if len(ch.Password) > 0 {
		q.Set("password", ch.Password)
	}
	if len(ch.DBName) > 0 {
		q.Set("database", ch.DBName)
	}
	req.URL.RawQuery = q.Encode()
	return req.URL.String()
}

// DBURL is a URL generator for Mail.RU's clickhouse connector
func (ch CHParams) DBURL() string {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/%s", ch.Host, ch.Port, ch.DBName), nil)
	if err != nil {
		panic(err)
	}
	q := req.URL.Query()
	if len(ch.User) > 0 {
		q.Set("user", ch.User)
	}
	if len(ch.Password) > 0 {
		q.Set("password", ch.Password)
	}
	req.URL.RawQuery = q.Encode()
	return req.URL.String()
}

//go:generate logparsergen --source=testingch.script --package=util

// ExtractCHParams extracts CH connection components from string
// [user[:password]@]host:port[/dbname]
func ExtractCHParams(data string) (res CHParams) {
	c := &components{}
	up := &userpass{}

	if ok, err := c.Parse([]byte(data)); !ok || err != nil {
		panic(err)
	}

	if ok, err := up.Parse(c.GetAuthData()); !ok || err != nil {
		panic(err)
	}

	res.Host = string(c.Host)
	res.Port = int(c.Port)
	res.DBName = string(c.DBName)
	res.User = string(up.User)
	res.Password = string(up.Password)
	return
}

// EnvCHParams extracts clickhouse connection parameters from CLICKHOUSE environment variable
func EnvCHParams() CHParams {
	CH := os.Getenv("CLICKHOUSE")
	if len(CH) == 0 {
		CH = "default@localhost:8123/default"
		message.Warningf("No CLICKHOUSE environment variable found, will use `%s` to connect to clickhouse", CH)
	}
	return ExtractCHParams(CH)
}
