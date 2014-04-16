/*
The config package is used to load config file.
It support Json/Yaml/Xml config.

Usage:
config.Parse(pathToConfigFile, codeType) // return an interface
codeType can be one of "config.Json, config.Yaml, config.Xml"
*/
package config

import (
  "os"
  "encoding/json"
  "errors"
  "gopkg.in/yaml.v1"
  "encoding/xml"
)

func ParseJson(filename string, solution interface{}) (err error) {
  rawConf, err := readFile(filename)
  if err != nil {
    return
  }
  err = json.Unmarshal(rawConf, solution)
  if err != nil {
    return
  }
  return
}

func ParseXml(filename string, solution interface{}) (err error) {
  rawConf, err := readFile(filename)
  if err != nil {
    return
  }
  err = xml.Unmarshal(rawConf, solution)
  if err != nil {
    return
  }
  return
}

func ParseYaml(filename string, solution interface{}) (err error) {
  rawConf, err := readFile(filename)
  if err != nil {
    return
  }
  err = yaml.Unmarshal(rawConf, solution)
  if err != nil {
    return
  }
  return
}

func readFile(filename string) (rawConf []byte, err error) {
  var fd *os.File
  fd, err = os.Open(filename)
  defer fd.Close()
  if err != nil {
    return
  }
  rawFile := make([]byte,10240)
  length, _ := fd.Read(rawFile)
  if length == 0 {
    err = errors.New("File length is zero")
    return
  }
  rawConf = rawFile[:length]
  return
}
