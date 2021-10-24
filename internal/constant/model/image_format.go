package model

import uuid "github.com/satori/go.uuid"

type ImageFormat struct {
	Name       string    `json:"name,omitempty"`
	Hash       string    `json:"hash,omitempty"`
	Ext        string    `json:"ext,omitempty"`
	Mime       string    `json:"mime,omitempty"`
	Width      int       `json:"width,omitempty"`
	Height     int       `json:"height,omitempty"`
	Size       int64     `json:"size,omitempty"`
	Path       string    `json:"path,omitempty"`
	Url        string    `json:"url,omitempty"`
	FormatType string    `json:"format_type,omitempty"`
	ImageID    uuid.UUID `json:"image_id,omitempty" `
}
