package handlers

import (
	"net/http"

	"github.com/stoleS/nbs-rest/internal/handlers/core_service"
	"github.com/stoleS/nbs-rest/internal/handlers/exchange_rate_service"
)

func noop(w http.ResponseWriter, r *http.Request) {}

func Handler(r *http.ServeMux) {
	// CoreService
	r.HandleFunc("POST /CoreService/GetBank", core_service.GetBank)
	r.HandleFunc("POST /CoreService/GetBankStatus", core_service.GetBankStatus)
	r.HandleFunc("POST /CoreService/GetBankType", core_service.GetBankType)
	r.HandleFunc("POST /CoreService/GetCompanyCount", core_service.GetCompanyCount)
	r.HandleFunc("POST /CoreService/GetCompany", core_service.GetCompany)
	r.HandleFunc("POST /CoreService/GetCompanyStatus", noop)
	r.HandleFunc("POST /CoreService/GetCompanyType", noop)
	r.HandleFunc("POST /CoreService/GetCountry", noop)
	r.HandleFunc("POST /CoreService/GetCountryByCurrency", noop)
	r.HandleFunc("POST /CoreService/GetCurrency", noop)
	r.HandleFunc("POST /CoreService/GetCurrencyByCountry", noop)
	r.HandleFunc("POST /CoreService/GetCurrencyConvertible", noop)
	r.HandleFunc("POST /CoreService/GetCurrencyGroup", noop)

	// ExchangeRateService
	r.HandleFunc("POST /ExchangeRateService/GetCurrentExchangeRate", exchange_rate_service.GetCurrentExchangeRate)
	r.HandleFunc("POST /ExchangeRateService/GetCurrentExchangeRateList", noop)
	r.HandleFunc("POST /ExchangeRateService/GetExchangeRateByCurrency", noop)
	r.HandleFunc("POST /ExchangeRateService/GetExchangeRateByDate", noop)
	r.HandleFunc("POST /ExchangeRateService/GetExchangeRateByListNumber", noop)
	r.HandleFunc("POST /ExchangeRateService/GetExchangeRateList", noop)
	r.HandleFunc("POST /ExchangeRateService/GetExchangeRateListCount", noop)
	r.HandleFunc("POST /ExchangeRateService/GetExchangeRateListType", noop)
	r.HandleFunc("POST /ExchangeRateService/GetCurrentExchangeRateByRateType", noop)
	r.HandleFunc("POST /ExchangeRateService/GetExchangeRateByRateType", noop)
	r.HandleFunc("POST /ExchangeRateService/GetExchangeRateRsdEur", noop)
	r.HandleFunc("POST /ExchangeRateService/GetExchangeRateRsdEurByPeriod", noop)
	r.HandleFunc("POST /ExchangeRateService/GetCurrentExchangeRateRsdEur", noop)
	r.HandleFunc("POST /ExchangeRateService/GetExchangeRateRsdEurType", noop)

	// CurrentExchangeRateService
	r.HandleFunc("POST /CurrentExchangeRateService/GetCurrentExchangeRate", noop)
	r.HandleFunc("POST /CurrentExchangeRateService/GetCurrentExchangeRateList", noop)
	r.HandleFunc("POST /CurrentExchangeRateService/GetCurrentExchangeRateByRateType", noop)
	r.HandleFunc("POST /CurrentExchangeRateService/GetExchangeRateListType", noop)

	// BankExchangeRateService
	r.HandleFunc("POST /BankExchangeRateService/GetBankExchangeRate", noop)
	r.HandleFunc("POST /BankExchangeRateService/GetBankExchangeRateByBank", noop)
	r.HandleFunc("POST /BankExchangeRateService/GetBankExchangeRateByCurrency", noop)
	r.HandleFunc("POST /BankExchangeRateService/GetBankExchangeRateByPeriod", noop)

	// CompanyAccountService
	r.HandleFunc("POST /CompanyAccountService/GetCompanyAccountCount", noop)
	r.HandleFunc("POST /CompanyAccountService/GetCompanyAccount", noop)
	r.HandleFunc("POST /CompanyAccountService/GetCompanyAccountTop", noop)
	r.HandleFunc("POST /CompanyAccountService/GetCompanyAccountUpdatedCount", noop)
	r.HandleFunc("POST /CompanyAccountService/GetCompanyAccountUpdated", noop)
	r.HandleFunc("POST /CompanyAccountService/GetCompanyAccountByNationalIdentificationNumber", noop)
	r.HandleFunc("POST /CompanyAccountService/GetCompanyAccountOriginByNationalIdNumber", noop)
	r.HandleFunc("POST /CompanyAccountService/GetCompanyAccountStatus", noop)
	r.HandleFunc("POST /CompanyAccountService/GetCompanyAccountType", noop)

	// DebtorService
	r.HandleFunc("POST /DebtorService/GetDebtorCount", noop)
	r.HandleFunc("POST /DebtorService/GetDebtor", noop)
	r.HandleFunc("POST /DebtorService/GetDebtorIlliquidityDayCount", noop)
	r.HandleFunc("POST /DebtorService/GetDebtorIlliquidityDay", noop)
	r.HandleFunc("POST /DebtorService/GetEnforcedCollectionDebtorBlockadeStatus", noop)
	r.HandleFunc("POST /DebtorService/GetCourtDecisionCount", noop)
	r.HandleFunc("POST /DebtorService/GetCourtDecision", noop)

	// EnforcedCollectionDecisionService
	r.HandleFunc("POST /EnforcedCollectionDecisionService/GetEnforcedCollectionDecisionCount", noop)
	r.HandleFunc("POST /EnforcedCollectionDecisionService/GetEnforcedCollectionDecision", noop)
	r.HandleFunc("POST /EnforcedCollectionDecisionService/GetEnforcedCollectionDecisionIssuerCount", noop)
	r.HandleFunc("POST /EnforcedCollectionDecisionService/GetEnforcedCollectionDecisionIssuer", noop)
	r.HandleFunc("POST /EnforcedCollectionDecisionService/GetEnforcedCollectionDecisionIssuerType", noop)

	// ValPanFundService
	r.HandleFunc("POST /ValPanFundService/GetValPanFund", noop)
	r.HandleFunc("POST /ValPanFundService/GetValPanFundCompany", noop)
	r.HandleFunc("POST /ValPanFundService/GetValPanFundCompanyStatus", noop)
	r.HandleFunc("POST /ValPanFundService/GetValPanFundReport", noop)
	r.HandleFunc("POST /ValPanFundService/GetValPanFundStatus", noop)
	r.HandleFunc("POST /ValPanFundService/GetFondex", noop)
	r.HandleFunc("POST /ValPanFundService/GetFondexByPeriod", noop)

	// BeoinaService
	r.HandleFunc("POST /BeoinaService/GetBeonia", noop)
	r.HandleFunc("POST /BeoinaService/GetBeoniaByPeriod", noop)
	r.HandleFunc("POST /BeoinaService/GetBeoniaByYear", noop)

	// InsuranceMarketService
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceActivityType", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceActivityTypeGroup", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceJobType", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceLegalEntityCount", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceLegalEntity", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceLegalEntityResidentType", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceLegalEntityStatus", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceLegalEntityType", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceMarketEntityCount", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceMarketEntity", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceMarketEntityJob", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceMarketEntityStatus", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceMarketEntityType", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceMarketMeasureCount", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceMarketMeasure", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceMarketMeasureStatus", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceMarketMeasureType", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceMarketRelationCount", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceMarketRelation", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceMarketRelationStatus", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceMarketRelationType", noop)
	r.HandleFunc("POST /InsuranceMarketService/GetInsuranceLegalEntityFileLink", noop)

	// LicenceService
	r.HandleFunc("POST /LicenceService/GetCustomer", noop)
	r.HandleFunc("POST /LicenceService/GetLicence", noop)
	r.HandleFunc("POST /LicenceService/GetLicenceByLicenceID", noop)
	r.HandleFunc("POST /LicenceService/GetLicenceHistory", noop)
	r.HandleFunc("POST /LicenceService/GetLicenceHistoryItem", noop)
	r.HandleFunc("POST /LicenceService/GetLicenceItem", noop)
	r.HandleFunc("POST /LicenceService/GetLicenceItemByLicenceID", noop)
	r.HandleFunc("POST /LicenceService/GetProduct", noop)
	r.HandleFunc("POST /LicenceService/GetUser", noop)
}
