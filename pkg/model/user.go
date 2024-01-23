package todo

type User struct {
	Id       int    `json:"-" gorm:"column:id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" gorm:"column:password_hash"`
}
