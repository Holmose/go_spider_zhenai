安装elasticsearch

```bash
docker run -d --name elasticsearch -p 9200:9200 elasticsearch:5.6.6
```

go 语言中使用

```bash
go get gopkg.in/olivere/elastic.v5
# 下载中文分词
https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v5.6.6/elasticsearch-analysis-ik-5.6.6.zip

# 解压到 Elasticsearch 安装目录的 /plugins/目录下即可，并改名为 ik
root@caf6be5603e3:~# ls -lh /usr/share/elasticsearch/plugins/ik/
total 1.4M
-rw-r--r--. 1 root root 258K Aug 13 14:30 commons-codec-1.9.jar
-rw-r--r--. 1 root root  61K Aug 13 14:30 commons-logging-1.2.jar
drwxr-xr-x. 2 root root 4.0K Aug 13 14:30 config
-rw-r--r--. 1 root root  51K Aug 13 14:30 elasticsearch-analysis-ik-5.6.6.jar
-rw-r--r--. 1 root root 720K Aug 13 14:30 httpclient-4.5.2.jar
-rw-r--r--. 1 root root 320K Aug 13 14:30 httpcore-4.4.4.jar
-rw-r--r--. 1 root root 2.7K Aug 13 14:30 plugin-descriptor.properties

# 重启服务
docker restart elasticsearch
```

