package config

import (
  "testing"
  "os"
  "fmt"
)

func createFile(content []byte) (testFile string) {
  testFile = "/tmp/goTestGoConfig"
  fd, err := os.Create(testFile)
  defer fd.Close()
  if err != nil {
    fmt.Println(err)
  }
  _, err = fd.Write(content)
  if err != nil {
    fmt.Println(err)
  }
  return
}

func TestParseJson(t *testing.T) {
  type confStruct struct {
    Name string
    Age int
    Parents []string
  }
  jsonSample := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
  testFile := createFile(jsonSample)
  defer os.Remove(testFile)

  var solution confStruct
  err := ParseJson(testFile, &solution)
  if err != nil {
    t.Error(err)
  }
  if solution.Name != "Wednesday" {
    t.Error("Parse json fail, should be Wednesday, get ", solution.Name)
  }

  if solution.Parents[0] != "Gomez" {
    t.Error("Parse json fail, should be Gomez, get ", solution.Parents)
  }
}

func TestParseYaml(t *testing.T) {
  type confStruct struct {
    A string
    B struct {
      C int
      D []int
    }
  }
  yamlSample := []byte(`
a: Hello
b:
  c: 2
  d: [3, 4]
  `)
  testFile := createFile(yamlSample)
  defer os.Remove(testFile)

  var solution confStruct
  err := ParseYaml(testFile, &solution)
  if err != nil {
    t.Error(err)
  }

  if solution.A != "Hello" {
    t.Error("Parse Yaml fail, should be Hello ,get", solution)
  }
  if solution.B.D[0] != 3 {
    t.Error("Parse Yaml fail, should be 3 ,get", solution)
  }
}

func TestReadFile(t *testing.T) {
  content := []byte("hello")
  testFile := createFile(content)
  defer os.Remove(testFile)

  rawFile, err := readFile(testFile)
  if err != nil {
    t.Error(err)
  }
  for k, v := range content {
    if v != rawFile[k] {
      t.Error("readFile error, get ", rawFile)
    }
  }
}

