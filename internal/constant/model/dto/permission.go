package dto

type Permission struct {
	ID        uint   `json:"id"`
	UserId    string `gorm:"column:v0" json:"user_id"`
	CompanyId string `gorm:"column:v1" json:"company,omitempty"`
	Name      string `gorm:"column:v2" json:"name"`
	Action    string `gorm:"column:v2" json:"description"`
}
