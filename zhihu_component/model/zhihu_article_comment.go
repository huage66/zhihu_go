package model

import "time"

// ZhihuArticleComment 文章评论表
type ZhihuArticleComment struct {
	ID        int64     `gorm:"primaryKey;column:id"` // 主键ID
	Title     string    `gorm:"column:title"`         // 话题标题
	Content   string    `gorm:"column:content"`       // 回答的内容
	UserID    int64     `gorm:"column:user_id"`       // 用户id
	ArticalID int64     `gorm:"column:artical_id"`    // 话题id
	Delete    *int8     `gorm:"column:delete"`        // 是否删除, 默认0-未删除 1-删除
	CreatedAt time.Time `gorm:"column:created_at"`    // 记录创建时间
	UpdatedAt time.Time `gorm:"column:updated_at"`    // 更新时间
}

func (z ZhihuArticleComment) TableName() string {
	return "zhihu_article_comment"
}

