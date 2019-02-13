package main

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)

}

func getRandStr(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	bytes := []byte(str)

	ret := []byte{}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l ; i ++ {

		ret = append(ret, bytes[rand.Intn(len(str))])
	}

	return string(ret)
}
func MD5(text string) string{
	ctx := md5.New()
	ctx.Write([]byte(text))

	return hex.EncodeToString(ctx.Sum(nil))
}

func MD52(text, salt string) string{
	ctx := md5.New()
	ctx.Write([]byte(text))

	saltByte := []byte(salt)
	return hex.EncodeToString(ctx.Sum(saltByte))
}

func main() {

	//fmt.Println(getRandStr(10))
	//
	//var b = getRandStr(3)
	//var a = ""
	//fmt.Println([]byte(a) == nil)
	//
	//fmt.Println(MD5(b + "hel"))
	//fmt.Println(MD52(b, "hel"))
	//var a = "34234madfa"
	//
	//fmt.Println(a[2:])
	//
	//fmt.Println(time.Now().Add(1 * time.Hour).Unix())

	//fmt.Println("" || "hello")

}
