package models

/*********
* MODEL *
*********/
// Job is a simple model to store job data and encoded/decoded through JSON.
// Has an id, title, company, location, and type.
type Job struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Type     string `json:"type"`
}
