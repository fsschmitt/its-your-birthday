package models

type GiftContributor struct {
	Name string `json:"name,omitempty"`

	// User has paid or not
	HasPaid bool `json:"hasPaid,omitempty"`

	// All the ids of the suggested gifts voted
	Upvotes []string `json:"upvotes,omitempty"`
}
