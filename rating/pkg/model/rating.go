package model

// RecordID unique identifier for record
type RecordID string

// RecordType context for record identifier
type RecordType string

// Existing record types
const (
	RecordTypeMovie = RecordType("movie")
)

// UserID unique user
type UserID string

// RatingValue score for rating
type RatingValue int

// Rating defines user rating full context
type Rating struct {
	RecordID   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserID     UserID      `json:"userId"`
	Value      RatingValue `json:"value"`
}
