package serializer

import v "github.com/go-ozzo/ozzo-validation/v4"

type VideoSummaryRequest struct {
	ID string `json:"id"`
}

type VideoSummaryContent struct {
	LocationName string   `json:"location_name"`
	StartTime    float64  `json:"start_time"`
	EndTime      float64  `json:"end_time"`
	Summary      string   `json:"summary"`
	PlaceID      string   `json:"place_id"`
	Lat          float64  `json:"lat"`
	Lng          float64  `json:"lng"`
	Category     string   `json:"category"`
	Photos       []string `json:"photos"`
}
type VideoSummaryResponse struct {
	Content         []VideoSummaryContent `json:"content"`
	CanGenerateTrip bool                  `json:"can_generate_trip"`
	UserID          string                `json:"user_id"`
	VideoID         string                `json:"video_id,omitempty"`
}

func (b VideoSummaryRequest) Validate() error {

	if err := v.ValidateStruct(&b, v.Field(&b.ID, v.Required)); err != nil {
		return err
	}

	return nil
}

type ListVideoSummaryRequest struct {
	UserID string `json:"user_id"`
}

func (b ListVideoSummaryRequest) Validate() error {

	if err := v.ValidateStruct(&b, v.Field(&b.UserID, v.Required)); err != nil {
		return err
	}

	return nil
}

type GetVideoSummaryByCategory struct {
	ID       string `json:"id"`
	Category string `json:"category"`
}

func (b GetVideoSummaryByCategory) Validate() error {

	if err := v.ValidateStruct(&b,
		v.Field(&b.ID, v.Required),
		v.Field(&b.Category, v.Required),
	); err != nil {
		return err
	}

	return nil
}
