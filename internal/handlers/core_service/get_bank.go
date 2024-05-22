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

type GetBankReqEnvelope struct {
	XMLName xml.Name       `xml:"soap12:Body"`
	GetBank GetBankReqBody `xml:"GetBank"`
}

type GetBankReqBody struct {
	XMLNs                        string `xml:"xmlns,attr"`
	BankID                       int    `xml:"bankID,omitempty"`
	BankCode                     int    `xml:"bankCode,omitempty"`
	NationalIdentificationNumber int    `xml:"nationalIdentificationNumber,omitempty"`
	TaxIdentificationNumber      int    `xml:"taxIdentificationNumber,omitempty"`
	Date                         string `xml:"date,omitempty"`
}

type _GetBankReqBody struct {
	BankID                       int
	BankCode                     int
	NationalIdentificationNumber int
	TaxIdentificationNumber      int
	Date                         string
}

type GetBankResEnvelope struct {
	Body struct {
		GetBankResponse struct {
			GetBankResult struct {
				Diffgram struct {
					BankDataSet struct {
						Bank []struct {
							ID                           string `xml:"id,attr"`
							RowOrder                     string `xml:"rowOrder,attr"`
							BankID                       string `xml:"BankID"`
							BankHistoryID                string `xml:"BankHistoryID"`
							StartDate                    string `xml:"StartDate"`
							BankCode                     string `xml:"BankCode"`
							NationalIdentificationNumber string `xml:"NationalIdentificationNumber"`
							Name                         string `xml:"Name"`
							Address                      string `xml:"Address"`
							City                         string `xml:"City"`
							PostalCode                   string `xml:"PostalCode"`
							Phone                        string `xml:"Phone"`
							Fax                          string `xml:"Fax"`
							Email                        string `xml:"Email"`
							WebAddress                   string `xml:"WebAddress"`
							Director                     string `xml:"Director"`
							BankTypeID                   string `xml:"BankTypeID"`
							BankStatusID                 string `xml:"BankStatusID"`
						} `xml:"Bank"`
					} `xml:"BankDataSet"`
				} `xml:"diffgram"`
			} `xml:"GetBankResult"`
		} `xml:"GetBankResponse"`
	} `xml:"Body"`
}

func GetBank(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	SOAP := *tools.SOAP()
	var reqBody _GetBankReqBody

	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		log.Error(err)
	}

	reqEnvelope := &GetBankReqEnvelope{
		GetBank: GetBankReqBody{
			XMLNs:                        SOAP.Config.XMLNs,
			BankID:                       reqBody.BankID,
			BankCode:                     reqBody.BankCode,
			NationalIdentificationNumber: reqBody.NationalIdentificationNumber,
			TaxIdentificationNumber:      reqBody.TaxIdentificationNumber,
			Date:                         reqBody.Date,
		},
	}

	request := SOAP.BuildRequest(*reqEnvelope)

	res, err := SOAP.ExecuteRequest("POST", GET_BANK_URL, bytes.NewBuffer([]byte(*request)))
	if err != nil {
		if res != nil {
			SOAP.HandleError(w, &res, err)
			return
		}
		api.RequestErrorHandler(w, err)
		return
	}

	var response = &GetBankResEnvelope{}

	xml.Unmarshal([]byte(res), response)

	err = json.NewEncoder(w).Encode(response.Body.GetBankResponse.GetBankResult.Diffgram.BankDataSet)
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
}
