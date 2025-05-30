package models

import "time"

type Course struct {
	CourseID   string     `json:"course_id"`
	CourseName string     `json:"course_name"`
	LecturerID string       `json:"lecturer_id"`
	Semester   int        `json:"semester"`
	Credit     int        `json:"credit"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

var CourseQuery = `
	CREATE TABLE IF NOT EXISTS courses (
		course_id VARCHAR(6) PRIMARY KEY,
		course_name VARCHAR(40) NOT NULL,
		lecturer_id VARCHAR(3) NOT NULL,
		semester INT NOT NULL,
		credit INT NOT NULL,
		deleted_at TIMESTAMP NULL DEFAULT NULL,
		FOREIGN KEY (lecturer_id) REFERENCES lecturers(lecturer_id)
	);`
