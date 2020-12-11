package models

type Game struct {
	ID          int    `gorm:"primarykey" json:"id"`
	Cid         int    `gorm:"column:pid" json:"pid"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	MinPrice    int
	Unit        string
	Desc        string
	Sort        int16
	Status      int8
	IsRecommend int8
	CreateAt    int64 `gorm:"autoCreateTime"`
	UpdateAt    int64 `gorm:"autoCreateTime"`
}

func (game *Game) TableName() string {
	return "idol_game"
}
