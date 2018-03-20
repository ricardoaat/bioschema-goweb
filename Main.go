package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ricardoaat/bioschemas-goweb/bioparser"
	"github.com/ricardoaat/bioschemas-goweb/config"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

var (
	version   string
	buildDate string
)

func logInit() {

	//now := time.Now()
	//logfile := config.Conf.Path.LogPath + fmt.Sprintf("govalid_%s.log", now.Format("20060102T150405"))
	logfile := config.Conf.Path.LogPath + "goweb.log"
	fmt.Println("Loging to " + logfile)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{})
	pathMap := lfshook.PathMap{
		log.DebugLevel: logfile,
		log.InfoLevel:  logfile,
		log.ErrorLevel: logfile,
		log.WarnLevel:  logfile,
		log.PanicLevel: logfile,
	}
	log.AddHook(lfshook.NewHook(
		pathMap,
		&log.JSONFormatter{},
	))
}

func main() {

	err := config.LoadConfig("config.toml")
	if err != nil {
		fmt.Println("Couldn't load config.toml ", err)
	}
	logInit()

	log.Info("--------------Init program--------------")
	log.Info(fmt.Sprintf("Version: %s Build Date: %s", version, buildDate))
	log.Debug("Loaded configuration " + fmt.Sprint(config.Conf))
	v := flag.Bool("v", false, "Returns the binary version and built date info")
	f := flag.String("f", "", "File path to Bioschemas CSV file to parse")
	u := flag.String("u", "", "Url Path to Bioschemas CSV info to parse")

	flag.Parse()

	if !*v {
		bioparser.Validate(*f, *u)
	}
}
