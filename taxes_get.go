package procountor

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-procountor/utils"
)

func (c *Client) NewTaxesGetRequest() TaxesGetRequest {
	r := TaxesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaxesGetRequest struct {
	client      *Client
	queryParams *TaxesGetRequestQueryParams
	pathParams  *TaxesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody TaxesGetRequestBody
}

func (r TaxesGetRequest) NewQueryParams() *TaxesGetRequestQueryParams {
	return &TaxesGetRequestQueryParams{}
}

type TaxesGetRequestQueryParams struct{}

func (p TaxesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TaxesGetRequest) QueryParams() *TaxesGetRequestQueryParams {
	return r.queryParams
}

func (r TaxesGetRequest) NewPathParams() *TaxesGetRequestPathParams {
	return &TaxesGetRequestPathParams{}
}

type TaxesGetRequestPathParams struct {
}

func (p *TaxesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TaxesGetRequest) PathParams() *TaxesGetRequestPathParams {
	return r.pathParams
}

func (r *TaxesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TaxesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *TaxesGetRequest) Method() string {
	return r.method
}

func (r TaxesGetRequest) NewRequestBody() TaxesGetRequestBody {
	return TaxesGetRequestBody{}
}

type TaxesGetRequestBody struct {
}

func (r *TaxesGetRequest) RequestBody() *TaxesGetRequestBody {
	return nil
}

func (r *TaxesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *TaxesGetRequest) SetRequestBody(body TaxesGetRequestBody) {
	r.requestBody = body
}

func (r *TaxesGetRequest) NewResponseBody() *TaxesGetResponseBody {
	return &TaxesGetResponseBody{}
}

type TaxesGetResponseBody Taxes

func (r *TaxesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("taxes", r.PathParams())
	return &u
}

func (r *TaxesGetRequest) Do() (TaxesGetResponseBody, error) {
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
