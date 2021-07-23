package bexio

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-bexio/utils"
)

func (c *Client) NewCompaniesGetRequest() CompaniesGetRequest {
	r := CompaniesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CompaniesGetRequest struct {
	client      *Client
	queryParams *CompaniesGetRequestQueryParams
	pathParams  *CompaniesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody CompaniesGetRequestBody
}

func (r CompaniesGetRequest) NewQueryParams() *CompaniesGetRequestQueryParams {
	return &CompaniesGetRequestQueryParams{}
}

type CompaniesGetRequestQueryParams struct{}

func (p CompaniesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CompaniesGetRequest) QueryParams() *CompaniesGetRequestQueryParams {
	return r.queryParams
}

func (r CompaniesGetRequest) NewPathParams() *CompaniesGetRequestPathParams {
	return &CompaniesGetRequestPathParams{}
}

type CompaniesGetRequestPathParams struct {
}

func (p *CompaniesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CompaniesGetRequest) PathParams() *CompaniesGetRequestPathParams {
	return r.pathParams
}

func (r *CompaniesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CompaniesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CompaniesGetRequest) Method() string {
	return r.method
}

func (r CompaniesGetRequest) NewRequestBody() CompaniesGetRequestBody {
	return CompaniesGetRequestBody{}
}

type CompaniesGetRequestBody struct {
}

func (r *CompaniesGetRequest) RequestBody() *CompaniesGetRequestBody {
	return nil
}

func (r *CompaniesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *CompaniesGetRequest) SetRequestBody(body CompaniesGetRequestBody) {
	r.requestBody = body
}

func (r *CompaniesGetRequest) NewResponseBody() *CompaniesGetResponseBody {
	return &CompaniesGetResponseBody{}
}

type CompaniesGetResponseBody Companies

func (r *CompaniesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("companies", r.PathParams())
	return &u
}

func (r *CompaniesGetRequest) Do() (CompaniesGetResponseBody, error) {
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
