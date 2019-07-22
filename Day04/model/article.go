package model

type Article struct {
	Model
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
	CoverImageUrl string `json:"cover_image_url"`
}

func ExistArticleByID(id int) bool {
	var article Article
	DB.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	DB.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	// Preload就是一个预加载器，它会执行两条SQL
	// 分别是SELECT * FROM blog_articles
	// SELECT * FROM blog_tag WHERE id IN (1,2,3,4);
	// 在查询出结构后，gorm内部处理对应的映射逻辑，将其填充到Article的Tag中，会特别方便，并且避免了循环查询
	DB.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

// gorm会通过类名+ID的方式去找到这两个类之间的关联关系
func GetArticle(id int) (article Article) {
	DB.Where("id = ?", id).First(&article)

	// 通过Related进行关联查询
	DB.Model(&article).Related(&article.Tag)

	return
}

func EditArticle(id int, data interface{}) bool {
	DB.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	DB.Create(&Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	})

	return true
}

func DeleteArticle(id int) bool {
	DB.Where("id = ?", id).Delete(Article{})

	return true
}

// 硬删除代码  要使用 Unscoped()，这是 GORM 的约定
func CleanAllArticle() bool {
	DB.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})
	return true
}

//func (article *Article) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("CreatedOn", time.Now().Unix())
//
//	return nil
//}
//
//func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("ModifiedOn", time.Now().Unix())
//
//	return nil
//}
