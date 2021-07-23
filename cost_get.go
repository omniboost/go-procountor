package bexio

import "github.com/omniboost/go-bexio/utils"

func (c *Client) NewCostGetRequest() CostGetRequest {
	r := CostGetRequest{
		BexioDataGetRequest: c.NewBexioDataGetRequest(),
	}
	r.BexioDataGetRequest.QueryParams().BusinessObject = "CC1"
	return r
}

type CostGetRequest struct {
	BexioDataGetRequest
}

func (r *CostGetRequest) NewResponseBody() *CostGetResponseBody {
	return &CostGetResponseBody{}
}

type CostGetResponseBody struct {
	Cost []Cost `json:"COST"`
}

func (r *CostGetRequest) Do() (CostGetResponseBody, error) {
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
