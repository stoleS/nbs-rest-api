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

type GetBankStatusReqEnvelope struct {
	XMLName       xml.Name             `xml:"soap12:Body"`
	GetBankStatus GetBankStatusReqBody `xml:"GetBankStatus"`
}

type GetBankStatusReqBody struct {
	XMLNs        string `xml:"xmlns,attr"`
	BankStatusID int    `xml:"bankStatusID,omitempty"`
}

type _GetBankStatusReqBody struct {
	BankStatusID int
}

type GetBankStatusResEnvelope struct {
	Body struct {
		GetBankStatusResponse struct {
			GetBankStatusResult struct {
				Diffgram struct {
					BankDataSet struct {
						BankStatus []struct {
							ID           string `xml:"id,attr"`
							RowOrder     string `xml:"rowOrder,attr"`
							HasChanges   string `xml:"hasChanges,attr"`
							BankStatusID string `xml:"BankStatusID"`
							Name         string `xml:"Name"`
						} `xml:"BankStatus"`
					} `xml:"BankDataSet"`
				} `xml:"diffgram"`
			} `xml:"GetBankStatusResult"`
		} `xml:"GetBankStatusResponse"`
	} `xml:"Body"`
}

func GetBankStatus(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	SOAP := *tools.SOAP()
	var reqBody _GetBankStatusReqBody

	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		log.Error(err)
	}

	reqEnvelope := &GetBankStatusReqEnvelope{
		GetBankStatus: GetBankStatusReqBody{
			XMLNs:        SOAP.Config.XMLNs,
			BankStatusID: reqBody.BankStatusID,
		},
	}

	request := SOAP.BuildRequest(*reqEnvelope)

	res, err := SOAP.ExecuteRequest("POST", GET_BANK_STATUS_URL, bytes.NewBuffer([]byte(*request)))
	if err != nil {
		if res != nil {
			SOAP.HandleError(w, &res, err)
			return
		}
		api.RequestErrorHandler(w, err)
		return
	}

	var response = &GetBankStatusResEnvelope{}

	xml.Unmarshal([]byte(res), response)

	err = json.NewEncoder(w).Encode(response.Body.GetBankStatusResponse.GetBankStatusResult.Diffgram.BankDataSet)
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
}
