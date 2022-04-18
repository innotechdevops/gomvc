package promotion

import "gorm.io/gorm"

type Promotion struct {
	gorm.Model
	Id   int    `json:"id"`
	Name string `json:"name"`
}
