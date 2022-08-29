package procountor

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-procountor/utils"
)

func (c *Client) NewCompanyGetRequest() CompanyGetRequest {
	r := CompanyGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CompanyGetRequest struct {
	client      *Client
	queryParams *CompanyGetRequestQueryParams
	pathParams  *CompanyGetRequestPathParams
	method      string
	headers     http.Header
	requestBody CompanyGetRequestBody
}

func (r CompanyGetRequest) NewQueryParams() *CompanyGetRequestQueryParams {
	return &CompanyGetRequestQueryParams{}
}

type CompanyGetRequestQueryParams struct{}

func (p CompanyGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CompanyGetRequest) QueryParams() *CompanyGetRequestQueryParams {
	return r.queryParams
}

func (r CompanyGetRequest) NewPathParams() *CompanyGetRequestPathParams {
	return &CompanyGetRequestPathParams{}
}

type CompanyGetRequestPathParams struct {
}

func (p *CompanyGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CompanyGetRequest) PathParams() *CompanyGetRequestPathParams {
	return r.pathParams
}

func (r *CompanyGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CompanyGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CompanyGetRequest) Method() string {
	return r.method
}

func (r CompanyGetRequest) NewRequestBody() CompanyGetRequestBody {
	return CompanyGetRequestBody{}
}

type CompanyGetRequestBody struct {
}

func (r *CompanyGetRequest) RequestBody() *CompanyGetRequestBody {
	return nil
}

func (r *CompanyGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *CompanyGetRequest) SetRequestBody(body CompanyGetRequestBody) {
	r.requestBody = body
}

func (r *CompanyGetRequest) NewResponseBody() *CompanyGetResponseBody {
	return &CompanyGetResponseBody{}
}

type CompanyGetResponseBody Company

func (r *CompanyGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/company", r.PathParams())
	return &u
}

func (r *CompanyGetRequest) Do() (CompanyGetResponseBody, error) {
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
