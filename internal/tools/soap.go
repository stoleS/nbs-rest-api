package tools

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/stoleS/nbs-rest/api"
)

func InitSoap() *SoapConfig {
	var soapConfig = &SoapConfig{
		Username:   GetUsername(),
		Password:   GetPassword(),
		LicenceKey: GetLicenceKey(),
		XMLNs:      "http://communicationoffice.nbs.rs",
	}

	return soapConfig
}

var instance *Soap

func SOAP() *Soap {
	var once sync.Once
	if instance == nil {
		once.Do(
			func() {
				soapInstance := InitSoap()
				instance = &Soap{Config: *soapInstance}
			})
	}

	return instance
}

type Soap struct {
	Config SoapConfig
}

type SoapConfig struct {
	XMLNs      string
	Username   string
	Password   string
	LicenceKey string
}

type SoapRequest struct {
	XMLName     xml.Name `xml:"soap12:Envelope"`
	XMLNsXsi    string   `xml:"xmlns:xsi,attr"`
	XMLNsXsd    string   `xml:"xmlns:xsd,attr"`
	XMLNsSoap12 string   `xml:"xmlns:soap12,attr"`
	Header      SoapHeader
	Body        any
}

type SoapHeader struct {
	XMLName xml.Name `xml:"soap12:Header"`
	NBSUrl  NBSUrl   `xml:"AuthenticationHeader"`
}

type SoapErorr struct {
	Body struct {
		Fault struct {
			Code struct {
				Value string `xml:"Value"`
			} `xml:"Code"`
			Reason struct {
				Text struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
				} `xml:"Text"`
			} `xml:"Reason"`
			Detail struct {
				ErrorInfo struct {
					ErrorType    string `xml:"ErrorType"`
					ErrorCode    string `xml:"ErrorCode"`
					ErrorMessage string `xml:"ErrorMessage"`
				} `xml:"ErrorInfo"`
			} `xml:"detail"`
		} `xml:"Fault"`
	} `xml:"Body,omitempty"`
}

type NBSUrl struct {
	XMLNs     string `xml:"xmlns,attr"`
	UserName  string `xml:"UserName"`
	Password  string `xml:"Password"`
	LicenceId string `xml:"LicenceID"`
}

func (s *Soap) BuildSoapHeader() *SoapHeader {
	nbsUrl := &NBSUrl{
		XMLNs:     s.Config.XMLNs,
		UserName:  s.Config.Username,
		Password:  s.Config.Password,
		LicenceId: s.Config.LicenceKey,
	}

	header := &SoapHeader{NBSUrl: *nbsUrl}

	return header
}

func (s *Soap) BuildRequest(reqEnvelope any) *[]byte {
	header := s.BuildSoapHeader()

	request := &SoapRequest{
		XMLNsXsi:    "http://www.w3.org/2001/XMLSchema-instance",
		XMLNsXsd:    "http://www.w3.org/2001/XMLSchema",
		XMLNsSoap12: "http://www.w3.org/2003/05/soap-envelope",
		Header:      *header,
		Body:        reqEnvelope,
	}

	out, err := xml.MarshalIndent(request, "", " ")
	if err != nil {
		log.Error(err)
	}

	return &out
}

func (s *Soap) ExecuteRequest(reqType string, url string, body io.Reader) ([]byte, error) {
	client := &http.Client{}
	var maybeErrResponse SoapErorr

	req, err := http.NewRequest(reqType, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	xml.Unmarshal([]byte(data), &maybeErrResponse)
	if maybeErrResponse.Body.Fault.Code.Value != "" {
		return data, errors.New(maybeErrResponse.Body.Fault.Detail.ErrorInfo.ErrorType)
	}

	return data, nil
}

func (s *Soap) HandleError(w http.ResponseWriter, soapErrorRes *[]byte, err error) {
	log.Error(err)
	w.WriteHeader(http.StatusBadRequest)

	var soapErrorEnvelope *SoapErorr

	xml.Unmarshal([]byte(*soapErrorRes), &soapErrorEnvelope)
	err = json.NewEncoder(w).Encode(soapErrorEnvelope.Body)
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
}
