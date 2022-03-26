# Antioxidant
更新一些想看的东西

# Features

* 监控 github 仓库并发送到企业微信上

# TODO

* [x] 查询新增的CVE
* [ ] 监控一些常用站点，比如expdb等
* [ ] 支持更多的企微消息格式
* [ ] 支持更多种接收形式
* [ ] 协程加速

# 使用

1. 修改 antioxidant.go 的15行 github token，18行的企业微信 webhook api。
2. 修改 repos/Repos.yaml 中想要监控的仓库。
3. 运行。
