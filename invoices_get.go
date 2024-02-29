package procountor

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-procountor/utils"
)

func (c *Client) NewInvoicesGetRequest() InvoicesGetRequest {
	r := InvoicesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InvoicesGetRequest struct {
	client      *Client
	queryParams *InvoicesGetRequestQueryParams
	pathParams  *InvoicesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody InvoicesGetRequestBody
}

func (r InvoicesGetRequest) NewQueryParams() *InvoicesGetRequestQueryParams {
	return &InvoicesGetRequestQueryParams{}
}

type InvoicesGetRequestQueryParams struct {
	Status                []string `schema:"status,omitempty"`
	StartDate             string   `schema:"startDate,omitempty"`
	EndDate               string   `schema:"endDate,omitempty"`
	CreatedStartDate      string   `schema:"createdStartDate,omitempty"`
	CreatedEndDate        string   `schema:"createdEndDate,omitempty"`
	VersionStartDate      string   `schema:"versionStartDate,omitempty"`
	VersionEndDate        string   `schema:"versionEndDate,omitempty"`
	Types                 []string `schema:"types,omitempty"`
	BusinessPartnerId     int      `schema:"businessPartnerId,omitempty"`
	FactoringContractID   int      `schema:"factoringContractId,omitempty"`
	OriginalInvoiceNumber string   `schema:"originalInvoiceNumber,omitempty"`
	InvoiceNumber         string   `schema:"invoiceNumber,omitempty"`
	PreviousId            int      `schema:"previousId,omitempty"`
	OrderById             string   `schema:"orderById,omitempty"`
	OrderByDate           string   `schema:"orderByDate,omitempty"`
	OrderByCreated        string   `schema:"orderByCreated,omitempty"`
	OrderByVersion        string   `schema:"orderByVersion,omitempty"`
	Size                  int      `schema:"size,omitempty"`
	Page                  int      `schema:"page,omitempty"`
	InvoiceChannelID      []string `schema:"invoiceChannelId,omitempty"`
}

func (p InvoicesGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *InvoicesGetRequest) QueryParams() *InvoicesGetRequestQueryParams {
	return r.queryParams
}

func (r InvoicesGetRequest) NewPathParams() *InvoicesGetRequestPathParams {
	return &InvoicesGetRequestPathParams{}
}

type InvoicesGetRequestPathParams struct {
}

func (p *InvoicesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *InvoicesGetRequest) PathParams() *InvoicesGetRequestPathParams {
	return r.pathParams
}

func (r *InvoicesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InvoicesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoicesGetRequest) Method() string {
	return r.method
}

func (r InvoicesGetRequest) NewRequestBody() InvoicesGetRequestBody {
	return InvoicesGetRequestBody{}
}

type InvoicesGetRequestBody struct {
}

func (r *InvoicesGetRequest) RequestBody() *InvoicesGetRequestBody {
	return nil
}

func (r *InvoicesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *InvoicesGetRequest) SetRequestBody(body InvoicesGetRequestBody) {
	r.requestBody = body
}

func (r *InvoicesGetRequest) NewResponseBody() *InvoicesGetResponseBody {
	return &InvoicesGetResponseBody{}
}

type InvoicesGetResponseBody struct {
	Results Invoices `json:"results"`
	Meta    Meta     `json:"meta"`
}

func (r *InvoicesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/invoices", r.PathParams())
	return &u
}

func (r *InvoicesGetRequest) Do() (InvoicesGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *InvoicesGetRequest) All() (InvoicesGetResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := InvoicesGetResponseBody{}
	concat.Meta.PageNumber = resp.Meta.PageNumber
	concat.Meta.PageSize = resp.Meta.PageSize
	concat.Meta.ResultCount = resp.Meta.ResultCount

	for concat.Meta.PageSize == concat.Meta.ResultCount {
		r.QueryParams().Page = concat.Meta.PageNumber + 1
		resp, err := r.Do()
		if err != nil {
			return resp, err
		}

		concat.Meta.PageNumber = resp.Meta.PageNumber
		concat.Meta.PageSize = resp.Meta.PageSize
		concat.Meta.ResultCount = resp.Meta.ResultCount
		concat.Results = append(concat.Results, resp.Results...)
	}

	return concat, nil
}
