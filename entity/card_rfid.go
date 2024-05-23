package entity

type Card struct {
	ID   int    `json:"id"`
	UID  string `json:"uid" validate:"required"`
	Type string `json:"type" validate:"required"`
}
