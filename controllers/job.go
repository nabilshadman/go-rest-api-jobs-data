package controllers

import (
	"database/sql"
	"encoding/json"
	"jobs-list/models"
	jobRepository "jobs-list/repository/job"
	"jobs-list/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct{}

var jobs []models.Job

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetJobs returns all jobs in collection as a JSON array.
func (c Controller) GetJobs(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var job models.Job
		var error models.Error

		jobs = []models.Job{}
		jobRepo := jobRepository.JobRepository{}
		jobs, err := jobRepo.GetJobs(db, job, jobs)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, jobs)
	}

}

// GetJob reads ID from URL (using mux) and queries database for job in question.
// Returns JSON data if job exists, otherwise returns 404.
func (c Controller) GetJob(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var job models.Job
		var error models.Error

		jobs = []models.Job{}
		jobRepo := jobRepository.JobRepository{}

		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		job, err := jobRepo.GetJob(db, job, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Not Found"
				utils.SendError(w, http.StatusNotFound, error) //404
				return
			} else {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error) //500
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, job)
	}

}

// AddJob reads in JSON data and creates a new job within database.
// Returns ID of the newly created job.
func (c Controller) AddJob(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var job models.Job
		var jobID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&job)

		if job.Title == "" || job.Company == "" || job.Location == "" || job.Type == "" {
			error.Message = "Enter missing fields"
			utils.SendError(w, http.StatusBadRequest, error) //400
			return
		}

		jobRepo := jobRepository.JobRepository{}
		jobID, err := jobRepo.AddJob(db, job)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, jobID)
	}

}

// UpdateJob queries database for ID of job to be updated.
// Replaces the writeable data within the job with JSON data,
// and rewrites job to the database.
// Returns rows updated from request.
func (c Controller) UpdateJob(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var job models.Job
		var error models.Error

		json.NewDecoder(r.Body).Decode(&job)

		if job.ID == 0 || job.Title == "" || job.Company == "" || job.Location == "" || job.Type == "" {
			error.Message = "All fields are required"
			utils.SendError(w, http.StatusBadRequest, error) //400
			return
		}

		jobRepo := jobRepository.JobRepository{}
		rowsUpdated, err := jobRepo.UpdateJob(db, job)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)
	}
}

// RemoveJob deletes job from database based off ID passed in through URL.
// Returns rows deleted if successful, 404 if no job exists for ID.
func (c Controller) RemoveJob(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		params := mux.Vars(r)
		jobRepo := jobRepository.JobRepository{}
		id, _ := strconv.Atoi(params["id"])

		rowsDeleted, err := jobRepo.RemoveJob(db, id)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error) //500
			return
		}

		if rowsDeleted == 0 {
			error.Message = "Not Found"
			utils.SendError(w, http.StatusNotFound, error) //404
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsDeleted)
	}
}
