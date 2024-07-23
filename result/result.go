package result

type Result struct {
	JobID  int
	Output string
}

func New(jobID int, output string) Result {
	return Result{
		JobID:  jobID,
		Output: output,
	}
}
