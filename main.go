package main

import (
	"flag"
	"fmt"
	"github.com/tawaku/zbxcui/api"
	"github.com/tawaku/zbxcui/gui"
	"log"
	"os"
)

const (
	DefaultLogFile = "logs/app.log"
)

func main() {
	// Get arguments
	var host, user, pass string
	var port int
	flag.StringVar(&host, "host", "127.0.0.1", "IP address of Zabbix server.")
	flag.IntVar(&port, "port", 80, "Port of Zabbix server.")
	flag.StringVar(&user, "user", "Admin", "Username to login Zabbix GUI.")
	flag.StringVar(&pass, "pass", "zabbix", "Password to login Zabbix GUI.")
	flag.Parse()

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
	url := fmt.Sprintf("http://%s:%d/api_jsonrpc.php", host, port)
	if c, err := api.MakeClient(url, user, pass, logger); err != nil {
		panic("Failed to make client: " + err.Error())
	} else {
		defer c.Close()
		d = gui.NewDashboard(c, logger)
	}

	// Start GUI
	d.Run()
}
