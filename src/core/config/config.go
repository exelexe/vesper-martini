package config

import (
	"core"
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

	filename = "./config/" + martini.Env + ".jsonp"
	log.Println("config file: " + filename)

	jsonString, err := ioutil.ReadFile(filename)
	core.CheckFatal(err, "read error")

	err = json.Unmarshal(jsonString, &Conf)
	core.CheckFatal(err, "parse error")
}
