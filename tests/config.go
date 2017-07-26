package tests

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Username string `json:username`
	Password string `json:password`
}

func ConfigureWithFile(f *os.File) Configuration {
	j, _ := ioutil.ReadAll(f)

	var c Configuration

	if err := json.Unmarshal(j, &c); err != nil {
		panic(err)
	}

	return c
}

func Configure(f io.Reader) Configuration {
	j, _ := ioutil.ReadAll(f)

	var c Configuration

	if err := json.Unmarshal(j, &c); err != nil {
		panic(err)
	}

	return c
}
