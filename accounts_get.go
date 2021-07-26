package procountor

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-procountor/utils"
)

func (c *Client) NewAccountsGetRequest() AccountsGetRequest {
	r := AccountsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountsGetRequest struct {
	client      *Client
	queryParams *AccountsGetRequestQueryParams
	pathParams  *AccountsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody AccountsGetRequestBody
}

func (r AccountsGetRequest) NewQueryParams() *AccountsGetRequestQueryParams {
	return &AccountsGetRequestQueryParams{}
}

type AccountsGetRequestQueryParams struct{}

func (p AccountsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AccountsGetRequest) QueryParams() *AccountsGetRequestQueryParams {
	return r.queryParams
}

func (r AccountsGetRequest) NewPathParams() *AccountsGetRequestPathParams {
	return &AccountsGetRequestPathParams{}
}

type AccountsGetRequestPathParams struct {
}

func (p *AccountsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccountsGetRequest) PathParams() *AccountsGetRequestPathParams {
	return r.pathParams
}

func (r *AccountsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountsGetRequest) Method() string {
	return r.method
}

func (r AccountsGetRequest) NewRequestBody() AccountsGetRequestBody {
	return AccountsGetRequestBody{}
}

type AccountsGetRequestBody struct {
}

func (r *AccountsGetRequest) RequestBody() *AccountsGetRequestBody {
	return nil
}

func (r *AccountsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *AccountsGetRequest) SetRequestBody(body AccountsGetRequestBody) {
	r.requestBody = body
}

func (r *AccountsGetRequest) NewResponseBody() *AccountsGetResponseBody {
	return &AccountsGetResponseBody{}
}

type AccountsGetResponseBody Accounts

func (r *AccountsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("accounts", r.PathParams())
	return &u
}

func (r *AccountsGetRequest) Do() (AccountsGetResponseBody, error) {
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
