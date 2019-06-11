package main

import (
   "log"
   // "github.com/howeyc/fsnotify"
   "github.com/asccclass/filewatch/libs/filemonitor"
)

func main() {
   M, err := filemonitor.NewMonitor("./filepool")
   if err != nil {
      log.Println(err)
      return
   }
   M.Do()
   M.FileWatcher.Watch("./filepool")
   select {}
}
