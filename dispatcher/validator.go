package dispatcher

/*
Step1:
- Job ID uniqueness (query s3)
- Language availability (config manager)
- Language version availability (config manager)
- Testset existence (query s3)
- Testset version existence (query s3)

Step2:
- Generate a unique `submission_id`.
- Put file content in {job_id}/{submission_id}.*
- Upload the source file to S3.
- Construct the Job Specification.
*/

func validateSubmission(submission SubmissionSpec) error {
	return nil
}
