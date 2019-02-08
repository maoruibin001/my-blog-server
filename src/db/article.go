package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"my-blog-server/src/utils"
	"strings"
	"time"
)

type ArticleSchema struct {
	Aid int `aid:"aid"`
	Title  string    `json:"title"`
	Content string	 `json:"content"`
	Tags []string `json:"tags"`
	Date int64 `json:"date"`
	IsPublish bool `json:"isPublish"`
	CommentN int `json:"commentN"`
}

func InsertArticle(data ArticleSchema) ArticleSchema {
	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()
	m := ArticleSchema{}
	m.Aid = data.Aid
	m.Title = data.Title
	m.Tags = data.Tags
	m.Date = data.Date
	m.IsPublish = data.IsPublish
	m.CommentN = data.CommentN

	utils.HandleError("insert error: ", c.Insert(&m))
	fmt.Println("插入一条数据", m)

	return m

}

func FindByArtcleId(id string) []ArticleSchema {
	return ArticleMultiFindByKV("aid", id)
}

func ArticleSingleFindByKV(key string, v interface{}) UserSchema {

	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()

	results := []UserSchema{}

	utils.HandleError("find error: ", c.Find(bson.M{key: v}).All(&results))

	result := UserSchema{}
	if len(results) > 0 {
		result = results[0]
	}
	return result
}

func ArticleMultiFindByKV( key string, v interface{}) []ArticleSchema {

	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()

	results := []ArticleSchema{}

	utils.HandleError("find error: ", c.Find(bson.M{key: v}).All(&results))


	return results
}

func generationAid() int {
	counter, session := GetCollect("my-blog-2", "counter")
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
func CreateArticle(title, tags string, isPublish, comentN int) ArticleSchema {
	var m = ArticleSchema{}
	m.Aid = generationAid()
	m.Title = title
	m.Tags = strings.Split(tags, ",")
	m.Date = time.Now().Unix()
	m.IsPublish = isPublish == 1
	m.CommentN = comentN
	return InsertArticle(m)
}

func ChangeArticle(id, name, age, phone, salt string) error {
	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()

	selector := bson.M{"id": id}
	data := bson.M{"name": name, "age": age, "phone": phone, "salt": salt}

	err := c.Update(selector, bson.M{"$set": data})

	return err
}

func RemoveArticle(k string, v interface{}) error {
	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()

	err := c.Remove(bson.M{k: v})

	return err
}
func InitArticleData()  {
	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()
	count, err := c.Count()
	utils.HandleError("查找错误：", err)
	if count == 0 {
		fmt.Println("数据库为空，初始化数据...")
		//var i int
		//for i = 0; i < 10; i ++ {
		//	//InsertUser(UserInfo{"mao", i, strconv.Itoa(i)})
		//}

	}
}
