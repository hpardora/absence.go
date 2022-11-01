package absence

import (
	"encoding/json"
	"fmt"
	"github.com/hiyosi/hawk"
	"github.com/hpardora/absence.go/pkg/randStrings"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
	"time"
)

const baseURL string = "https://app.absence.io/api/v2"

type Client struct {
	hawkClient *hawk.Client
	config     *Config
	user       *User
	company    *Company
	logger     *log.Logger
}

func New(c *Config) *Client {
	logger := log.New()
	nonce := randStrings.RandStringRunes(6)
	hawkClient := hawk.NewClient(
		&hawk.Credential{
			ID:  c.ID,
			Key: c.Key,
			Alg: hawk.SHA256,
		},
		&hawk.Option{
			TimeStamp: time.Now().Unix(),
			Nonce:     nonce,
			Ext:       "some-app-data",
		},
	)
	client := &Client{
		hawkClient: hawkClient,
		config:     c,
		logger:     logger,
	}
	// Retrieve User Data
	me, err := client.me()
	if err != nil {
		return nil
	}
	client.user = me

	// Retrieve User Company Data
	comp, err := client.myCompany()
	if err != nil {
		return nil
	}
	client.company = comp

	logger.Infof("connected as %s %s on company %s", me.FirstName, me.LastName, comp.Name)

	return client
}

func (c *Client) buildHeader(method string, path string) string {
	// build request header
	header, _ := c.hawkClient.Header(method, path)
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
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Authorization", header)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	c.logger.Infof("method: %s\tresult_code: %d\t url: %s", method, res.StatusCode, url)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}

func (c *Client) me() (*User, error) {
	url := fmt.Sprintf("%s/users/%s", baseURL, c.config.ID)
	method := "GET"
	payload := strings.NewReader(``)

	respBytes := c.doRequest(url, method, payload)

	me := User{}
	if err := json.Unmarshal(respBytes, &me); err != nil {
		return nil, err
	}
	return &me, nil
}

func (c *Client) myCompany() (*Company, error) {
	url := fmt.Sprintf("%s/companies/%s", baseURL, c.user.Company)
	method := "GET"
	payload := strings.NewReader(``)

	respBytes := c.doRequest(url, method, payload)
	comp := Company{}
	if err := json.Unmarshal(respBytes, &comp); err != nil {
		return nil, err
	}
	return &comp, nil

}

func (c *Client) getHolydays(filter string) []byte {
	url := fmt.Sprintf("%s/holidays", baseURL)
	method := "POST"
	payload := strings.NewReader(`{
		"skip":0,
		"limit":5000,
		"filter": {` + filter + `},
		"sortBy":{
			"date":1
		}
	}`)

	respBytes := c.doRequest(url, method, payload)
	return respBytes
}

func (c *Client) getHolidaysRegion() []byte {
	url := fmt.Sprintf("%s/holidayregions", baseURL)
	method := "POST"
	payload := strings.NewReader(`{
		"skip":0,
		"limit":5000,
		"filter": {},
		"sortBy":{
			"date":1
		}
	}`)

	respBytes := c.doRequest(url, method, payload)
	return respBytes
}

func (c *Client) GetAllHolydays() ([]HolidayDetail, error) {
	// Get Holydays based on Company
	filter := `"$and":[{"company":"` + c.user.Company + `"}]`
	respBytes := c.getHolydays(filter)
	holidays := Holidays{}
	if err := json.Unmarshal(respBytes, &holidays); err != nil {
		return nil, err
	}

	// Get Holydays based on
	respBytes = c.getHolidaysRegion()
	regionHolidays := HolidayRegion{}

	if err := json.Unmarshal(respBytes, &regionHolidays); err != nil {
		return nil, err
	}
	var hrd HolidayRegionDetail
	for _, value := range regionHolidays.Data {
		if value.Name == c.company.Country {
			hrd = value
			break
		}
	}
	dateStart := fmt.Sprintf("%d-01-01T00:00:00.000Z", c.company.CurrentCompanyYear)
	dateEnd := fmt.Sprintf("%d-12-31T23:59:59.999Z", c.company.CurrentCompanyYear)
	ids := fmt.Sprintf(`"%s"`, strings.Join(hrd.HolidayIds, `","`))

	filter = fmt.Sprintf(`"_id":{"$in":[%s]},"dates":{"$gte":"%s","$lte":"%s"}`, ids, dateStart, dateEnd)
	respBytes = c.getHolydays(filter)
	fmt.Println(string(respBytes))
	holidaysRegional := Holidays{}
	if err := json.Unmarshal(respBytes, &holidaysRegional); err != nil {
		return nil, err
	}

	return append(holidays.Data, holidaysRegional.Data...), nil
}
