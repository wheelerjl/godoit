CREATE SCHEMA godoit;
SET SCHEMA 'godoit';

CREATE TABLE subject (
  subject_id UUID PRIMARY KEY,
  name VARCHAR NOT NULL,
  image_url TEXT
);

CREATE TABLE activity (
  activity_id UUID PRIMARY KEY,
  subject_id UUID NOT NULL REFERENCES subject(subject_id),
  assigned_user_id UUID,
  name TEXT NOT NULL,
  location TEXT,
  long_desc TEXT,
  start_time TIMESTAMP WITH TIME ZONE NOT NULL
);
