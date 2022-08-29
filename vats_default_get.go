package procountor

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-procountor/utils"
)

func (c *Client) NewVatsDefaultGetRequest() VatsDefaultGetRequest {
	r := VatsDefaultGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VatsDefaultGetRequest struct {
	client      *Client
	queryParams *VatsDefaultGetRequestQueryParams
	pathParams  *VatsDefaultGetRequestPathParams
	method      string
	headers     http.Header
	requestBody VatsDefaultGetRequestBody
}

func (r VatsDefaultGetRequest) NewQueryParams() *VatsDefaultGetRequestQueryParams {
	return &VatsDefaultGetRequestQueryParams{}
}

type VatsDefaultGetRequestQueryParams struct{}

func (p VatsDefaultGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VatsDefaultGetRequest) QueryParams() *VatsDefaultGetRequestQueryParams {
	return r.queryParams
}

func (r VatsDefaultGetRequest) NewPathParams() *VatsDefaultGetRequestPathParams {
	return &VatsDefaultGetRequestPathParams{}
}

type VatsDefaultGetRequestPathParams struct {
}

func (p *VatsDefaultGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *VatsDefaultGetRequest) PathParams() *VatsDefaultGetRequestPathParams {
	return r.pathParams
}

func (r *VatsDefaultGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VatsDefaultGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *VatsDefaultGetRequest) Method() string {
	return r.method
}

func (r VatsDefaultGetRequest) NewRequestBody() VatsDefaultGetRequestBody {
	return VatsDefaultGetRequestBody{}
}

type VatsDefaultGetRequestBody struct {
}

func (r *VatsDefaultGetRequest) RequestBody() *VatsDefaultGetRequestBody {
	return nil
}

func (r *VatsDefaultGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *VatsDefaultGetRequest) SetRequestBody(body VatsDefaultGetRequestBody) {
	r.requestBody = body
}

func (r *VatsDefaultGetRequest) NewResponseBody() *VatsDefaultGetResponseBody {
	return &VatsDefaultGetResponseBody{}
}

type VatsDefaultGetResponseBody struct {
	VatInformation VatInformations `json:"vatInformation"`
	VatStatuses    VatStatuses     `json:"vatStatuses"`
}

func (r *VatsDefaultGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/vats/default", r.PathParams())
	return &u
}

func (r *VatsDefaultGetRequest) Do() (VatsDefaultGetResponseBody, error) {
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
