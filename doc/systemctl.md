# systemctl

## 驻守

/lib/systemd/system/api-temp.service

```conf
[Unit]
Description=api-temp web service
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

每次更新配置后 `supervisorctl update api-temp`

- 生效

```sh
# 运行日志
journalctl -u api-temp
# 查看服务状态
sudo systemctl status api-temp

# 修改配置后需要
sudo systemctl daemon-reload
sudo systemctl restart api-temp

# 启动测试
sudo systemctl start api-temp
# 停止
sudo systemctl stop api-temp
# 测试通过打开开机自启动
sudo systemctl enable api-temp

# update
cd [version]
cp config.yaml ~/api/api-swtich-subscription/
supervisorctl stop api-temp
cp [new file] ~/api/api-swtich-subscription/
supervisorctl update api-temp
```

## 日志设置

```bash
# 查看日志服务
systemctl list-units | grep journal*
```