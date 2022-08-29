package procountor

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-procountor/utils"
)

func (c *Client) NewDimensionsGetRequest() DimensionsGetRequest {
	r := DimensionsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DimensionsGetRequest struct {
	client      *Client
	queryParams *DimensionsGetRequestQueryParams
	pathParams  *DimensionsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody DimensionsGetRequestBody
}

func (r DimensionsGetRequest) NewQueryParams() *DimensionsGetRequestQueryParams {
	return &DimensionsGetRequestQueryParams{}
}

type DimensionsGetRequestQueryParams struct{}

func (p DimensionsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DimensionsGetRequest) QueryParams() *DimensionsGetRequestQueryParams {
	return r.queryParams
}

func (r DimensionsGetRequest) NewPathParams() *DimensionsGetRequestPathParams {
	return &DimensionsGetRequestPathParams{}
}

type DimensionsGetRequestPathParams struct {
}

func (p *DimensionsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DimensionsGetRequest) PathParams() *DimensionsGetRequestPathParams {
	return r.pathParams
}

func (r *DimensionsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DimensionsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *DimensionsGetRequest) Method() string {
	return r.method
}

func (r DimensionsGetRequest) NewRequestBody() DimensionsGetRequestBody {
	return DimensionsGetRequestBody{}
}

type DimensionsGetRequestBody struct {
}

func (r *DimensionsGetRequest) RequestBody() *DimensionsGetRequestBody {
	return nil
}

func (r *DimensionsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *DimensionsGetRequest) SetRequestBody(body DimensionsGetRequestBody) {
	r.requestBody = body
}

func (r *DimensionsGetRequest) NewResponseBody() *DimensionsGetResponseBody {
	return &DimensionsGetResponseBody{}
}

type DimensionsGetResponseBody Dimensions

func (r *DimensionsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/dimensions", r.PathParams())
	return &u
}

func (r *DimensionsGetRequest) Do() (DimensionsGetResponseBody, error) {
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
