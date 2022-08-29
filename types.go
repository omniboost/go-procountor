package procountor

import (
	"github.com/cydev/zero"
	"github.com/omniboost/go-procountor/omitempty"
)

type LedgerAccounts []LedgerAccount

type LedgerAccount struct {
	LedgerAccountCode string `json:"ledgerAccountCode"`
	Name              string `json:"name"`
}

type Company struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Country              string `json:"country"`
	ProductVersion       string `json:"productVersion"`
	OperationType        string `json:"operationType"`
	AccountingOfficeName string `json:"accountingOfficeName"`
	AccountingOfficeID   int    `json:"accountingOfficeId"`
	CompanyAddress       struct {
		Street  string `json:"street"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"companyAddress"`
	BillingAddress struct {
		Street string `json:"street"`
		Zip    string `json:"zip"`
		City   string `json:"city"`
	} `json:"billingAddress"`
	Mva                bool `json:"mva"`
	InTradeRegister    bool `json:"inTradeRegister"`
	BusinessIdentifier struct {
		BrnType string `json:"brnType"`
		Code    string `json:"code"`
	} `json:"businessIdentifier"`
}

type Dimensions []Dimension

type Dimension struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Items []struct {
		ID          int    `json:"id"`
		CodeName    string `json:"codeName"`
		Status      string `json:"status,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"items,omitempty"`
}

type VatStatuses []VatStatus

type VatInformations []VatInformation

type VatInformation struct {
	Country     string `json:"country"`
	Percentages []struct {
		VatPercent float64 `json:"vatPercent"`
		Sales      bool    `json:"sales"`
		Purchase   bool    `json:"purchase"`
	} `json:"percentages"`
}

type VatStatus struct {
	VatStatus   int    `json:"vatStatus"`
	Description string `json:"description"`
	Sales       bool   `json:"sales"`
	Purchase    bool   `json:"purchase"`
}

type Invoices []Invoice

type Invoice struct {
	PartnerID              int                    `json:"partnerId,omitempty"`
	Type                   string                 `json:"type"`
	Status                 string                 `json:"status"`
	Date                   string                 `json:"date"`
	CounterParty           CounterParty           `json:"counterParty"`
	BillingAddress         Address                `json:"billingAddress,omitempty"`
	DeliveryAddress        Address                `json:"deliveryAddress,omitempty"`
	PaymentInfo            PaymentInfo            `json:"paymentInfo"`
	DeliveryTermsInfo      DeliveryTermsInfo      `json:"deliveryTermsInfo,omitempty"`
	ExtraInfo              ExtraInfo              `json:"extraInfo"`
	DiscountPercent        int                    `json:"discountPercent"`
	OrderReference         string                 `json:"orderReference"`
	InvoiceRows            InvoiceRows            `json:"invoiceRows"`
	VatStatus              int                    `json:"vatStatus"`
	OriginalInvoiceNumber  string                 `json:"originalInvoiceNumber"`
	DeliveryStartDate      string                 `json:"deliveryStartDate,omitempty"`
	DeliveryEndDate        string                 `json:"deliveryEndDate,omitempty"`
	DeliveryMethod         string                 `json:"deliveryMethod,omitempty"`
	DeliveryInstructions   string                 `json:"deliveryInstructions,omitempty"`
	InvoiceChannel         string                 `json:"invoiceChannel,omitempty"`
	InvoiceOperatorInfo    InvoiceOperatorInfo    `json:"invoiceOperatorInfo,omitempty"`
	PenaltyPercent         int                    `json:"penaltyPercent,omitempty"`
	Language               string                 `json:"language,omitempty"`
	InvoiceTemplateID      int                    `json:"invoiceTemplateId,omitempty"`
	AdditionalInformation  string                 `json:"additionalInformation,omitempty"`
	VatCountry             string                 `json:"vatCountry,omitempty"`
	Notes                  string                 `json:"notes,omitempty"`
	FactoringContractID    int                    `json:"factoringContractId,omitempty"`
	FactoringText          string                 `json:"factoringText,omitempty"`
	TravelInformationItems TravelInformationItems `json:"travelInformationItems,omitempty"`
	// InvoiceApprovalInformation struct {
	// 	Acceptors []struct {
	// 		UserID int `json:"userId"`
	// 	} `json:"acceptors"`
	// 	Verifiers []struct {
	// 		UserID int `json:"userId"`
	// 	} `json:"verifiers"`
	// } `json:"invoiceApprovalInformation,omitempty"`
	OrderNumber     string `json:"orderNumber,omitempty"`
	AgreementNumber string `json:"agreementNumber,omitempty"`
	AccountingCode  string `json:"accountingCode,omitempty"`
	DeliverySite    string `json:"deliverySite,omitempty"`
	TenderReference string `json:"tenderReference,omitempty"`
	IsOffer         bool   `json:"isOffer,omitempty"`
}

func (i Invoice) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

