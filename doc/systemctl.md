## systemctl

daemon process by systemctl

/lib/systemd/system/temp-gin-api-self.service

```conf
[Unit]
Description=temp-gin-api-self web service
After=network-online.target network.target syslog.target
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

when update config use  `supervisorctl update temp-gin-api-self`

- Effective

```sh
# 运行日志
journalctl -u temp-gin-api-self
# 查看服务状态
sudo systemctl status temp-gin-api-self

# 修改配置后需要
sudo systemctl daemon-reload
sudo systemctl restart temp-gin-api-self

# 启动测试
sudo systemctl start temp-gin-api-self
# 停止
sudo systemctl stop temp-gin-api-self
# 测试通过打开开机自启动
sudo systemctl enable temp-gin-api-self

# update
cd [version]
cp config.yaml ~/api/api-swtich-subscription/
supervisorctl stop temp-gin-api-self
cp [new file] ~/api/api-swtich-subscription/
supervisorctl update temp-gin-api-self
```

## 日志设置

```bash
# 查看日志服务
systemctl list-units | grep journal*
```