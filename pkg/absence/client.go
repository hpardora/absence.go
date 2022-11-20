package absence

import (
	"encoding/json"
	"fmt"
	"github.com/hiyosi/hawk"
	"github.com/hpardora/absence.go/pkg/randStrings"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"time"
)

const baseURL string = "https://app.absence.io/api/v2"

const (
	EndpointReasons        string = "/reasons"
	EndpointHolidays       string = "/holidays"
	EndpointAbsences       string = "/absences"
	EndpointTimeSpan       string = "/timespans"
	EndpointTimeSpanCreate string = "/timespans/create"
)

type Client struct {
	config *Config
	logger *logrus.Logger
}

func New(c *Config, logger *logrus.Logger) *Client {
	client := &Client{
		config: c,
		logger: logger,
	}
	return client
}

func (c *Client) GetHawkClient() *hawk.Client {
	nonce := randStrings.RandStringRunes(6)
	hawkClient := hawk.NewClient(
		&hawk.Credential{
			ID:  c.config.ID,
			Key: c.config.Key,
			Alg: hawk.SHA256,
		},
		&hawk.Option{
			TimeStamp: time.Now().Unix(),
			Nonce:     nonce,
			Ext:       "some-app-data",
		},
	)
	return hawkClient
}

func (c *Client) buildHeader(method string, path string) string {
	// build request header
	hawkClient := c.GetHawkClient()
	header, _ := hawkClient.Header(method, path)
	return header
}

func (c *Client) doRequest(url string, method string, payload *strings.Reader) []byte {
	header := c.buildHeader(method, url)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		c.logger.WithError(err).Error("unable to create new request")
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Authorization", header)
	res, err := client.Do(req)
	if err != nil {
		c.logger.WithError(err).Errorf("unable to do request %s", url)
		return nil
	}
	defer res.Body.Close()
	c.logger.Infof("method: %s\tresult_code: %d\turl: %s", method, res.StatusCode, url)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.logger.WithError(err).Error("unable to read body content")
		return nil
	}
	return body
}

func (c *Client) doPostRequest(endpoint string, payload *strings.Reader) []byte {
	url := fmt.Sprintf("%s%s", baseURL, endpoint)
	method := "POST"
	respBytes := c.doRequest(url, method, payload)
	return respBytes
}

func (c *Client) doPutRequest(endpoint string, payload *strings.Reader) []byte {
	url := fmt.Sprintf("%s%s", baseURL, endpoint)
	method := "PUT"
	respBytes := c.doRequest(url, method, payload)
	return respBytes
}

func (c *Client) doPostRequestCommon(endpoint string, filter string, sortBy string) []byte {
	url := fmt.Sprintf("%s%s", baseURL, endpoint)
	method := "POST"
	payload := strings.NewReader(`{
		"skip":0,
		"limit":5000,
		"filter":` + filter + `,
		"sortBy":` + sortBy +
		`}`)

	respBytes := c.doRequest(url, method, payload)
	return respBytes
}

func (c *Client) Me() (*User, error) {
	url := fmt.Sprintf("%s/users/%s", baseURL, c.config.ID)
	method := "GET"
	payload := strings.NewReader(``)
	respBytes := c.doRequest(url, method, payload)
	c.logger.Debug(string(respBytes))
	me := User{}
	if err := json.Unmarshal(respBytes, &me); err != nil {
		return nil, err
	}
	return &me, nil
}

func (c *Client) MyCompany(companyID string) (*Company, error) {
	url := fmt.Sprintf("%s/companies/%s", baseURL, companyID)
	method := "GET"
	payload := strings.NewReader(``)

	respBytes := c.doRequest(url, method, payload)
	c.logger.Debug(string(respBytes))
	comp := Company{}
	if err := json.Unmarshal(respBytes, &comp); err != nil {
		return nil, err
	}
	return &comp, nil
}

func (c *Client) GetReasons() ([]Reason, error) {
	filter := "{}"
	sortBy := "{}"
	respBytes := c.doPostRequestCommon(EndpointReasons, filter, sortBy)
	reasons := ReasonsResponse{}
	if err := json.Unmarshal(respBytes, &reasons); err != nil {
		return nil, err
	}

	return reasons.Data, nil

}

func (c *Client) GetMyHolydays(companyCurrentYear int, userHolidayIds []string) ([]Holiday, error) {
	dateStart := fmt.Sprintf("%d-01-01T00:00:00.000Z", companyCurrentYear)
	dateEnd := fmt.Sprintf("%d-12-31T23:59:59.999Z", companyCurrentYear)
	ids := fmt.Sprintf(`"%s"`, strings.Join(userHolidayIds, `","`))

	filter := fmt.Sprintf(`{"_id":{"$in":[%s]},"dates":{"$gte":"%s","$lte":"%s"}}`, ids, dateStart, dateEnd)
	sortBy := `{"date":1}`
	respBytes := c.doPostRequestCommon(EndpointHolidays, filter, sortBy)
	c.logger.Debug(string(respBytes))
	holidays := HolidaysResponse{}
	if err := json.Unmarshal(respBytes, &holidays); err != nil {
		return nil, err
	}

	return holidays.Data, nil
}

func (c *Client) GetMyAbsences(userID string, companyCurrentYear int) ([]Absence, error) {
	filter := fmt.Sprintf(`{"assignedToId":"%s","start":{"$gte":"%d-01-01"}}`, userID, companyCurrentYear-1)
	sortBy := `{"start":-1}`
	respBytes := c.doPostRequestCommon(EndpointAbsences, filter, sortBy)
	c.logger.Debug(string(respBytes))
	absences := AbsencesResponse{}
	if err := json.Unmarshal(respBytes, &absences); err != nil {
		return nil, err
	}

	return absences.Data, nil
}

func (c *Client) ClockInApi(userID string) (*TimeSpan, error) {
	now := time.Now()
	nowStr := now.Format("2006-01-02T15:04:05")
	time.Sleep(1 * time.Second)
	payload := strings.NewReader(`{
	  	"userId": "` + userID + `",
		"start":"` + nowStr + `.001Z",
	  	"end": null,
	  	"timezoneName": "CEST",
	  	"timezone": "+0100",
	  	"type": "work"
	}`)
	respBytes := c.doPostRequest(EndpointTimeSpanCreate, payload)
	tSpan := TimeSpan{}
	if err := json.Unmarshal(respBytes, &tSpan); err != nil {
		return nil, err
	}

	return &tSpan, nil

}

func (c *Client) ClockOutApi(timeSpan *TimeSpan) string {
	now := time.Now()
	nowStr := now.Format("2006-01-02T15:04:05")
	startStr := timeSpan.Start.Format("2006-01-02T15:04:05")
	time.Sleep(1 * time.Second)
	payload := strings.NewReader(`{
		"start":"` + startStr + `.001Z",
		"end":"` + nowStr + `.001Z",
	  	"timezoneName": "CEST",
	  	"timezone": "+0100",
	}`)
	respBytes := c.doPutRequest(fmt.Sprintf("%s/%s", EndpointTimeSpan, timeSpan.Id), payload)
	c.logger.Infof("%s", string(respBytes))
	return string(respBytes)
}
