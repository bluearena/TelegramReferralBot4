package main

import(
	"github.com/segmentio/ksuid"
	"os"
	"log"
	"io/ioutil"
	"bytes"
	"encoding/json"
)

func generateToken() string{
	return ksuid.New().String()
}

func readJson(obj interface{}, filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		log.Print(err)
	}
	body, err := ioutil.ReadAll(file)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(obj)
	if err != nil {
		log.Print(err)
	}
	file.Close()
}