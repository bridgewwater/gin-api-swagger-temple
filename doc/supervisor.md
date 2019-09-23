## supervisor

daemon process by supervisor 

/etc/supervisor/conf.d/temp-gin-api-self.conf

```conf
[program:temp-gin-api-self]
autorestart=true
redirect_stderr=false
command=/root/Document/temp-gin-api-self/1.0.0/main -c /root/Document/temp-gin-api-self/1.0.0/config.yaml
stdout_logfile_maxbytes = 20MB
stdout_logfile_backups = 49
stdout_logfile = /root/Document/temp-gin-api-self/log/supervisor_stdout.log
stderr_logfile_maxbytes = 1MB
stderr_logfile_backups = 30
stderr_logfile = /root/Document/temp-gin-api-self/log/supervisor_stderr.log
```

redirect_stderr=true 如果为 true ，则stderr的日志会被写入stdout日志文件中 ; 默认为 false ，非必须设置

when update config use `supervisorctl update temp-gin-api-self`

- Effective

```bash
mkdir -p /root/Document/temp-gin-api-self/log/
# service supervisor restart
supervisorctl update

# check
supervisorctl status
tail -n 40 /root/Document/temp-gin-api-self/log/supervisor_stdout.log

tail -n 40 /root/Document/temp-gin-api-self/log/supervisor_stderr.log

# see more info
tail -f -n 30 /root/Document/temp-gin-api-self/log/supervisor_stdout.log

# update
cd [version]
cp config.yaml ~/Document/temp-gin-api-self/
supervisorctl stop temp-gin-api-self
cp main ~/Document/temp-gin-api-self/
supervisorctl update temp-gin-api-self
```

# supervisor-install

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