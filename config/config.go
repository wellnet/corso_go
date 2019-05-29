package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

type Config struct {
	ForecastDownloader string `json:"forecastDownloader"`
}

var (
	c Config
)

func Read() {
	err := read()
	if err != nil {
		panic(err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	if err := watcher.Add("config.json"); err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case <-watcher.Events:
				err := read()
				if err != nil {
					panic(err)
				}
			case err := <-watcher.Errors:
				log.Printf("Error watching config: %v", err)
			}
		}
	}()
}

func read() error {
	log.Printf("Reading config")

	jsonFile, err := os.Open("config.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		return err
	}

	return nil
}

func Get() Config {
	return c
}
