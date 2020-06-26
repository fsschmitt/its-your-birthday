package models

// Describes the instance of the gift where people will all suggest stuff and pay later
type Gift struct {
	Id int64 `json:"id,omitempty"`

	Contributors []GiftContributor `json:"contributors,omitempty"`
}
