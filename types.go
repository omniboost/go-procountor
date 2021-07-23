package bexio

type Accounts []Account

type Account struct {
	CurrencyID     int    `json:"currency_id"`
	ID             int    `json:"id"`
	UUID           string `json:"uuid"`
	Type           string `json:"type"`
	AccountNo      string `json:"account_no"`
	Name           string `json:"name"`
	TaxID          int    `json:"tax_id"`
	AccountGroupID int    `json:"account_group_id"`
	InternalKey    string `json:"internal_key"`
	IsActive       bool   `json:"is_active"`
	IsLocked       bool   `json:"is_locked"`
}

type Taxes []Tax

type Tax struct {
	ID                int    `json:"id"`
	UUID              string `json:"uuid"`
	Name              string `json:"name"`
	Code              string `json:"code"`
	Digit             string `json:"digit"`
	Type              string `json:"type"`
	AccountID         int    `json:"account_id"`
	TaxSettlementType string `json:"tax_settlement_type"`
	Value             Number `json:"value"`
	NetTaxValue       Number `json:"net_tax_value"`
	StartYear         int    `json:"start_year"`
	EndYear           int    `json:"end_year"`
	IsActive          bool   `json:"is_active"`
	DisplayName       string `json:"display_name"`
	StartMonth        int    `json:"start_month"`
	EndMonth          int    `json:"end_month"`
}
