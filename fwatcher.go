package main

import (
   "log"
   "github.com/asccclass/filewatch/libs/filemonitor"
)

func main() {
   M, err := filemonitor.NewMonitor("./tmp")
   if err != nil {
      log.Println(err)
      return
   }
   M.Do()
   M.FileWatcher.Watch("./filepool")
   select {}
}
