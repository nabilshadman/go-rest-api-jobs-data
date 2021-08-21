package models

/*********
* MODEL *
*********/
// Error represents errors throughout project encoded through JSON.
// Has a message field.
type Error struct {
	Message string `json:"message"`
}
