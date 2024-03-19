package identityModel

type Identity struct {
	ID int `gorm:"primarykey" json:"ID"`
	Name      string `json:"Name" binding:"required"`
	Age       int    `json:"Age" binding:"required,numeric"`
	Job       string `json:"Job" binding:"required"`
	IsMarried bool   `json:"IsMarried"`
}
