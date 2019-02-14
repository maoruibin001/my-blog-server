package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"my-blog-server/src/utils"
	"time"
)

type ArticleSchema struct {
	Aid int `json:"aid"`
	Title  string    `json:"title"`
	Content string	 `json:"content"`
	Tags []string `json:"tags"`
	Date int64 `json:"date"`
	IsPublish bool `json:"isPublish"`
	CommentN int `json:"commentN"`
	DateStr string `json:"dateStr"`
}

func InsertArticle(data ArticleSchema) ArticleSchema {
	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()
	m := ArticleSchema{}
	m.Aid = data.Aid
	m.Title = data.Title
	m.Tags = data.Tags
	m.Date = data.Date
	m.DateStr = data.DateStr
	m.IsPublish = data.IsPublish
	m.CommentN = data.CommentN
	m.Content = data.Content


	utils.HandleError("insert error: ", c.Insert(&m))
	fmt.Println("插入一条数据", m)

	return m

}

func FindByArtcleId(id int) []ArticleSchema {
	return ArticleMultiFindByKV("aid", id)
}

func ArticleSingleFindByKV(key string, v interface{}) ArticleSchema {

	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()

	results := []ArticleSchema{}

	utils.HandleError("find error: ", c.Find(bson.M{key: v}).All(&results))

	result := ArticleSchema{}
	if len(results) > 0 {
		result = results[0]
	}
	return result
}

func GetSomeArticles(conditions bson.M, pageSize, pageNo int)  ([]ArticleSchema, error) {
	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()

	results := []ArticleSchema{}
	var err error = nil
	err = c.Find(conditions).Limit(pageSize).Skip((pageNo - 1) * pageSize).Sort("-date").All(&results)

	fmt.Println("results:", results)
	return results, err
}


func GetArticles(key string, v interface{}, pageSize, pageNo, all int) ([]ArticleSchema, error)  {
	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()

	results := []ArticleSchema{}
	var err error = nil
	if key == "" {
		if all == 1 {
			err = c.Find(nil).All(&results)
		} else {
			err = c.Find(nil).Limit(pageSize).Skip((pageNo - 1) * pageSize).Sort("-date").All(&results)
		}
	} else {
		err = c.Find(bson.M{key: v}).Limit(pageSize).Skip((pageNo - 1) * pageSize).All(&results)
	}

	fmt.Println("results:", results)
	return results, err
}

func ArticleMultiFindByKV( key string, v interface{}) []ArticleSchema {

	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()

	results := []ArticleSchema{}

	utils.HandleError("find error: ", c.Find(bson.M{key: v}).All(&results))


	return results
}

func generationAid() int {
	counter, session := GetCollect("my-blog-2", "acounter")
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
func CreateArticle(title string, tags []string, isPublish, comentN int, content string) ArticleSchema {
	var m = ArticleSchema{}
	m.Aid = generationAid()
	m.Title = title
	m.Tags = tags
	m.Date = time.Now().UnixNano() / 1e6
	m.DateStr = time.Now().Format("2006年01月02日 15时04分05秒")
	m.IsPublish = isPublish == 1
	m.CommentN = comentN
	m.Content = content
	return InsertArticle(m)
}

func ChangeArticle(aid int, title string, tags []string, isPublish, comentN int, content string) error {
	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()

	selector := bson.M{"aid": aid}
	data := bson.M{"title": title, "tags": tags, "isPublish": isPublish, "comentN": comentN, "content": content}

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
