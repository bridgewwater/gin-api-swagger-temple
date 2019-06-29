package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
)

var urlCheckList = [...]string{
	"/v1/api/checked",
}

var urlMatchList = [...]string{
	"/v1/api/checked/id/.*",
}

var superAccessToken = "557c0689faf80e61daeee6e343"

// this method use HEAD x-access-token
func GinUnAuthCheck(c *gin.Context) error {
	authorization := c.GetHeader("x-access-token")
	if authorization == "" {
		return fmt.Errorf("no head x-access-token")
	}
	//TODO xAccessToken := config.GetAuthCfg().SuperAccessToken
	xAccessToken := superAccessToken
	url := c.Request.URL.String()

	for _, val := range urlCheckList {
		if url == val {
			if authorization == xAccessToken {
				return nil
			}
		}
	}

	for _, matchStr := range urlMatchList {
		matched, _ := regexp.MatchString(matchStr, url)
		if matched {
			if authorization == xAccessToken {
				return nil
			}
		}
	}

	return fmt.Errorf("other check, please fix Auth check error")
}

func GinUnAuthHead(c *gin.Context, key, val string) error {
	xAccessToken := c.GetHeader(key)
	if xAccessToken == val {
		return nil
	}
	return fmt.Errorf("can not pass auth")
}
