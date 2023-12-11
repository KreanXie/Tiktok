package common

import "time"

type User struct {
	UserId     uint   `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username   string `gorm:"index" json:"username"`
	Password   string `gorm:"index" json:"password"`
	Signature  string `gorm:"index" json:"signature"`
	ProfileUrl string `gorm:"index" json:"profile_url"`
	Token      string `gorm:"index" json:"token"`
}

func (User) TableName() string {
	return "users"
}

type Video struct {
	VideoId  uint   `gorm:"primaryKey;autoIncrement"`
	AuthorId uint   `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title    string `gorm:"size:255;index" json:"title"`
	PlayUrl  string `gorm:"size:255;index" json:"play_url"`
	CoverUrl string `gorm:"size:255;index" json:"cover_url"`
	Author   User   `gorm:"foreignKey:AuthorId"`
}

func (Video) TableName() string {
	return "videos"
}

type Like struct {
	LikeId  uint `gorm:"primaryKey;autoIncrement"`
	VideoId uint `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserId  uint `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Like) TableName() string {
	return "likes"
}

type View struct {
	ViewId  uint `gorm:"primaryKey;autoIncrement"`
	UserId  uint `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	VideoId uint `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (View) TableName() string {
	return "views"
}

type Comment struct {
	CommentId uint      `gorm:"primaryKey;autoIncrement"`
	VideoId   uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserId    uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Text      string    `gorm:"index"`
	Timestamp time.Time `gorm:"type:datetime"`
}

func (Comment) TableName() string {
	return "comments"
}

type Report struct {
	ReportId uint `gorm:"primaryKey;autoIncrement"`
	UserId   uint `gorm:"index"`
	VideoId  uint `gorm:"index"`
	Text     uint `gorm:"index"`
}
