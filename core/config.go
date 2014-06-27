package core

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"io/ioutil"
	"log"
)

type Config struct {
	Dsn string
}

var Conf Config

// json parser
func init() {
	var err error
	var filename string

	filename = "./config/" + martini.Env + ".json"
	log.Println("config file: " + filename)

	jsonString, err := ioutil.ReadFile(filename)
	CheckFatal(err, "read error")

	err = json.Unmarshal(jsonString, &Conf)
	CheckFatal(err, "parse error")
}
