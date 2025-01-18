## supervisor

daemon process by supervisor 

/etc/supervisor/conf.d/gin-api-swagger-temple.conf

```conf
[program:gin-api-swagger-temple]
autorestart=true
redirect_stderr=false
command=/root/Document/gin-api-swagger-temple/1.0.0/main -c /root/Document/gin-api-swagger-temple/1.0.0/config.yaml
stdout_logfile_maxbytes = 20MB
stdout_logfile_backups = 49
stdout_logfile = /root/Document/gin-api-swagger-temple/log/supervisor_stdout.log
stderr_logfile_maxbytes = 1MB
stderr_logfile_backups = 30
stderr_logfile = /root/Document/gin-api-swagger-temple/log/supervisor_stderr.log
```

redirect_stderr=true 如果为 true ，则stderr的日志会被写入stdout日志文件中 ; 默认为 false ，非必须设置

when update config use `supervisorctl update gin-api-swagger-temple`

- Effective

```bash
mkdir -p /root/Document/gin-api-swagger-temple/log/
# service supervisor restart
supervisorctl update

# check
supervisorctl status
tail -n 40 /root/Document/gin-api-swagger-temple/log/supervisor_stdout.log

tail -n 40 /root/Document/gin-api-swagger-temple/log/supervisor_stderr.log

# see more info
tail -f -n 30 /root/Document/gin-api-swagger-temple/log/supervisor_stdout.log

# update
cd [version]
cp config.yaml ~/Document/gin-api-swagger-temple/
supervisorctl stop gin-api-swagger-temple
cp main ~/Document/gin-api-swagger-temple/
supervisorctl update gin-api-swagger-temple
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