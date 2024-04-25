package serializer

import v "github.com/go-ozzo/ozzo-validation/v4"

type TripSummaryRequest struct {
	ID string `json:"id"`
}

type LocationWithSummary struct {
	LocationName string   `json:"location_name"`
	Summary      string   `json:"summary"`
	PlaceID      string   `json:"place_id"`
	Lat          float64  `json:"lat"`
	Lng          float64  `json:"lng"`
	Rating       float64  `json:"rating"`
	Category     string   `json:"category"`
	Photos       []string `json:"photos"`
	// HasRecommendedRestaurant bool             `json:"has_recommended_restaurant"`
	// RecommendedRestaurant    RestaurantDetail `json:"recommended_restaurant,omitempty"`
}

type Photo struct {
	Reference string `json:"reference"`
	MaxWidth  int    `json:"max_width"`
	MaxHeight int    `json:"max_height"`
}

type RestaurantDetail struct {
	Name    string   `json:"name,omitempty"`
	Summary string   `json:"summary,omitempty"`
	Rating  float64  `json:"rating,omitempty"`
	PlaceID string   `json:"place_id,omitempty"`
	Lat     float64  `json:"lat,omitempty"`
	Lng     float64  `json:"lng,omitempty"`
	Photos  []string `json:"photos"`
}
type TripSummaryContent struct {
	Day                 string                `json:"day"`
	LocationWithSummary []LocationWithSummary `json:"location_with_summary"`
	CountDining         int                   `json:"count_dining"`
	// BreakfastRestaurant RestaurantDetail      `json:"breakfast_restaurant"`
	// LunchRestaurant     RestaurantDetail      `json:"lunch_restaurant"`
	// DinnerRestaurant    RestaurantDetail      `json:"dinner_restaurant"`
}

type TripSummaryResponse struct {
	Content []TripSummaryContent `json:"content"`
	UserID  string               `json:"user_id"`
	VideoID string               `json:"video_id,omitempty"`
}

func (b TripSummaryRequest) Validate() error {

	if err := v.ValidateStruct(&b, v.Field(&b.ID, v.Required)); err != nil {
		return err
	}

	return nil
}

type ListTripSummaryRequest struct {
	UserID string `json:"user_id"`
}

func (b ListTripSummaryRequest) Validate() error {

	if err := v.ValidateStruct(&b, v.Field(&b.UserID, v.Required)); err != nil {
		return err
	}

	return nil
}
