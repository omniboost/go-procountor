package procountor

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-procountor/utils"
)

func (c *Client) NewBankAccountsGetRequest() BankAccountsGetRequest {
	r := BankAccountsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BankAccountsGetRequest struct {
	client      *Client
	queryParams *BankAccountsGetRequestQueryParams
	pathParams  *BankAccountsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody BankAccountsGetRequestBody
}

func (r BankAccountsGetRequest) NewQueryParams() *BankAccountsGetRequestQueryParams {
	return &BankAccountsGetRequestQueryParams{}
}

type BankAccountsGetRequestQueryParams struct{}

func (p BankAccountsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BankAccountsGetRequest) QueryParams() *BankAccountsGetRequestQueryParams {
	return r.queryParams
}

func (r BankAccountsGetRequest) NewPathParams() *BankAccountsGetRequestPathParams {
	return &BankAccountsGetRequestPathParams{}
}

type BankAccountsGetRequestPathParams struct {
}

func (p *BankAccountsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BankAccountsGetRequest) PathParams() *BankAccountsGetRequestPathParams {
	return r.pathParams
}

func (r *BankAccountsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BankAccountsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *BankAccountsGetRequest) Method() string {
	return r.method
}

func (r BankAccountsGetRequest) NewRequestBody() BankAccountsGetRequestBody {
	return BankAccountsGetRequestBody{}
}

type BankAccountsGetRequestBody struct {
}

func (r *BankAccountsGetRequest) RequestBody() *BankAccountsGetRequestBody {
	return nil
}

func (r *BankAccountsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *BankAccountsGetRequest) SetRequestBody(body BankAccountsGetRequestBody) {
	r.requestBody = body
}

func (r *BankAccountsGetRequest) NewResponseBody() *BankAccountsGetResponseBody {
	return &BankAccountsGetResponseBody{}
}

type BankAccountsGetResponseBody struct {
	Results BankAccounts `json:"results"`
	Meta    Meta         `json:"meta"`
}

func (r *BankAccountsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/bankaccounts", r.PathParams())
	return &u
}

func (r *BankAccountsGetRequest) Do() (BankAccountsGetResponseBody, error) {
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
