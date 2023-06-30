package entities

type Employee struct {
	ID         int    `json:"-"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"projectID"`
}
