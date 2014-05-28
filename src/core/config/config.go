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

var c Config = parse()

// json parser
func parse() Config {
  var c Config
  var filename string

  filename = "config/" + martini.Env + ".json"
  log.Println("config file: " + filename)

  jsonString, err := ioutil.ReadFile(filename)
  if err != nil {
    log.Println("ioutil.ReadFile error file: " + filename)
  }

  err = json.Unmarshal(jsonString, &c)
  if err != nil {
    log.Println("json.Unmarshal error file: " + filename)
  }

  return c
}

func Get() Config {
  return c
}
