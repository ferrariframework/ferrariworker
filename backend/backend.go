package backend

import "github.com/ottogiron/ferrariworker/processor"

//Backend defines a data store for jobs to be persisted
type Backend interface {
	Persist(workerId string, jobID string, jobResults []processor.JobResult) error
	JobResults(workerId string, jobID) ([]processor.JobResult, error)
}
