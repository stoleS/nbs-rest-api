package exchange_rate_service

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/stoleS/nbs-rest/api"
	"github.com/stoleS/nbs-rest/internal/tools"
)

type ReqEnvelope struct {
	XMLName                xml.Name                      `xml:"soap12:Body"`
	GetCurrentExchangeRate GetCurrentExchangeRateReqBody `xml:"GetCurrentExchangeRate"`
}

type GetCurrentExchangeRateReqBody struct {
	XMLNs                  string `xml:"xmlns,attr"`
	ExchangeRateListTypeID int    `xml:"exchangeRateListTypeID"`
}

type ResEnvelope struct {
	Body struct {
		GetCurrentExchangeRateResponse struct {
			GetCurrentExchangeRateResult struct {
				Diffgram struct {
					ExchangeRateDataSet struct {
						ExchangeRate []struct {
							Text                   string `xml:",chardata"`
							ID                     string `xml:"id,attr"`
							RowOrder               string `xml:"rowOrder,attr"`
							ExchangeRateListNumber string `xml:"ExchangeRateListNumber"`
							Date                   string `xml:"Date"`
							CreateDate             string `xml:"CreateDate"`
							DateTo                 string `xml:"DateTo"`
							ExchangeRateListTypeID string `xml:"ExchangeRateListTypeID"`
							CurrencyGroupID        string `xml:"CurrencyGroupID"`
							CurrencyCode           string `xml:"CurrencyCode"`
							CurrencyCodeNumChar    string `xml:"CurrencyCodeNumChar"`
							CurrencyCodeAlfaChar   string `xml:"CurrencyCodeAlfaChar"`
							CurrencyNameSerCyrl    string `xml:"CurrencyNameSerCyrl"`
							CurrencyNameSerLat     string `xml:"CurrencyNameSerLat"`
							CurrencyNameEng        string `xml:"CurrencyNameEng"`
							CountryNameSerCyrl     string `xml:"CountryNameSerCyrl"`
							CountryNameSerLat      string `xml:"CountryNameSerLat"`
							CountryNameEng         string `xml:"CountryNameEng"`
							Unit                   string `xml:"Unit"`
							BuyingRate             string `xml:"BuyingRate"`
							MiddleRate             string `xml:"MiddleRate"`
							SellingRate            string `xml:"SellingRate"`
							FixingRate             string `xml:"FixingRate"`
						} `xml:"ExchangeRate"`
					} `xml:"ExchangeRateDataSet"`
				} `xml:"diffgram"`
			} `xml:"GetCurrentExchangeRateResult"`
		} `xml:"GetCurrentExchangeRateResponse"`
	} `xml:"Body"`
}

func GetCurrentExchangeRate(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	SOAP := *tools.SOAP()

	reqEnvelope := &ReqEnvelope{
		GetCurrentExchangeRate: GetCurrentExchangeRateReqBody{
			ExchangeRateListTypeID: 2,
			XMLNs:                  SOAP.Config.XMLNs,
		},
	}

	request := SOAP.BuildRequest(*reqEnvelope)

	res, err := SOAP.ExecuteRequest("POST", GET_CURRENT_EXCHANGE_RATE_URL, bytes.NewBuffer([]byte(*request)))
	if err != nil {
		fmt.Println(err)
	}

	var response = &ResEnvelope{}

	xml.Unmarshal([]byte(res), response)

	err = json.NewEncoder(w).Encode(response.Body.GetCurrentExchangeRateResponse.GetCurrentExchangeRateResult.Diffgram.ExchangeRateDataSet)
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
}