func (i Invoice) IsEmpty() bool {
	return zero.IsZero(i)
}

type CounterParty struct {
	ContactPersonName   string      `json:"contactPersonName"`
	Identifier          string      `json:"identifier"`
	TaxCode             string      `json:"taxCode"`
	CustomerNumber      string      `json:"customerNumber"`
	Email               string      `json:"email"`
	CounterPartyAddress Address     `json:"counterPartyAddress"`
	BankAccount         BankAccount `json:"bankAccount,omitempty"`
	EinvoiceAddress     struct {
		Operator string `json:"operator"`
		Address  string `json:"address"`
		EdiID    string `json:"ediId"`
	} `json:"einvoiceAddress"`
}

func (cp CounterParty) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(cp)
}

func (cp CounterParty) IsEmpty() bool {
	return zero.IsZero(cp)
}

type Address struct {
	Name        string `json:"name"`
	Specifier   string `json:"specifier"`
	Street      string `json:"street"`
	Zip         string `json:"zip"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Subdivision string `json:"subdivision"`
}

func (a Address) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(a)
}

func (a Address) IsEmpty() bool {
	return zero.IsZero(a)
}

type PaymentInfo struct {
	PaymentMethod         string      `json:"paymentMethod"`
	Currency              string      `json:"currency"`
	BankAccount           BankAccount `json:"bankAccount,omitempty"`
	DueDate               string      `json:"dueDate"`
	CurrencyRate          int         `json:"currencyRate"`
	PaymentTermPercentage int         `json:"paymentTermPercentage,omitempty"`
	CashDiscount          struct {
		NumberOfDays       int `json:"numberOfDays"`
		DiscountPercentage int `json:"discountPercentage"`
	} `json:"cashDiscount,omitempty"`
	BankReferenceCode     string `json:"bankReferenceCode,omitempty"`
	BankReferenceCodeType string `json:"bankReferenceCodeType,omitempty"`
	ClearingCode          string `json:"clearingCode,omitempty"`
}

func (i PaymentInfo) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

func (i PaymentInfo) IsEmpty() bool {
	return zero.IsZero(i)
}

type InvoiceRows []InvoiceRow

type InvoiceRow struct {
	ProductID       int    `json:"productId,omitempty"`
	Product         string `json:"product"`
	ProductCode     string `json:"productCode,omitempty"`
	Quantity        int    `json:"quantity"`
	Unit            string `json:"unit"`
	UnitPrice       int    `json:"unitPrice"`
	DiscountPercent int    `json:"discountPercent"`
	VatPercent      int    `json:"vatPercent"`
	VatStatus       int    `json:"vatStatus"`
	Comment         string `json:"comment"`
	StartDate       string `json:"startDate,omitempty"`
	EndDate         string `json:"endDate,omitempty"`
}

type TravelInformationItems []TravelInformationItem

type TravelInformationItem struct {
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
	Places    string `json:"places"`
	Purpose   string `json:"purpose"`
}

func (i TravelInformationItem) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

func (i TravelInformationItem) IsEmpty() bool {
	return zero.IsZero(i)
}

type InvoiceOperatorInfo struct {
	Operator         string `json:"operator"`
	ReceivingAddress string `json:"receivingAddress"`
}

func (i InvoiceOperatorInfo) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

func (i InvoiceOperatorInfo) IsEmpty() bool {
	return zero.IsZero(i)
}

type DeliveryTermsInfo struct {
	Name         string `json:"name"`
	Municipality string `json:"municipality"`
}

func (i DeliveryTermsInfo) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

func (i DeliveryTermsInfo) IsEmpty() bool {
	return zero.IsZero(i)
}

type ExtraInfo struct {
	AccountingByRow      bool `json:"accountingByRow"`
	UnitPricesIncludeVat bool `json:"unitPricesIncludeVat"`
}

func (i ExtraInfo) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

func (i ExtraInfo) IsEmpty() bool {
	return zero.IsZero(i)
}

func (a BankAccount) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(a)
}

func (a BankAccount) IsEmpty() bool {
	return zero.IsZero(a)
}

type BankAccounts []BankAccount

type BankAccount struct {
	ID                int    `json:"id,omitempty"`
	AccountNumber     string `json:"accountNumber,omitempty"`
	IBAN              string `json:"iban,omitempty"`
	BIC               string `json:"bic,omitempty"`
	BankName          string `json:"bankName,omitempty"`
	Currency          string `json:"currency,omitempty"`
	DefaultForInvoice bool   `json:"defaultForInvoice,omitempty"`
	DefaultForPayment bool   `json:"defaultForPayment,omitempty"`
	Status            string `json:"status,omitempty"`
	OrderNo           int    `json:"orderNo,omitempty"`
	Version           string `json:"version,omitempty"`
}

type Meta struct {
	PageNumber  int `json:"pageNumber"`
	PageSize    int `json:"pageSize"`
	ResultCount int `json:"resultCount"`
}
