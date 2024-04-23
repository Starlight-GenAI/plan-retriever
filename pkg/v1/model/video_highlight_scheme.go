package model

type VideoHighlightContent struct {
	HighlightName   string `firestore:"highlight_name"`
	HighlightDetail string `firestore:"highlight_detail"`
}

type VideoHighlightFirestore struct {
	Content []VideoHighlightContent `firestore:"content"`
	QueueID string                  `firestore:"queue_id"`
	UserID  string                  `firestore:"user_id"`
	VideoID string                  `firestore:"video_id"`
}

type VideoHightlight struct {
	Content []VideoHighlightContent
	QueueID string
	UserID  string
	VideoID string
}
