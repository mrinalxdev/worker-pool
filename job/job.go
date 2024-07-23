package job

type Job struct {
	ID          int
	Description string
}

func New(id int, description string) Job {
	return Job{
		ID:          id,
		Description: description,
	}
}
