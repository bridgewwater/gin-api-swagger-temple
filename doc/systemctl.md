## systemctl

daemon process by systemctl

/lib/systemd/system/gin-api-swagger-temple.service

```conf
[Unit]
Description=gin-api-swagger-temple web service
After=network-online.target network.target syzlog.S().target
Wants=network.target

[Service]
Type=simple
# start path need change
ExecStart=/root/api/api-swtich-subscription/1.0.0/main -c /root/api/api-swtich-subscription/1.0.0/config.yaml
# auto restart open, default is no, other is always on-success
Restart=always
# auto restart time default is 0.1s
RestartSec=5
# auto restart count 0 is unlimited
StartLimitInterval=10
# if ExitStatus 143 137 SIGTERM SIGKILL will not restart
#RestartPreventExitStatus=143 137 SIGTERM SIGKILL

[Install]
WantedBy=multi-user.target
```

when update config use  `supervisorctl update gin-api-swagger-temple`

- Effective

```sh
# 运行日志
journalctl -u gin-api-swagger-temple
# 查看服务状态
sudo systemctl status gin-api-swagger-temple

# 修改配置后需要
sudo systemctl daemon-reload
sudo systemctl restart gin-api-swagger-temple

# 启动测试
sudo systemctl start gin-api-swagger-temple
# 停止
sudo systemctl stop gin-api-swagger-temple
# 测试通过打开开机自启动
sudo systemctl enable gin-api-swagger-temple

# update
cd [version]
cp config.yaml ~/api/api-swtich-subscription/
supervisorctl stop gin-api-swagger-temple
cp [new file] ~/api/api-swtich-subscription/
supervisorctl update gin-api-swagger-temple
```

## 日志设置

```bash
# 查看日志服务
systemctl list-units | grep journal*
```