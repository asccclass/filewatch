package filemonitor

import (
   "log"
   "github.com/howeyc/fsnotify"
)

type monitor struct {
   Dir   string      // 監控目錄
   Ext   string      // 附檔名
   FileWatcher *fsnotify.Watcher
}

func(m *monitor) Do() {
   go func() {
      for {
         select {
         case w := <-m.FileWatcher.Event:
            if w.IsCreate() {
               log.Println("新增文件", w.Name)
               continue
            } else if w.IsModify() {
               continue
            } else if w.IsDelete() {
               log.Println(w.Name, "文件被删除.")
               continue
            } else if w.IsRename() {
               w = <-m.FileWatcher.Event
               m.FileWatcher.RemoveWatch(w.Name)
               log.Println(w.Name, " 被重命名.")
            }
         case err := <-m.FileWatcher.Error:
            log.Fatalln(err)
         }
      }
   }()
}

// Initial
func NewMonitor(dir, ext string) (*monitor, error) {
   Mon, err := fsnotify.NewWatcher()
   return &monitor {
      Dir:  dir,
      Ext:  ext,
      FileWatcher: Mon,
   }, err
}
