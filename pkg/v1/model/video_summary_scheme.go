package model

type VideoSummaryContent struct {
	LocationName string  `firestore:"location_name"`
	StartTime    float64 `firestore:"start_time"`
	EndTime      float64 `firestore:"end_time"`
	Summary      string  `firestore:"summary"`
	PlaceID      string  `firestore:"place_id"`
	Lat          float64 `firestore:"lat"`
	Lng          float64 `firestore:"lng"`
	Category     string  `firestore:"category"`
	Photos       []Photo `firestore:"photos"`
}
type VideoSummaryFireStore struct {
	QueueID         string                `firestore:"queue_id"`
	UserID          string                `firestore:"user_id"`
	Content         []VideoSummaryContent `firestore:"content"`
	CanGenerateTrip bool                  `firestore:"can_generate_trip"`
	VideoID         string                `firestore:"video_id"`
}

type VideoSummary struct {
	Content         []VideoSummaryContent
	CanGenerateTrip bool
	UserID          string
	VIdeoID         string
}
