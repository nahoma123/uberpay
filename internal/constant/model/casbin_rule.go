package model

type CasbinRule struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement;"`
	Ptype string `json:"ptype,omitempty" gorm:"size:512;"    validate:"required" `
	V0    string `json:"v_0,omitempty"   gorm:"size:512;"    validate:"required"`
	V1    string `json:"v_1,omitempty"   gorm:"size:512;"    validate:"required"`
	V2    string `json:"v_2,omitempty"   gorm:"size:512;"    validate:"required"`
	V3    string `json:"v_3,omitempty"   gorm:"size:512;"    validate:"required"`
	V4    string ` json:"v_4,omitempty"  gorm:"size:512;"    validate:"required"`
	V5    string `json:"v_5,omitempty"   gorm:"size:512;"    validate:"required"`
}
