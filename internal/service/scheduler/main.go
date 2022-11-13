package scheduler

import (
	"github.com/go-co-op/gocron"
	"github.com/hpardora/absence.go/pkg/absence"
	"github.com/sirupsen/logrus"
	"time"
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
	cron := gocron.NewScheduler(time.Local)
	_, err := cron.Every(1).Day().Tag("initializer").At(s.conf.CronExecutionTime).Do(func() {
		s.Calculate()
	})
	if err != nil {
		s.logger.WithError(err).Errorf("unable to append job")
	}
	cron.StartBlocking()
}
