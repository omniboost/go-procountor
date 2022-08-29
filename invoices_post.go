package procountor

import (
	"net/http"
	"net/url"

	"github.com/cydev/zero"
	"github.com/omniboost/go-procountor/omitempty"
	"github.com/omniboost/go-procountor/utils"
)

func (c *Client) NewInvoicesPostRequest() InvoicesPostRequest {
	r := InvoicesPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InvoicesPostRequest struct {
	client      *Client
	queryParams *InvoicesPostRequestQueryParams
	pathParams  *InvoicesPostRequestPathParams
	method      string
	headers     http.Header
	requestBody InvoicesPostRequestBody
}

func (r InvoicesPostRequest) NewQueryParams() *InvoicesPostRequestQueryParams {
	return &InvoicesPostRequestQueryParams{}
}

type InvoicesPostRequestQueryParams struct{}

func (p InvoicesPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *InvoicesPostRequest) QueryParams() *InvoicesPostRequestQueryParams {
	return r.queryParams
}

func (r InvoicesPostRequest) NewPathParams() *InvoicesPostRequestPathParams {
	return &InvoicesPostRequestPathParams{}
}

type InvoicesPostRequestPathParams struct {
}

func (p *InvoicesPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *InvoicesPostRequest) PathParams() *InvoicesPostRequestPathParams {
	return r.pathParams
}

func (r *InvoicesPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InvoicesPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoicesPostRequest) Method() string {
	return r.method
}

func (r InvoicesPostRequest) NewRequestBody() InvoicesPostRequestBody {
	return InvoicesPostRequestBody{}
}

type InvoicesPostRequestBody Invoice

func (rb InvoicesPostRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(rb)
}

func (rb InvoicesPostRequestBody) IsEmpty() bool {
	return zero.IsZero(rb)
}

func (r *InvoicesPostRequest) RequestBody() *InvoicesPostRequestBody {
	return &r.requestBody
}

func (r *InvoicesPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *InvoicesPostRequest) SetRequestBody(body InvoicesPostRequestBody) {
	r.requestBody = body
}

func (r *InvoicesPostRequest) NewResponseBody() *InvoicesPostResponseBody {
	return &InvoicesPostResponseBody{}
}

type InvoicesPostResponseBody Invoice

func (r *InvoicesPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/invoices", r.PathParams())
	return &u
}

func (r *InvoicesPostRequest) Do() (InvoicesPostResponseBody, error) {
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
