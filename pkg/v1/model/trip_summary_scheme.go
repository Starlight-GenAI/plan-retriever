package model

const (
	DINING = "dining"
)

type LocationWithSummary struct {
	LocationName             string           `firestore:"location_name"`
	Summary                  string           `firestore:"summary"`
	PlaceID                  string           `firestore:"place_id"`
	Lat                      float64          `firestore:"lat"`
	Lng                      float64          `firestore:"lng"`
	Rating                   float64          `firestore:"rating"`
	Category                 string           `firestore:"category"`
	Photos                   []Photo          `firestore:"photos"`
	HasRecommendedRestaurant bool             `firestore:"has_recommended_restaurant"`
	RecommendedRestaurant    RestaurantDetail `firestore:"recommended_restaurant"`
}

type Photo struct {
	Reference string `firestore:"reference"`
	MaxWidth  int    `firestore:"max_width"`
	MaxHeight int    `firestore:"max_height"`
}
type RestaurantDetail struct {
	Name    string  `firestore:"name"`
	Summary string  `firestore:"summary"`
	Rating  float64 `firestore:"rating"`
	PlaceID string  `firestore:"place_id"`
	Lat     float64 `firestore:"lat"`
	Lng     float64 `firestore:"lng"`
	Photos  []Photo `firestore:"photos"`
}
type TripSummaryContent struct {
	Day                 string                `firestore:"day"`
	LocationWithSummary []LocationWithSummary `firestore:"location_with_summary"`
}

type TripSummaryFirestore struct {
	QueueID string               `firestore:"queue_id"`
	Content []TripSummaryContent `firestore:"content"`
	UserID  string               `firestore:"user_id"`
	VideoID string               `firestore:"video_id"`
}

type TripSummary struct {
	Content []TripSummaryContent
	UserID  string
	VideoID string
}
