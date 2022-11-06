package printer

import (
	"github.com/hpardora/absence.go/pkg/absence"
	"github.com/sirupsen/logrus"
)

type Printer struct {
	client *absence.Client
	logger *logrus.Logger

	reasons []absence.Reason
}

func New(client *absence.Client, logger *logrus.Logger) *Printer {
	return &Printer{
		client: client,
		logger: logger,
	}
}

func (p *Printer) Process() {
	// Retrieve User Data
	me, err := p.client.Me()
	if err != nil {
		p.logger.WithError(err).Errorln("unable to retrieve user information")
		return
	}
	p.printUser(me)

	// Retrieve User Company Data
	comp, err := p.client.MyCompany(me.Company)
	if err != nil {
		p.logger.WithError(err).Errorln("unable to retrieve company information")
		return
	}
	p.printCompany(comp)

	holidays, err := p.client.GetMyHolydays(comp.CurrentCompanyYear, me.HolidayIds)
	if err != nil {
		p.logger.WithError(err).Errorln("unable to retrieve user holidays")
	}
	p.printHolidays(holidays)

	absences, err := p.client.GetMyAbsences(me.ID, comp.CurrentCompanyYear)
	if err != nil {
		p.logger.WithError(err).Errorln("unable to retrieve user absences")
	}
	p.printAbsences(absences)

}

//func (c *Client) HaveToWork() (*Holiday, bool) {
//	hasToWork := false
//	today := time.Now().Weekday()
//	for _, workDay := range c.config.WorkingDays {
//		if today == workDay {
//			hasToWork = true
//		}
//	}
//
//	absences, err := c.GetMyAbsences()
//	if err != nil {
//		panic("unable to  retrieve absences")
//	}
//
//	onAbsence := c.todayHaveAbsence(absences)
//	if onAbsence {
//		return nil, false
//	}
//
//	holiday, isHoliDay := c.todayIsHoliday()
//	return holiday, !isHoliDay && hasToWork
//}
//
//func (c *Client) todayHaveAbsence(absences []Absence) bool {
//	now := time.Now()
//	for _, absence := range absences {
//
//		if now.After(absence.Start) && now.Before(absence.End) {
//			return true
//		}
//	}
//	return false
//}
//
//func (p *Printer) todayIsHoliday() (*absence.Holiday, bool) {
//	for _, holiday := range c.hollydays {
//		if len(holiday.Dates) == 1 {
//			if isToday(holiday.Dates[0]) {
//				return &holiday, true
//			}
//		} else {
//			for _, date := range holiday.Dates {
//				if date.Year() != c.company.CurrentCompanyYear {
//					continue
//				}
//				if isToday(date) {
//					return &holiday, true
//				}
//			}
//		}
//	}
//	return nil, false
//}
//
//func isToday(date time.Time) bool {
//	now := time.Now()
//	return now.Day() == date.Day() && now.Month() == date.Month()
//}
