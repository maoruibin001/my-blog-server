package db

func GetTags() ([]string, error)  {
	c, session := GetCollect("my-blog-2", "article")
	defer session.Close()

	results := []string{}

	err := c.Find(nil).Distinct("tags", &results)

	return results, err

}
