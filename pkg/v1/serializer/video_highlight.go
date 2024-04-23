package serializer

import v "github.com/go-ozzo/ozzo-validation/v4"

type VideoHighlightRequest struct {
	ID string `json:"id"`
}

type VideoHighlightContent struct {
	HighlightName   string `json:"hightlight_name"`
	HighlightDetail string `json:"highlight_detail"`
}

type VideoHightlightResponse struct {
	Content []VideoHighlightContent `json:"content"`
	QueueID string                  `json:"queue_id"`
	UserID  string                  `json:"user_id"`
	VideoID string                  `json:"video_id,omitempty"`
}

func (b VideoHighlightRequest) Validate() error {
	if err := v.ValidateStruct(&b, v.Field(&b.ID, v.Required)); err != nil {
		return err
	}

	return nil
}
