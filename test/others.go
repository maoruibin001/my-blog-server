package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"my-blog-server/src/db"
	"my-blog-server/src/utils"
	"strings"
)


func generationAid() int {
	counter, session := db.GetCollect("my-blog-2", "counter")
	defer session.Close()

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"aid": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	doc := struct{ Aid int }{}
	if _, err := counter.Find(bson.M{}).Apply(change, &doc); err != nil {
		utils.HandleError("查找出错", err)
	}
	log.Println("doc:", doc)
	return doc.Aid
}
func main() {
	var a = ""

	fmt.Println(strings.Split(a, ","))
}
