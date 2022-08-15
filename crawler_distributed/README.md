### 使用命令行创建服务

#### 启动worker
```bash
go run .\crawler_distributed\worker\server\worker.go --port=9000
go run .\crawler_distributed\worker\server\worker.go --port=9001
```

#### 启动ItemSaver
```bash
go run .\crawler_distributed\persist\server\itemsaver.go --port=1234
```

#### 启动爬虫
```bash
go run .\crawler_distributed\main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9
001"
```