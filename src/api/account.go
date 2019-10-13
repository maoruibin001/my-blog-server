package api

import (
	"album-server/src/config"
	"album-server/src/db"
	"album-server/src/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

//https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx95bd67c3b7c5c989&redirect_uri=https%3A%2F%2F88w87.com&response_type=code&scope=snsapi_userinfo#wechat_redirect
//https://api.weixin.qq.com/sns/oauth2/access_token?appid=wx95bd67c3b7c5c989&secret=2b5bf66816586287281b22c323c11743&code=021Zz77M1K9IH7132z4M1UTN6M1Zz77c&grant_type=authorization_code
//
//https://api.weixin.qq.com/sns/userinfo?access_token=26_JW9lagV1yfAZkWWJkvAakp-ot2Pid9I3NUkdDnUVdRz6iLxxH1VKd0LAUE2O894sq1Q8qyrQ-4WxJ8M7fy9Ip85Q0cjGUH44hQvZvqvg-sY&openid=oym6WuPQYckGFN0PMq1xaAha-8iU&lang=zh_CN

type ConfigStruct struct {
	AccessToken    string   `json:"access_token"`
	ExpireIn       int      `json:"expires_in"`
	RefreshToken   string   `json:"refresh_token"`
	OpenId 		   string `json:"openid"`
	Scope          string   `json:"scope"`
}

func initAccount(router *gin.Engine) {
	//通过code获取微信用户信息
	router.GET("/api/userInfo", func(context *gin.Context) {
		code := context.DefaultQuery("code", "")
		uidStr := context.DefaultQuery("uid", "")
		uid, err := strconv.Atoi(uidStr)
		if err != nil {
			utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
				"errorMsg": "uid错误，请输入正确的uid: " + uidStr,
			})
			return

		}
		if (uid == 0) {
			//获取accesstoken
			accessTokenUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", config.APPID, config.SECRET, code)
			jsonStr := utils.Get(accessTokenUrl)
			var config ConfigStruct
			if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
				if config.OpenId == "" {
					utils.ResponseJson(context, http.StatusOK, utils.RESPONSEPARAMERROR, gin.H{
						"errorMsg": "code已过期, code: " + code,
					})
					return
				}

				fmt.Println("================json str 转struct==")
				fmt.Println(config)
				fmt.Println(config.AccessToken)
			}

			//获取用户信息
			userInfoUrl := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", config.AccessToken, config.OpenId)

			fmt.Println("type is: ", reflect.TypeOf(uidStr))

			userInfoStr := utils.Get(userInfoUrl)

			fmt.Println(userInfoStr)

			var userInfo db.WxAccountSchema
			if err := json.Unmarshal([]byte(userInfoStr), &userInfo); err == nil {
				fmt.Println("================json str 转struct==")
				fmt.Println(userInfo)
				fmt.Println(userInfo.NickName)
			}

			fmt.Println("userInfoStr: ", userInfo)
			if userInfo.OpenId == "" {
				utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOUSER, nil)
				return
			}
			accountInfo := db.FindByOpenId(userInfo.OpenId)

			if accountInfo.OpenId == "" { //系统中没有这个用户，则创建用户
				accountInfo = db.CreateAccount(userInfo)
				utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, accountInfo)
			} else if accountInfo.OpenId == userInfo.OpenId { //系统中有此用户，则返回用户
				utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, accountInfo)
			} else { //系统中已存在openid对应的其他用户，则返回用户已存在
				utils.ResponseJson(context, http.StatusOK, utils.RESPONSEUSEREXSIST, accountInfo)
			}

		} else {
			accountInfo := db.AccountSingleFindByKV("uid", uid)
			if accountInfo.UId == uid {
				utils.ResponseJson(context, http.StatusOK, utils.RESPONSEOK, accountInfo)
				return
			} else {
				utils.ResponseJson(context, http.StatusOK, utils.RESPONSENOUSER, nil)
			}
		}

	})
}
