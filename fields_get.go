package bexio

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-bexio/utils"
)

func (c *Client) NewFieldsGetRequest() FieldsGetRequest {
	r := FieldsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type FieldsGetRequest struct {
	client      *Client
	queryParams *FieldsGetRequestQueryParams
	pathParams  *FieldsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody FieldsGetRequestBody
}

func (r FieldsGetRequest) NewQueryParams() *FieldsGetRequestQueryParams {
	return &FieldsGetRequestQueryParams{}
}

type FieldsGetRequestQueryParams struct {
	BusinessObject string `schema:"BusinessObject"`
	PageSize       int    `schema:"PageSize"`
}

func (p FieldsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *FieldsGetRequest) QueryParams() *FieldsGetRequestQueryParams {
	return r.queryParams
}

func (r FieldsGetRequest) NewPathParams() *FieldsGetRequestPathParams {
	return &FieldsGetRequestPathParams{}
}

type FieldsGetRequestPathParams struct {
}

func (p *FieldsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *FieldsGetRequest) PathParams() *FieldsGetRequestPathParams {
	return r.pathParams
}

func (r *FieldsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *FieldsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *FieldsGetRequest) Method() string {
	return r.method
}

func (r FieldsGetRequest) NewRequestBody() FieldsGetRequestBody {
	return FieldsGetRequestBody{}
}

type FieldsGetRequestBody struct {
}

func (r *FieldsGetRequest) RequestBody() *FieldsGetRequestBody {
	return nil
}

func (r *FieldsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *FieldsGetRequest) SetRequestBody(body FieldsGetRequestBody) {
	r.requestBody = body
}

func (r *FieldsGetRequest) NewResponseBody() *FieldsGetResponseBody {
	return &FieldsGetResponseBody{}
}

type FieldsGetResponseBody struct {
}

func (r *FieldsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("accountviewfields", r.PathParams())
	return &u
}

func (r *FieldsGetRequest) Do() (FieldsGetResponseBody, error) {
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
