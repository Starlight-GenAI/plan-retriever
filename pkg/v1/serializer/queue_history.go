package serializer

import (
	"time"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type ListQueueHistoryRequest struct {
	UserID string `json:"user_id"`
}

type QueueHistory struct {
	ID            string    `json:"queue_id"`
	VideoUrl      string    `json:"video_url"`
	VideoID       string    `json:"video_id"`
	Status        string    `json:"status"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Thumbnails    string    `json:"thumbnails"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	ChannelName   string    `json:"channel_name"`
	IsUseSubTitle bool      `json:"is_use_subtitle"`
}
type ListQueueHistoryResponse struct {
	Data []QueueHistory `json:"data"`
}

func (b ListQueueHistoryRequest) Validate() error {

	if err := v.ValidateStruct(&b, v.Field(&b.UserID, v.Required)); err != nil {
		return err
	}
	return nil
}
