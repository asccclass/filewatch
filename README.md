## 監控特定目錄內檔案是否有異動程式


### 安裝
```
go mod download  // 取得函式
```


### 使用函數
```
https://github.com/howeyc/fsnotify
```

### 升級函數

```
檢查： go list -m -u all
升級： go get <套件名稱>@latest (針對您需要的套件)
清理： go mod tidy
測試： go test ./... (確保升級沒把程式弄壞)
```
