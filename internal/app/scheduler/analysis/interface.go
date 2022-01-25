package analysis

import "github.com/gocraft/work"

// Scheduler ...
type Scheduler interface {
	// StartJob ...
	StartJob(*work.Job) error

	// Run ...
	Run()
}
