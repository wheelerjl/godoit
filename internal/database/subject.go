package database

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
)

type Subject struct {
	SubjectID string `json:"subject_id"`
	Name      string `json:"name"`
	ImageURL  string `json:"image_url"`
}

func (c Client) AddSubject(ctx context.Context, subject Subject) (Subject, error) {
	if subject.SubjectID == "" {
		subject.SubjectID = uuid.NewString()
	}

	if _, err := c.DB.Exec(ctx, "INSERT INTO subjects (subject_id, name, image_url) VALUES ($1, $2, $3)",
		subject.SubjectID,
		subject.Name,
		subject.ImageURL); err != nil {
		return Subject{}, err
	}

	return subject, nil
}

func (c Client) GetSubjects(ctx context.Context) (subjects []Subject, err error) {
	subjects = []Subject{}
	return subjects, pgxscan.Select(ctx, c.DB, &subjects, "SELECT * FROM subjects")
}

func (c Client) GetSubject(ctx context.Context, subjectID string) (subject Subject, err error) {
	return subject, pgxscan.Get(ctx, c.DB, &subject, "SELECT * FROM subjects WHERE subject_id=$1", subjectID)
}

func (c Client) RemoveSubject(ctx context.Context, subjectID string) error {
	_, err := c.DB.Exec(ctx, "DELETE FROM subjects WHERE subject_id =$1", subjectID)
	return err
}
