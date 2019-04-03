# APIBase

[TOC]

- `http://127.0.0.1:38080`  根据实际情况更换

这组 API 是为了检查服务器状态

## health

```bash
curl http://127.0.0.1:38080/ssc/health \
	-X GET
```

## disk

```bash
curl http://127.0.0.1:38080/ssc/disk \
	-X GET
```

## cpu

```bash
curl http://127.0.0.1:38080/ssc/cpu \
	-X GET
```

## ram

```bash
curl http://127.0.0.1:38080/ssc/ram \
	-X GET
```