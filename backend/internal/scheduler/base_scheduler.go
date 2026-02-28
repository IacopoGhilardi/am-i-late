package scheduler

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/iacopoGhilardi/amILate/internal/utils/logger"
)

type JobRegistrar interface {
	RegisterJobs(s gocron.Scheduler)
}

type Scheduler struct {
	Instance gocron.Scheduler
	Jobs     []JobRegistrar
}

func NewScheduler(registrars ...JobRegistrar) (*Scheduler, error) {
	instance, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	return &Scheduler{
		Instance: instance,
		Jobs:     registrars,
	}, nil
}

func (s *Scheduler) Start() {
	for _, job := range s.Jobs {
		job.RegisterJobs(s.Instance)
	}

	s.Instance.Start()
	logger.Info("Scheduler started")
}
