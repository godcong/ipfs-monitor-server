package main

import (
	"flag"
	"github.com/godcong/elogrus"
	"github.com/godcong/ipfs-monitor-server/config"
	"github.com/godcong/ipfs-monitor-server/service"
	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var configPath = flag.String("config", "config.toml", "config path")
var logPath = flag.String("log", "", "log path")

func main() {
	flag.Parse()

	initLog(*logPath)

	err := config.Initialize(os.Args[0], *configPath)
	if err != nil {
		panic(err)
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//start
	service.Start(config.Config())

	go func() {
		sig := <-sigs
		log.Println(sig, "exiting")
		service.Stop()
		done <- true
	}()
	<-done

}

func initLog(logPath string) {

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200"))
	if err != nil {
		log.Panic(err)
	}

	t, err := elogrus.NewElasticHook(client, "localhost", log.TraceLevel, "ipfs-monitor-server")
	if err != nil {
		log.Panic(err)
	}
	log.AddHook(t)

	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{})

	if logPath != "" {
		dir, _ := filepath.Split(logPath)
		_ = os.MkdirAll(dir, os.ModePerm)
		file, err := os.OpenFile(logPath, os.O_SYNC|os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			log.Error(err)
		} else {
			log.SetOutput(file)
		}
	}
}
