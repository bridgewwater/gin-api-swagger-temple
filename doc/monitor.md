## config

```yaml
monitor: # monitor 自检测
  count: 10           # pingServer 函数try的次数
  status: true        # 存活检查接口 建议常开
  health: /status/health # api health
  hardware: true      # 硬件信息检查 按需开放
  status_hardware:
    disk: /status/hardware/disk     # hardware api disk
    cpu: /status/hardware/cpu       # hardware api cpu
    ram: /status/hardware/ram       # hardware api ram
  debug: true         # 调试接口，按需开放
  pprof: true         # 性能检测，按需开放
```

- `http://127.0.0.1:38080`  根据实际情况更换

这组 API 是为了检查服务器状态

## health

```bash
curl http://127.0.0.1:38080/status/health \
	-X GET
```

## disk

```bash
curl http://127.0.0.1:38080/status/hardware/disk \
	-X GET
```

## cpu

```bash
curl http://127.0.0.1:38080/status/hardware/cpu \
	-X GET
```

## ram

```bash
curl http://127.0.0.1:38080/status/hardware/ram \
	-X GET
```