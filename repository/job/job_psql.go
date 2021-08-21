package jobRepository

import (
	"database/sql"
	"jobs-list/models"
)

type JobRepository struct{}

func (b JobRepository) GetJobs(db *sql.DB, job models.Job, jobs []models.Job) ([]models.Job, error) {
	rows, err := db.Query("select * from jobs")

	if err != nil {
		return []models.Job{}, err
	}

	for rows.Next() {
		err = rows.Scan(&job.ID, &job.Title, &job.Company, &job.Location, &job.Type)
		jobs = append(jobs, job)
	}

	if err != nil {
		return []models.Job{}, err
	}

	return jobs, nil
}

func (b JobRepository) GetJob(db *sql.DB, job models.Job, id int) (models.Job, error) {
	rows := db.QueryRow("select * from jobs where id=$1", id)
	err := rows.Scan(&job.ID, &job.Title, &job.Company, &job.Location, &job.Type)

	return job, err
}

func (b JobRepository) AddJob(db *sql.DB, job models.Job) (int, error) {
	err := db.QueryRow(
		"insert into jobs (title, company, location, type) values($1, $2, $3, $4) RETURNING id;",
		job.Title, job.Company, job.Location, job.Type).Scan(&job.ID)

	if err != nil {
		return 0, err
	}

	return job.ID, nil
}

func (b JobRepository) UpdateJob(db *sql.DB, job models.Job) (int64, error) {
	result, err := db.Exec(
		"update jobs set title=$1, company=$2, location=$3, type=$4 where id=$5 RETURNING id",
		&job.Title, &job.Company, &job.Location, &job.Type, &job.ID)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func (b JobRepository) RemoveJob(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("delete from jobs where id = $1", id)

	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}
