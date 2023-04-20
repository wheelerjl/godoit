-- +goose Up
CREATE SCHEMA godoit;
CREATE TABLE godoit.subjects (
  subject_id UUID PRIMARY KEY,
  name VARCHAR NOT NULL,
  image_url TEXT
);

CREATE TABLE godoit.activities (
  activity_id UUID PRIMARY KEY,
  subject_id UUID NOT NULL REFERENCES godoit.subjects(subject_id),
  assigned_user_id TEXT,
  name TEXT NOT NULL,
  location TEXT,
  description TEXT,
  start_time TIMESTAMP WITH TIME ZONE NOT NULL
);

-- +goose Down
DROP TABLE godoit.activities;
DROP TABLE godoit.subjects;
DROP SCHEMA godoit;
