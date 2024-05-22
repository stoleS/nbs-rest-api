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

type GetCompanyReqEnvelope struct {
	XMLName    xml.Name          `xml:"soap12:Body"`
	GetCompany GetCompanyReqBody `xml:"GetCompany"`
}

type GetCompanyReqBody struct {
	XMLNs                        string `xml:"xmlns,attr"`
	CompanyID                    int    `xml:"companyID,omitempty"`
	CompanyCode                  int    `xml:"companyCode,omitempty"`
	Name                         string `xml:"name,omitempty"`
	City                         string `xml:"city,omitempty"`
	NationalIdentificationNumber int    `xml:"nationalIdentificationNumber,omitempty"`
	TaxIdentificationNumber      int    `xml:"taxIdentificationNumber,omitempty"`
	StartItemNumber              int    `xml:"startItemNumber,omitempty"`
	EndItemNumber                int    `xml:"endItemNumber,omitempty"`
}

type _GetCompanyReqBody struct {
	CompanyID                    int
	CompanyCode                  int
	Name                         string
	City                         string
	NationalIdentificationNumber int
	TaxIdentificationNumber      int
	StartItemNumber              int
	EndItemNumber                int
}

type GetCompanyResEnvelope struct {
	Body struct {
		GetCompanyResponse struct {
			GetCompanyResult struct {
				Diffgram struct {
					CompanyDataSet struct {
						Company struct {
							CompanyID                    string `xml:"CompanyID"`
							NationalIdentificationNumber string `xml:"NationalIdentificationNumber"`
							TaxIdentificationNumber      string `xml:"TaxIdentificationNumber"`
							Name                         string `xml:"Name"`
							ShortName                    string `xml:"ShortName"`
							Address                      string `xml:"Address"`
							City                         string `xml:"City"`
							Municipality                 string `xml:"Municipality"`
							Region                       string `xml:"Region"`
							PostalCode                   string `xml:"PostalCode"`
							ActivityName                 string `xml:"ActivityName"`
							RegistrationDate             string `xml:"RegistrationDate"`
							FoundingDate                 string `xml:"FoundingDate"`
							CompanyTypeID                string `xml:"CompanyTypeID"`
							CompanyStatusID              string `xml:"CompanyStatusID"`
							NBSDate                      string `xml:"NBSDate"`
							Resource                     string `xml:"Resource"`
						} `xml:"Company"`
					} `xml:"CompanyDataSet"`
				} `xml:"diffgram"`
			} `xml:"GetCompanyResult"`
		} `xml:"GetCompanyResponse"`
	} `xml:"Body"`
}

func GetCompany(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	SOAP := *tools.SOAP()
	var reqBody _GetCompanyReqBody

	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		log.Error(err)
	}

	reqEnvelope := &GetCompanyReqEnvelope{
		GetCompany: GetCompanyReqBody{
			XMLNs:                        SOAP.Config.XMLNs,
			CompanyID:                    reqBody.CompanyID,
			CompanyCode:                  reqBody.CompanyCode,
			Name:                         reqBody.Name,
			City:                         reqBody.City,
			NationalIdentificationNumber: reqBody.NationalIdentificationNumber,
			TaxIdentificationNumber:      reqBody.TaxIdentificationNumber,
			StartItemNumber:              reqBody.StartItemNumber,
			EndItemNumber:                reqBody.EndItemNumber,
		},
	}

	request := SOAP.BuildRequest(*reqEnvelope)

	res, err := SOAP.ExecuteRequest("POST", GET_COMPANY_URL, bytes.NewBuffer([]byte(*request)))
	if err != nil {
		if res != nil {
			SOAP.HandleError(w, &res, err)
			return
		}
		api.RequestErrorHandler(w, err)
		return
	}

	var response = &GetCompanyResEnvelope{}

	xml.Unmarshal([]byte(res), response)

	err = json.NewEncoder(w).Encode(response.Body.GetCompanyResponse.GetCompanyResult.Diffgram.CompanyDataSet)
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
}
