package main

import (

   log "github.com/sirupsen/logrus"
   "github.com/asccclass/filewatch/libs/filemonitor"
)

func main() {
   nightswatch, err := filemonitor.NewMonitor("./tmp", "")
   if err != nil {
      log.Println(err)
      return
   }
   log.Println("Night gathers, and now my watch begins,It shall not end until my death....")
   nightswatch.Do()
   nightswatch.FileWatcher.Watch("./filepool")
   select {}
}
