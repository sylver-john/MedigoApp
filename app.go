package main

import (
  "github.com/sylver-john/Medigo/router"
	log "github.com/cihub/seelog"
  "github.com/sylver-john/Medigo/service"
)

func main() {
  service.SetLogger()
  defer log.Flush()
  log.Info("Starting the server")
  router := router.NewRouter()
  router.Run()
}
