package bexio

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-bexio/utils"
)

func (c *Client) NewBexioDataGetRequest() BexioDataGetRequest {
	r := BexioDataGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BexioDataGetRequest struct {
	client      *Client
	queryParams *BexioDataGetRequestQueryParams
	pathParams  *BexioDataGetRequestPathParams
	method      string
	headers     http.Header
	requestBody BexioDataGetRequestBody
}

func (r BexioDataGetRequest) NewQueryParams() *BexioDataGetRequestQueryParams {
	return &BexioDataGetRequestQueryParams{}
}

type BexioDataGetRequestQueryParams struct {
	BusinessObject string `schema:"BusinessObject"`
	PageSize       int    `schema:"PageSize"`
	Fields         string `schema:"Fields,omitempty"`

	FilterControlSource1  string `schema:"FilterControlSource1,omitempty"`
	FilterOperator1       string `schema:"FilterOperator1,omitempty"`
	FilterValueType1      string `schema:"FilterValueType1,omitempty"`
	FilterValue1          string `schema:"FilterValue1,omitempty"`
	FilterIsListOfValues1 bool   `schema:"FilterIsListOfValues1"`

	FilterControlSource2  string `schema:"FilterControlSource2,omitempty"`
	FilterOperator2       string `schema:"FilterOperator2,omitempty"`
	FilterValueType2      string `schema:"FilterValueType2,omitempty"`
	FilterValue2          string `schema:"FilterValue2,omitempty"`
	FilterIsListOfValues2 bool   `schema:"FilterIsListOfValues2,omitempty"`

	SortFields string `schema:"SortFields,omitempty"`
	SortOrder  string `schema:"SortOrder,omitempty"`
}

func (p BexioDataGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BexioDataGetRequest) QueryParams() *BexioDataGetRequestQueryParams {
	return r.queryParams
}

func (r BexioDataGetRequest) NewPathParams() *BexioDataGetRequestPathParams {
	return &BexioDataGetRequestPathParams{}
}

type BexioDataGetRequestPathParams struct {
}

func (p *BexioDataGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BexioDataGetRequest) PathParams() *BexioDataGetRequestPathParams {
	return r.pathParams
}

func (r *BexioDataGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BexioDataGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *BexioDataGetRequest) Method() string {
	return r.method
}

func (r BexioDataGetRequest) NewRequestBody() BexioDataGetRequestBody {
	return BexioDataGetRequestBody{}
}

type BexioDataGetRequestBody struct {
}

func (r *BexioDataGetRequest) RequestBody() *BexioDataGetRequestBody {
	return nil
}

func (r *BexioDataGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *BexioDataGetRequest) SetRequestBody(body BexioDataGetRequestBody) {
	r.requestBody = body
}

func (r *BexioDataGetRequest) NewResponseBody() *BexioDataGetResponseBody {
	return &BexioDataGetResponseBody{}
}

type BexioDataGetResponseBody struct{}

func (r *BexioDataGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("accountviewdata", r.PathParams())
	return &u
}

func (r *BexioDataGetRequest) Do() (BexioDataGetResponseBody, error) {
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
