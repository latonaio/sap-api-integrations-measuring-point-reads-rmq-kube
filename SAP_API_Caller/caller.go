package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-measuring-point-reads-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL      string
	apiKey       string
	outputQueues []string
	outputter    RMQOutputter
	log          *logger.Logger
}

func NewSAPAPICaller(baseUrl string, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:      baseUrl,
		apiKey:       GetApiKey(),
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:          l,
	}
}

func (c *SAPAPICaller) AsyncGetMeasuringPoint(measuringPoint, equipment string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(measuringPoint)
				wg.Done()
			}()
		case "Equipment":
			func() {
				c.Equipment(equipment)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(measuringPoint string) {
	data, err := c.callMeasuringPointSrvAPIRequirementHeader("MeasuringPoint", measuringPoint)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "MeasuringpointHeaderData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callMeasuringPointSrvAPIRequirementHeader(api, measuringPoint string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "api_measuringpoint/srvd_a2x/sap/measuringpoint/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, measuringPoint)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("API status code %d. API request failed", resp.StatusCode)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Equipment(equipment string) {
	data, err := c.callMeasuringPointSrvAPIRequirementEquipment("MeasuringPoint", equipment)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "MeasuringpointHeaderData"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callMeasuringPointSrvAPIRequirementEquipment(api, equipment string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "api_measuringpoint/srvd_a2x/sap/measuringpoint/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithEquipment(req, equipment)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("API status code %d. API request failed", resp.StatusCode)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, measuringPoint string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("MeasuringPoint eq '%s'", measuringPoint))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithEquipment(req *http.Request, equipment string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Equipment eq '%s'", equipment))
	req.URL.RawQuery = params.Encode()
}
