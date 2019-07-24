package main

import (
   "os"
   log "github.com/sirupsen/logrus"
   "github.com/asccclass/filewatch/libs/filemonitor"
)

func main() {
   dir, err := os.Getwd()
   if err != nil {
      log.Printf("Initial error:%v.", err)
   }

   documentRoot := os.Getenv("DocumentRoot")
   if documentRoot == "" {
      documentRoot = "www"
   }
   documentRoot = dir + "/" + documentRoot

   nightswatch, err := filemonitor.NewMonitor(documentRoot, "")
   if err != nil {
      log.Println(err)
      return
   }
   log.Println("Night gathers, and now my watch begins,It shall not end until my death....")
   log.Printf("I'm watching %v....", documentRoot)
   nightswatch.Do()
   nightswatch.FileWatcher.Watch(documentRoot)
   select {}
}
