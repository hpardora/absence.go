package scheduler

import (
	"github.com/hpardora/absence.go/pkg/absence"
	"github.com/sirupsen/logrus"
)

type Scheduler struct {
	client *absence.Client
	logger *logrus.Logger
	conf   *Config

	user     *absence.User
	company  *absence.Company
	holidays []absence.Holiday
	absences []absence.Absence
	reasons  []absence.Reason
}

func New(cfg *Config, client *absence.Client, logger *logrus.Logger) *Scheduler {
	return &Scheduler{
		conf:   cfg,
		client: client,
		logger: logger,
	}
}

func (s *Scheduler) Process() {
	s.populateAllInformation()
	s.Calculate()
}
