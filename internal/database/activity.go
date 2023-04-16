package database

import (
	"context"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
)

type Activity struct {
	ActivityID     string    `json:"activity_id"`
	SubjectID      string    `json:"subject_id"`
	AssignedUserID string    `json:"assigned_user_id"`
	Name           string    `json:"name"`
	Location       string    `json:"location"`
	Description    string    `json:"description"`
	StartTime      time.Time `json:"start_time"`
}

func (c Client) AddActivity(ctx context.Context, activity Activity) (Activity, error) {
	if activity.ActivityID == "" {
		activity.ActivityID = uuid.NewString()
	}

	if _, err := c.DB.Exec(ctx, "INSERT INTO activities (activity_id, subject_id, assigned_user_id, name, location, description, start_time) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		activity.ActivityID,
		activity.SubjectID,
		activity.AssignedUserID,
		activity.Name,
		activity.Location,
		activity.Description,
		activity.StartTime); err != nil {
		return Activity{}, err
	}

	return activity, nil
}

func (c Client) GetActivities(ctx context.Context, assignedUserID string) (activities []Activity, err error) {
	activities = []Activity{}
	if assignedUserID == "" {
		return activities, pgxscan.Select(ctx, c.DB, &activities, "SELECT * FROM activities")
	}
	return activities, pgxscan.Select(ctx, c.DB, &activities, "SELECT * FROM activities WHERE assigned_user_id=$1", assignedUserID)
}

func (c Client) GetActivity(ctx context.Context, activityID string) (activity Activity, err error) {
	return activity, pgxscan.Get(ctx, c.DB, &activity, "SELECT * FROM activities WHERE activity_id=$1", activityID)
}

func (c Client) RemoveActivity(ctx context.Context, activityID string) error {
	_, err := c.DB.Exec(ctx, "DELETE FROM activities WHERE activity_id =$1", activityID)
	return err
}
