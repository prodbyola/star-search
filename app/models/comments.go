package models

import "gorm.io/gorm"

type Comment struct {
	*gorm.Model
	Movie   int64
	Author  string
	Content string
}

func (c *Comment) Save(DB *gorm.DB) error {
	err := DB.AutoMigrate(&c)
	err = DB.Create(&c).Error
	return err
}

func (c *Comment) GetData() CommentData {
	return CommentData{
		Movie:     c.Movie,
		Content:   c.Content,
		CreatedAt: c.CreatedAt.UTC().String(),
		Author:    c.Author,
		ID:        c.ID,
	}
}

type CommentData struct {
	ID        uint   `json:"id"`
	Movie     int64  `json:"movie_id"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func GetComments(DB *gorm.DB, movie_id int64, rtype string) (interface{}, error) {
	comments := []Comment{}
	results := DB.Model(&Comment{}).Where("movie = ?", movie_id).Order("created_at").Find(&comments)
	if err := results.Error; err != nil {
		return nil, err
	}

	if rtype == "count" {
		return results.RowsAffected, nil
	}

	data := []CommentData{}
	for i := 0; i < len(comments); i++ {
		comment := comments[i]
		data = append(data, comment.GetData())
	}

	return data, nil
}
