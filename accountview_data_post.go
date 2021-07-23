package bexio

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/omniboost/go-bexio/utils"
)

func (c *Client) NewBexioDataPostRequest() BexioDataPostRequest {
	r := BexioDataPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BexioDataPostRequest struct {
	client      *Client
	queryParams *BexioDataPostRequestQueryParams
	pathParams  *BexioDataPostRequestPathParams
	method      string
	headers     http.Header
	requestBody BexioDataPostRequestBody
}

func (r BexioDataPostRequest) NewQueryParams() *BexioDataPostRequestQueryParams {
	return &BexioDataPostRequestQueryParams{}
}

type BexioDataPostRequestQueryParams struct{}

func (p BexioDataPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BexioDataPostRequest) QueryParams() *BexioDataPostRequestQueryParams {
	return r.queryParams
}

func (r BexioDataPostRequest) NewPathParams() *BexioDataPostRequestPathParams {
	return &BexioDataPostRequestPathParams{}
}

type BexioDataPostRequestPathParams struct {
}

func (p *BexioDataPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BexioDataPostRequest) PathParams() *BexioDataPostRequestPathParams {
	return r.pathParams
}

func (r *BexioDataPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BexioDataPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *BexioDataPostRequest) Method() string {
	return r.method
}

func (r BexioDataPostRequest) NewRequestBody() BexioDataPostRequestBody {
	return BexioDataPostRequestBody{}
}

type BexioDataPostRequestBody struct {
	BookDate       string    `json:"BookDate"`
	BusinessObject string    `json:"BusinessObject"`
	Table          Table     `json:"Table"`
	TableData      TableData `json:"TableData"`
}

func (r *BexioDataPostRequest) RequestBody() *BexioDataPostRequestBody {
	return &r.requestBody
}

func (r *BexioDataPostRequest) RequestBodyInterface() interface{} {
	return r.RequestBody()
}

func (r *BexioDataPostRequest) SetRequestBody(body BexioDataPostRequestBody) {
	r.requestBody = body
}

func (r *BexioDataPostRequest) NewResponseBody() *BexioDataPostResponseBody {
	return &BexioDataPostResponseBody{}
}

type BexioDataPostResponseBody struct {
	json.RawMessage
}

func (r *BexioDataPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("accountviewdata", r.PathParams())
	return &u
}

func (r *BexioDataPostRequest) Do() (BexioDataPostResponseBody, error) {
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
