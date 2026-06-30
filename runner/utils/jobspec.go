package utils

type JobSpec struct {
	JobId        string `json:"job_id"`
	Language     string `json:"language"`
	Version      string `json:"version"`
	SubmissionID string `json:"submission_id"`
	FilePath     string `json:"filepath"`
	// S3Key          string
	Testset        string `json:"testset"`
	TestsetVersion string `json:"testset_version"`
}
