package database

type Activity struct {
	ID             string
	SubjectID      string
	AssignedUserID string
	Name           string
	Location       string
	Description    string
	StartTime      string
}
