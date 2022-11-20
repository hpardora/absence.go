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
