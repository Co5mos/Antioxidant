# Antioxidant
更新一些想看的东西

# Features

* 监控 github 仓库并发送到企业微信上
* 查询新增的CVE
* 热点漏洞监控，比如 spring rce, spring4shell

# TODO

* [x] 查询新增的CVE
* [x] 热点漏洞监控
* [ ] 监控一些常用站点，比如expdb等
* [ ] 支持更多的企微消息格式
* [ ] 支持更多种接收形式

# 使用

1. 修改 config.yaml 中的 github token 和 webhook api。
2. 修改 repos/Repos.yaml 中想要监控的仓库。
3. 修改 repos/HotKey.yaml 中想要监控的漏洞关键字。
4. 运行。
