package bexio

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-bexio/utils"
)

func (c *Client) NewBusinessObjectsGetRequest() BusinessObjectsGetRequest {
	r := BusinessObjectsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BusinessObjectsGetRequest struct {
	client      *Client
	queryParams *BusinessObjectsGetRequestQueryParams
	pathParams  *BusinessObjectsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody BusinessObjectsGetRequestBody
}

func (r BusinessObjectsGetRequest) NewQueryParams() *BusinessObjectsGetRequestQueryParams {
	return &BusinessObjectsGetRequestQueryParams{}
}

type BusinessObjectsGetRequestQueryParams struct {
	// BusinessObject string `schema:"BusinessObject"`
	// PageSize       int    `schema:"PageSize"`
}

func (p BusinessObjectsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BusinessObjectsGetRequest) QueryParams() *BusinessObjectsGetRequestQueryParams {
	return r.queryParams
}

func (r BusinessObjectsGetRequest) NewPathParams() *BusinessObjectsGetRequestPathParams {
	return &BusinessObjectsGetRequestPathParams{}
}

type BusinessObjectsGetRequestPathParams struct {
}

func (p *BusinessObjectsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BusinessObjectsGetRequest) PathParams() *BusinessObjectsGetRequestPathParams {
	return r.pathParams
}

func (r *BusinessObjectsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BusinessObjectsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *BusinessObjectsGetRequest) Method() string {
	return r.method
}

func (r BusinessObjectsGetRequest) NewRequestBody() BusinessObjectsGetRequestBody {
	return BusinessObjectsGetRequestBody{}
}

type BusinessObjectsGetRequestBody struct {
}

func (r *BusinessObjectsGetRequest) RequestBody() *BusinessObjectsGetRequestBody {
	return nil
}

func (r *BusinessObjectsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *BusinessObjectsGetRequest) SetRequestBody(body BusinessObjectsGetRequestBody) {
	r.requestBody = body
}

func (r *BusinessObjectsGetRequest) NewResponseBody() *BusinessObjectsGetResponseBody {
	return &BusinessObjectsGetResponseBody{}
}

type BusinessObjectsGetResponseBody BusinessObjects

func (r *BusinessObjectsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("accountviewsystembusinessobjects", r.PathParams())
	return &u
}

func (r *BusinessObjectsGetRequest) Do() (BusinessObjectsGetResponseBody, error) {
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
