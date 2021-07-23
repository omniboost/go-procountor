package bexio

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-bexio/utils"
)

func (c *Client) NewVATPeriodsGetRequest() VATPeriodsGetRequest {
	r := VATPeriodsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VATPeriodsGetRequest struct {
	client      *Client
	queryParams *VATPeriodsGetRequestQueryParams
	pathParams  *VATPeriodsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody VATPeriodsGetRequestBody
}

func (r VATPeriodsGetRequest) NewQueryParams() *VATPeriodsGetRequestQueryParams {
	return &VATPeriodsGetRequestQueryParams{}
}

type VATPeriodsGetRequestQueryParams struct{}

func (p VATPeriodsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VATPeriodsGetRequest) QueryParams() *VATPeriodsGetRequestQueryParams {
	return r.queryParams
}

func (r VATPeriodsGetRequest) NewPathParams() *VATPeriodsGetRequestPathParams {
	return &VATPeriodsGetRequestPathParams{}
}

type VATPeriodsGetRequestPathParams struct {
}

func (p *VATPeriodsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *VATPeriodsGetRequest) PathParams() *VATPeriodsGetRequestPathParams {
	return r.pathParams
}

func (r *VATPeriodsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VATPeriodsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *VATPeriodsGetRequest) Method() string {
	return r.method
}

func (r VATPeriodsGetRequest) NewRequestBody() VATPeriodsGetRequestBody {
	return VATPeriodsGetRequestBody{}
}

type VATPeriodsGetRequestBody struct {
}

func (r *VATPeriodsGetRequest) RequestBody() *VATPeriodsGetRequestBody {
	return nil
}

func (r *VATPeriodsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *VATPeriodsGetRequest) SetRequestBody(body VATPeriodsGetRequestBody) {
	r.requestBody = body
}

func (r *VATPeriodsGetRequest) NewResponseBody() *VATPeriodsGetResponseBody {
	return &VATPeriodsGetResponseBody{}
}

type VATPeriodsGetResponseBody VATPeriods

func (r *VATPeriodsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("accounting/vat_periods", r.PathParams())
	return &u
}

func (r *VATPeriodsGetRequest) Do() (VATPeriodsGetResponseBody, error) {
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
