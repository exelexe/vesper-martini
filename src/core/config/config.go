package config

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

    filename = "config/" + martini.Env + ".json"
    log.Println("config file: " + filename)

    jsonString, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Println("ioutil.ReadFile error file: " + filename)
    }

    err = json.Unmarshal(jsonString, &Conf)
    if err != nil {
        log.Println("json.Unmarshal error file: " + filename)
    }
}
