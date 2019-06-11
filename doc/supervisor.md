## 驻守

/etc/supervisor/conf.d/api-temp.conf

```conf
[program:api-temp]
autorestart=true
redirect_stderr=false
command=/root/Document/api-temp/1.0.0/main -c /root/Document/api-temp/1.0.0/config.yaml
stdout_logfile_maxbytes = 20MB
stdout_logfile_backups = 49
stdout_logfile = /root/Document/api-temp/log/supervisor_stdout.log
stderr_logfile_maxbytes = 1MB
stderr_logfile_backups = 30
stderr_logfile = /root/Document/api-temp/log/supervisor_stderr.log
```

redirect_stderr=true 如果为 true ，则stderr的日志会被写入stdout日志文件中 ; 默认为 false ，非必须设置

每次更新配置后 `supervisorctl update api-temp`

- 生效

```sh
mkdir -p /root/Document/api-temp/log/
# service supervisor restart
supervisorctl update

# check
supervisorctl status
tail -n 40 /root/Document/api-temp/log/supervisor_stdout.log

tail -n 40 /root/Document/api-temp/log/supervisor_stderr.log

# see more info
tail -f -n 30 /root/Document/api-temp/log/supervisor_stdout.log

# update
cd [version]
cp config.yaml ~/Document/api-temp/
supervisorctl stop api-temp
cp main ~/Document/api-temp/
supervisorctl update api-temp
```

# supervisor

## centOS

```bash
yum check-update
yum install epel-release

yum search supervisor
yum install -y supervisor

systemctl enable supervisord
systemctl start supervisord

systemctl status supervisord
```

## ubuntu

```bash
apt-get install -y supervisord
```