# go-redis

- [github.com/go-redis/redis](https://github.com/go-redis/redis)

```bash
dep ensure -add github.com/go-redis/redis@6.15.3
dep ensure -v
```

```go
import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
	)
const (
	defaultRedisMaxRetries = 0
	defaultDialTimeout     = 5 * time.Second
	defaultReadTimeout     = 3 * time.Second
	defaultWriteTimeout    = defaultReadTimeout
)

// init redis client
//	addr as localhost:6379
//	pwd if is "" will not use password
//	db redis db, default use 0
// use at code
//	initRedisClient(addr, password, db)
func initRedisClient(addr, pwd string, db int) (error, *redis.Client) {
	var opts *redis.Options
	if pwd == "" {
		opts = &redis.Options{
			Addr:         addr,
			DB:           db,
			MaxRetries:   defaultRedisMaxRetries,
			DialTimeout:  defaultDialTimeout,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
		}
	} else {
		opts = &redis.Options{
			Addr:         addr,
			Password:     pwd,
			DB:           db,
			MaxRetries:   defaultRedisMaxRetries,
			DialTimeout:  defaultDialTimeout,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
		}
	}
	redisClient := redis.NewClient(opts)
	_, err := redisClient.Ping().Result()
	if err != nil {
		return err, nil
	}
	return nil, redisClient
}

// init RedisByViper
// In config.yaml you can use as
//	redis:
//		addr: 127.0.0.1:6379
//		db: 1
//		is_no_pwd: false
//		password: 5290fb79983a8505
//	Warning!
//	if is_no_pwd true, wil not use password
//	if is_no_pwd false, will check redis.password not empty
func initRedisByViper() (error, *redis.Client) {
	addr := viper.GetString("redis.addr")
	if addr == "" {
		return fmt.Errorf("config file has not string at [ redis.addr ]"), nil
	}
	db := viper.GetInt("redis.db")
	noPwd := viper.GetBool("redis.is_no_pwd")
	if noPwd {
		return initRedisClient(addr, "", db)
	} else {
		password := viper.GetString("redis.password")
		if password == "" {
			return fmt.Errorf("config file [ is_no_pwd ] is true, but has not string at [ redis.password ]"), nil
		}
		return initRedisClient(addr, password, db)
	}
}
```

use see [https://github.com/go-redis/redis#quickstart](https://github.com/go-redis/redis#quickstart)