package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/tawaku/zbxcui/api"
	"github.com/tawaku/zbxcui/gui"
	"log"
	"os"
)

const (
	DefaultConfig  = "config.json"
	DefaultLogFile = "log/app.log"
	DefaultHost    = "127.0.0.1"
	DefaultPort    = 80
	DefaultUser    = "Admin"
	DefaultPass    = "zabbix"
)

type Config struct {
	Host string `json:"host"`
	Port uint   `json:"port,string"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

func main() {
	// Args constructor
	config := new(Config)
	// Args from json file
	if f, err := os.Open(DefaultConfig); err == nil {
		json.NewDecoder(f).Decode(&config)
		f.Close()
	}
	// Args from command line
	var host, user, pass string
	var port uint
	flag.StringVar(&host, "host", DefaultHost, "IP address of Zabbix server.")
	flag.UintVar(&port, "port", DefaultPort, "Port of Zabbix server.")
	flag.StringVar(&user, "user", DefaultUser, "Username to login Zabbix GUI.")
	flag.StringVar(&pass, "pass", DefaultPass, "Password to login Zabbix GUI.")
	flag.Parse()
	if host != DefaultHost {
		config.Host = host
	}
	if port != DefaultPort {
		config.Port = port
	}
	if user != DefaultUser {
		config.User = user
	}
	if pass != DefaultPass {
		config.Pass = pass
	}

	// Logger initialization
	var logger *log.Logger
	if logfile, err := os.OpenFile(DefaultLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666); err != nil {
		panic("Cannnot open logfile: " + err.Error())
	} else {
		defer logfile.Close()
		logger = log.New(logfile, "[ZBXCUI]", log.LstdFlags|log.LUTC)
	}

	// Make client to get information from Zabbix
	var d *gui.Dashboard
	url := fmt.Sprintf("http://%s:%d/api_jsonrpc.php", config.Host, config.Port)
	if c, err := api.MakeClient(url, config.User, config.Pass, logger); err != nil {
		panic("Failed to make client: " + err.Error())
	} else {
		defer c.Close()
		d = gui.NewDashboard(c, logger)
	}

	// Start GUI
	d.Run()
}
