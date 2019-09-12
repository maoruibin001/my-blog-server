package middleware

import (
	"fmt"
	"strings"
	"github.com/gin-gonic/gin"
	"album-server/src/utils"
)


func SelectDb() gin.HandlerFunc {
	return func(c *gin.Context) {
		flag := c.DefaultQuery("flag", "")
		if flag == "" {
			authorization := c.Request.Header.Get("Authorization")
			if s := strings.Split(authorization, " "); len(s) == 3 {
				flag = s[2]
			}
		}
		fmt.Println("flag is: ", flag)
		utils.SetDbName(flag)
	}

}