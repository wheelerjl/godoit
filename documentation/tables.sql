CREATE SCHEMA godoit;
SET SCHEMA 'godoit';

CREATE TABLE subjects (
  subject_id UUID PRIMARY KEY,
  name VARCHAR NOT NULL,
  image_url TEXT
);

CREATE TABLE activities (
  activity_id UUID PRIMARY KEY,
  subject_id UUID NOT NULL REFERENCES subject(subject_id),
  assigned_user_id TEXT,
  name TEXT NOT NULL,
  location TEXT,
  long_desc TEXT,
  start_time TIMESTAMP WITH TIME ZONE NOT NULL
);
