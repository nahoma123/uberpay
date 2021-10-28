package model

import (
	"mime/multipart"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Image struct {
	ID              uuid.UUID             `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name            string                `json:"name,omitempty" validate:"required"`
	ImageFile       *multipart.FileHeader `json:"-" form:"image" gorm:"-"`
	AlternativeText string                `json:"alternativeText,omitempty" form:"alternativeText"`
	Caption         string                `json:"caption,omitempty" form:"caption"`
	Width           int                   `json:"width,omitempty"`
	Height          int                   `json:"height,omitempty"`
	Hash            string                `json:"hash,omitempty"`
	Ext             string                `json:"ext,omitempty"`
	Mime            string                `json:"mime,omitempty"`
	Size            int64                 `json:"size,omitempty"`
	Url             string                `json:"url,omitempty"`
	PreviewUrl      string                `json:"previewUrl,omitempty" form:"previewUrl"`
	CreatedAt       time.Time             `json:"created_at,omitempty"`
	UpdatedAt       time.Time             `json:"updated_at,omitempty"`
}
type CompanyImage struct {
	Image   *Image `json:"images,omitempty"`
	Formats struct {
		Thumbnail *ImageFormat `json:"thumbnail,omitempty"`
		Small     *ImageFormat `json:"small,omitempty"`
	} `json:"formats,omitempty"`
}
type Form struct {
}
