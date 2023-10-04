package common

type User struct {
	UserId     uint   `gorm:"index" json:user_id gorm:"primary_key; auto_increment"`
	Account    string `gorm:"index" json:"account"`
	Password   string `gorm:"index" json:"password"`
	Username   string `gorm:"index" json:"username"`
	Signature  string `gorm:"index" json:"signature"`
	ProfileUrl string `gorm:"index" json:"profile_url"`
}

func (User) TableName() string {
	return "users"
}

type Video struct {
	VideoId  uint   `gorm:"index" json:"video_id" gorm:"primary_key; atuo_increment"`
	AuthorId uint   `gorm:"index" json:author_id`
	Title    string `gorm:"index", json:"title"`
	PlayUrl  string `gorm:"index" json:"play_url"`
	CoverUrl string `gorm:"index" json:"cover_url"`
}

func (Video) TableName() string {
	return "videos"
}

type UserVideoRelation struct {
	RelationId uint   `gorm:index json:"relation_id" gorm:"primary_key; auto_increment"`
	UserId     uint   `gorm:index json:"user_id"`
	VideoId    uint   `gorm:index json:"video_id"`
	Liked      bool   `gorm:index json:"liked"`
	Comment    string `gorm:index json:"comment"`
}

func (UserVideoRelation) TableName() string {
	return "user_video_relations"
}
