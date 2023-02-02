package main

import (
	"fmt"
	"log"
	"time"
	"io/ioutil"
	"encoding/json"

	mega "github.com/wade-welles/go-mega"
)


type Client struct {
	config  *Config
	mega    *mega.Mega
}

type Config struct {
	BaseUrl         string
	Retries         int
	DownloadWorkers int
	UploadWorkers   int
	TimeOut         int
	User            string
	Password        string
	Recursive       bool
	Force           bool
	SkipSameSize    bool
	SkipError       bool
	Verbose         int
}

type Path struct {
	prefix string
	path   []string
	size   int64
	t      int
	ts     time.Time
}

func (cfg *Config) Parse(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, cfg)
	if err != nil {
		panic(err)
	}

	return nil
}

func (client *Client) Login(otp string) error {
	err := client.mega.Login(client.config.User, client.config.Password, otp)
	return err
}

func main() {
	fmt.Println("mega-cli")
	fmt.Println("========")

	config := new(mega.Config)

	client := &Client{
		config:  config,
		mega: mega.New(),
	}

 mega.

	err = client.Login(otp)
	if err != nil {
		if err == mega.ENOENT {
			log.Fatal("Login failed, Please verify username or password")
		} else {
			log.Fatal("Unable to establish connection to mega service")
		}
	}
}
