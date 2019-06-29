# simple Auth

- at `${project_root}/config` package

```go
type AuthCfg struct {
	SuperAccessToken string
}

var authCfg *AuthCfg

func GetAuthCfg() *AuthCfg {
	return authCfg
}

// in config.yaml must set
//	auth:
//		super_access_token: adadaddatrsgrwfgaf
// or will return error
func initAuthCfg() error {
	superAccessToken := viper.GetString("auth.super_access_token")
	if superAccessToken == "" {
		return fmt.Errorf("auth.super_access_token is empty")
	}
	authCfg = &AuthCfg{
		SuperAccessToken: superAccessToken,
	}
	return nil
}
```

and init after viper

```go
	err := initAuthCfg()
	if err != nil {
		return err
	}
```

# use

```go
func HandlerFunc(c *gin.Context) {
	err := auth.GinUnAuthCheck(c)
	if err != nil {
		// not pass auth
	}
}
```