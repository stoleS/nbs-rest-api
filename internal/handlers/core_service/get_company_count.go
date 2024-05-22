package core_service

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/stoleS/nbs-rest/api"
	"github.com/stoleS/nbs-rest/internal/tools"
)

type GetCompanyCountReqEnvelope struct {
	XMLName         xml.Name               `xml:"soap12:Body"`
	GetCompanyCount GetCompanyCountReqBody `xml:"GetCompanyCount"`
}

type GetCompanyCountReqBody struct {
	XMLNs                        string `xml:"xmlns,attr"`
	CompanyID                    int    `xml:"companyID,omitempty"`
	CompanyCode                  int    `xml:"companyCode,omitempty"`
	Name                         string `xml:"name,omitempty"`
	City                         string `xml:"city,omitempty"`
	NationalIdentificationNumber int    `xml:"nationalIdentificationNumber,omitempty"`
	TaxIdentificationNumber      int    `xml:"taxIdentificationNumber,omitempty"`
}

type _GetCompanyCountReqBody struct {
	CompanyID                    int
	CompanyCode                  int
	Name                         string
	City                         string
	NationalIdentificationNumber int
	TaxIdentificationNumber      int
}

type GetCompanyCountResEnvelope struct {
	Body struct {
		GetCompanyCountResponse struct {
			GetCompanyCountResult string `xml:"GetCompanyCountResult"`
		} `xml:"GetCompanyCountResponse"`
	} `xml:"Body"`
}

func GetCompanyCount(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	SOAP := *tools.SOAP()
	var reqBody _GetCompanyCountReqBody

	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		log.Error(err)
	}

	reqEnvelope := &GetCompanyCountReqEnvelope{
		GetCompanyCount: GetCompanyCountReqBody{
			XMLNs:                        SOAP.Config.XMLNs,
			CompanyID:                    reqBody.CompanyID,
			CompanyCode:                  reqBody.CompanyCode,
			Name:                         reqBody.Name,
			City:                         reqBody.City,
			NationalIdentificationNumber: reqBody.NationalIdentificationNumber,
			TaxIdentificationNumber:      reqBody.TaxIdentificationNumber,
		},
	}

	request := SOAP.BuildRequest(*reqEnvelope)

	res, err := SOAP.ExecuteRequest("POST", GET_COMPANY_COUNT_URL, bytes.NewBuffer([]byte(*request)))
	if err != nil {
		if res != nil {
			SOAP.HandleError(w, &res, err)
			return
		}
		api.RequestErrorHandler(w, err)
		return
	}

	var response = &GetCompanyCountResEnvelope{}

	xml.Unmarshal([]byte(res), response)

	err = json.NewEncoder(w).Encode(response.Body.GetCompanyCountResponse)
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
}
