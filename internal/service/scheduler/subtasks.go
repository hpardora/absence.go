package scheduler

import (
	"sort"
	"time"
)

type Holiday struct {
	ID   string
	Date time.Time
	Name string
}

func (s *Scheduler) populateAllInformation() {
	s.populateUser()
	s.populateCompany()
	s.populateReasons()
	s.populateHolidays()
	s.populateAbsences()
}

func (s *Scheduler) populateUser() {
	s.logger.Infof("retrieving user information")
	me, err := s.client.Me()
	if err != nil {
		s.logger.WithError(err).Errorf("unable to retrieve my information")
	}
	s.user = me
}

func (s *Scheduler) populateCompany() {
	s.logger.Infof("retrieving company information")
	comp, err := s.client.MyCompany(s.user.Company)
	if err != nil {
		s.logger.WithError(err).Errorf("unable to retrieve company information")
	}
	s.company = comp

}

func (s *Scheduler) populateReasons() {
	s.logger.Infof("retrieving reasons information")
	reasons, err := s.client.GetReasons()
	if err != nil {
		s.logger.WithError(err).Errorf("unable to retrieve reason information")
	}
	s.reasons = reasons
}

func (s *Scheduler) populateHolidays() {
	s.logger.Infof("retrieving reasons information")
	holidays, err := s.client.GetMyHolydays(s.company.CurrentCompanyYear, s.user.HolidayIds)
	if err != nil {
		s.logger.WithError(err).Errorf("unable to retrieve reason information")
	}
	s.holidays = holidays
}

func (s *Scheduler) populateAbsences() {
	s.logger.Infof("retrieving absences information")
	absences, err := s.client.GetMyAbsences(s.user.ID, s.company.CurrentCompanyYear)
	if err != nil {
		s.logger.WithError(err).Errorf("unable to retrieve reason information")
	}
	s.absences = absences
}

func (s *Scheduler) clearHolidays() []*Holiday {
	now := time.Now()
	var dates []*Holiday
	for _, h := range s.holidays {
		date := time.Time{}
		if len(h.Dates) == 1 {
			date = h.Dates[0]
		} else {
			for _, d := range h.Dates {
				if d.Year() == now.Year() {
					date = d
					break
				}
			}
		}
		dates = append(dates, &Holiday{
			ID:   h.Id,
			Date: date,
			Name: h.Name,
		})
	}
	sort.SliceStable(dates, func(i, j int) bool {
		return dates[i].Date.Before(dates[j].Date)
	})
	return dates
}
