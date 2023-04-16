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

	if _, err := c.DB.Exec(ctx, "INSERT INTO subjects (subject_id, name, image_url) VALUES ($1, $2, $3)", subject.SubjectID, subject.Name, subject.ImageURL); err != nil {
		return Subject{}, err
	}

	return subject, nil
}

func (c Client) GetSubjects(ctx context.Context) ([]Subject, error) {
	var subjects []Subject
	if err := pgxscan.Select(ctx, c.DB, &subjects, "SELECT * FROM subjects"); err != nil {
		return []Subject{}, err
	}

	return subjects, nil
}

func (c Client) GetSubject(ctx context.Context, subjectID string) (Subject, error) {
	var subject Subject
	if err := pgxscan.Get(ctx, c.DB, &subject, "SELECT subject_id, name, image_url FROM subjects WHERE subject_id=$1", subjectID); err != nil {
		return Subject{}, err
	}

	return subject, nil
}

func (c Client) RemoveSubject(ctx context.Context, subjectID string) error {
	if _, err := c.DB.Exec(ctx, "DELETE FROM subjects WHERE subject_id =$1", subjectID); err != nil {
		return err
	}

	return nil
}
