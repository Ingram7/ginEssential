package models

type Category struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Name      string     `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreatedAt TimeNormal `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt TimeNormal `json:"updated_at" gorm:"type:timestamp"`
}
