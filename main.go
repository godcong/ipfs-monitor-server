package main

import (
	"flag"
	"github.com/godcong/go-trait"
	"github.com/godcong/ipfs-monitor-server/config"
	"github.com/godcong/ipfs-monitor-server/service"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var configPath = flag.String("config", "config.toml", "config path")
var logPath = flag.String("log", "", "log path")

func main() {
	flag.Parse()

	trait.InitElasticLog("ipfs-monitor-server", nil)

	err := config.Initialize(os.Args[0], *configPath)
	if err != nil {
		panic(err)
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	service.InitRedis(service.NewIPFSServiceRedis())
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
