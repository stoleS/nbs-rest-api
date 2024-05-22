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

type GetBankTypeReqEnvelope struct {
	XMLName       xml.Name           `xml:"soap12:Body"`
	GetBankStatus GetBankTypeReqBody `xml:"GetBankType"`
}

type GetBankTypeReqBody struct {
	XMLNs      string `xml:"xmlns,attr"`
	BankTypeID int    `xml:"bankTypeID,omitempty"`
}

type _GetBankTypeReqBody struct {
	BankTypeID int
}

type GetBankTypeResEnvelope struct {
	Body struct {
		GetBankTypeResponse struct {
			GetBankTypeResult struct {
				Diffgram struct {
					BankDataSet struct {
						BankType []struct {
							ID         string `xml:"id,attr"`
							RowOrder   string `xml:"rowOrder,attr"`
							HasChanges string `xml:"hasChanges,attr"`
							BankTypeID string `xml:"BankTypeID"`
							Name       string `xml:"Name"`
						} `xml:"BankType"`
					} `xml:"BankDataSet"`
				} `xml:"diffgram"`
			} `xml:"GetBankTypeResult"`
		} `xml:"GetBankTypeResponse"`
	} `xml:"Body"`
}

func GetBankType(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	SOAP := *tools.SOAP()
	var reqBody _GetBankTypeReqBody

	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		log.Error(err)
	}

	reqEnvelope := &GetBankTypeReqEnvelope{
		GetBankStatus: GetBankTypeReqBody{
			XMLNs:      SOAP.Config.XMLNs,
			BankTypeID: reqBody.BankTypeID,
		},
	}

	request := SOAP.BuildRequest(*reqEnvelope)

	res, err := SOAP.ExecuteRequest("POST", GET_BANK_TYPE_URL, bytes.NewBuffer([]byte(*request)))
	if err != nil {
		if res != nil {
			SOAP.HandleError(w, &res, err)
			return
		}
		api.RequestErrorHandler(w, err)
		return
	}

	var response = &GetBankTypeResEnvelope{}

	xml.Unmarshal([]byte(res), response)

	err = json.NewEncoder(w).Encode(response.Body.GetBankTypeResponse.GetBankTypeResult.Diffgram.BankDataSet)
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
}
