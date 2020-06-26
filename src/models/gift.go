package models

// Describes the instance of the gift where people will all suggest stuff and pay later
type Gift struct {
	ID int64 `json:"id,omitempty"`

	Giftee string `json:"giftee,omitempty"`

	Contributors []GiftContributor `json:"contributors,omitempty"`
}
