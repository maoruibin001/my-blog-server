package db

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"album-server/src/utils"
	"strconv"
	"time"
)
type CommentSetSchema struct {
	Aid int `json:"aid"`
	Count int `json:"count"`
	Comments []CommentSchema `json:"comments"`
}
type CommentSchema struct {
	ImgName string `json:"imgName"`
	Name string `json:"name"`
	Address string `json:"address"`
	Content string `json:"content"`
	Aid int `json:"aid"`
	Cid int `json:"cid"`
	Date int64 `json:"date"`
	Like int `json:"like"`
	LikeEmails []string `json:"likeEmails"`
}



func generationCid() int {
	counter, session := GetCollect("my-blog-2", "ccounter")
	defer session.Close()

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"cid": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	doc := struct{ Cid int }{}
	if _, err := counter.Find(bson.M{}).Apply(change, &doc); err != nil {
		utils.HandleError("查找出错", err)
	}
	log.Println("doc:", doc)
	return doc.Cid
}

func CommentSingleFindByKV(key string, v interface{}) CommentSetSchema {

	c, session := GetCollect("my-blog-2", "comment")
	defer session.Close()

	results := []CommentSetSchema{}

	utils.HandleError("find error: ", c.Find(bson.M{key: v}).All(&results))

	result := CommentSetSchema{}
	if len(results) > 0 {
		result = results[0]
	}
	return result
}


func CommentMultiFindByKV( key string, v interface{}) []CommentSchema {

	c, session := GetCollect("my-blog-2", "comment")
	defer session.Close()

	results := []CommentSchema{}

	utils.HandleError("find error: ", c.Find(bson.M{key: v}).All(&results))


	return results
}

func AddComment(aid int, imgName, name, address, content string) (CommentSchema, error) {

	m := CommentSchema{}
	comments, err := QueryCommentSet(aid)


	fmt.Println("query err: ", err)



	m.Like = 0
	m.LikeEmails = []string{}
	m.Date = time.Now().UnixNano() / 1e6
	m.Aid = aid
	m.Cid = generationCid()
	m.Content = content
	m.ImgName = imgName
	m.Name = name
	m.Address = address



	c, session := GetCollect("my-blog-2", "comment")
	defer session.Close()

	selector := bson.M{"aid": aid}


	if len(comments.Comments) == 0 && comments.Aid == 0 {
		fmt.Println("err", err, "插入一条新数据")
		comments = CommentSetSchema{}
		comments.Comments = append(comments.Comments, m)
		comments.Aid = aid
		comments.Count = len(comments.Comments)

		fmt.Println("comments is: ", comments)
		err = c.Insert(&comments)

		fmt.Println("插入出错", err)

	} else {

		comments.Comments = append(comments.Comments, m)

		comments.Count = len(comments.Comments)

		fmt.Println("更新数据", m)
		err = c.Update(selector, bson.M{"$set": comments})

	}



	return m, err



}

func DeleteComment(aid, cid int) error  {
	comments, index, err := QueryCommentIndex(aid, cid)

	if err != nil {
		return nil
	}

	m := comments.Comments[index]
	comments.Comments = append(comments.Comments[:index], comments.Comments[index + 1:]...)

	comments.Count = len(comments.Comments)
	c, session := GetCollect("my-blog-2", "comment")
	defer session.Close()

	selector := bson.M{"aid": aid}
	err = c.Update(selector, bson.M{"$set": comments})

	fmt.Println("删除一条数据", m)
	return err
}

func ChangeComment(aid, cid int, imgName, name, address, content string, like int, likeEmails []string) (CommentSchema, error)  {

	m := CommentSchema{}
	comments, index,  err := QueryCommentIndex(aid, cid)


	if err != nil {
		return m, err
	}

	m.Like = like
	m.Date = time.Now().UnixNano() / 1e6
	m.Aid = aid
	m.Cid = cid
	m.Content = content
	m.ImgName = imgName
	m.Name = name
	m.Address = address
	m.LikeEmails = likeEmails

	comments.Comments[index] = m


	c, session := GetCollect("my-blog-2", "comment")
	defer session.Close()

	selector := bson.M{"aid": aid}

	err = c.Update(selector, bson.M{"$set": comments})

	fmt.Println("更新一条数据", m)

	return m, nil
}

func QueryCommentSet(aid int)  (CommentSetSchema, error) {

	fmt.Println("aid is: ", aid)
	c, session := GetCollect("my-blog-2", "comment")
	defer session.Close()


	result := CommentSetSchema{}
	results := []CommentSetSchema{}
	var err error = nil

	if aid == 0 {
		return result, err
	}

	err = c.Find(bson.M{"aid": aid}).All(&results)

	if len(results) == 0 {
		return result, err
	} else {
		return results[0], err
	}
}

func QueryCommentIndex(aid, cid int) (CommentSetSchema, int, error) {
	commentSet, err := QueryCommentSet(aid)

	var index = -1

	if err != nil {
		return commentSet, index, err
	}

	for i, v := range commentSet.Comments {

		if v.Cid == cid {
			index = i
		}
	}

	fmt.Println("index: ", index)

	if index == -1 {
		err = errors.New("no cid: " + strconv.Itoa(cid))
		return commentSet, index, err
	}


	return commentSet, index, err

}