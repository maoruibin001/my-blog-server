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
type ResData struct {
	Articles []ArticleSchema `json:"articles"`
	Count int `json:"count"`
	IsEnd int `json:"isEnd"`
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

//func FindByArtcleId(id int) []ArticleSchema {
//	return ArticleMultiFindByKV("aid", id)
//}

func ArticleSingleFindByKV(condition bson.M) ArticleSchema {

	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()

	results := []ArticleSchema{}

	fmt.Println("condition: ", condition)
	utils.HandleError("find error: ", c.Find(condition).All(&results))

	fmt.Println("result www: ", results)
	result := ArticleSchema{}
	if len(results) > 0 {
		result = results[0]
	}
	return result
}


func GetSomeArticles(conditions bson.M, pageSize, pageNo int)  (ResData, error) {
	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()


	var ret = ResData{}
	results := []ArticleSchema{}
	var err error = nil

	count, err := c.Find(conditions).Count()

	err = c.Find(conditions).Limit(pageSize).Skip((pageNo - 1) * pageSize).Sort("-date").All(&results)


	fmt.Println("results:", results)
	ret.Articles = results
	ret.Count = count
	if pageNo * pageSize >= count {
		ret.IsEnd = 1
	} else {
		ret.IsEnd = 0
	}
	return ret, err
}


func GetArticles(condition bson.M, pageSize, pageNo int) (ResData, error)  {
	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()

	results := []ArticleSchema{}
	var err error = nil
	var count = 0
	var ret = ResData{}

	err = c.Find(condition).Limit(pageSize).Skip((pageNo - 1) * pageSize).All(&results)
	count, err = c.Find(condition).Count()
	fmt.Println("results:", results)

	ret.Articles = results
	ret.Count = count
	if pageNo * pageSize >= count {
		ret.IsEnd = 1
	} else {
		ret.IsEnd = 0
	}
	return ret, err
}

func ArticleMultiFindByKV(condition bson.M) []ArticleSchema {

	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()

	results := []ArticleSchema{}

	utils.HandleError("find error: ", c.Find(condition).All(&results))


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

func ChangeArticle(aid int, title string, tags []string, isPublish bool, comentN int, content string) error {
	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()

	selector := bson.M{"aid": aid}

	fmt.Println("isPublish", isPublish)
	data := bson.M{"title": title, "tags": tags, "ispublish": isPublish, "comentn": comentN, "content": content}

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
