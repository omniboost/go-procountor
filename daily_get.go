package bexio

import "github.com/omniboost/go-bexio/utils"

func (c *Client) NewDailyGetRequest() DailyGetRequest {
	r := DailyGetRequest{
		BexioDataGetRequest: c.NewBexioDataGetRequest(),
	}
	r.BexioDataGetRequest.QueryParams().BusinessObject = "DJ1"
	return r
}

type DailyGetRequest struct {
	BexioDataGetRequest
}

func (r *DailyGetRequest) NewResponseBody() *DailyGetResponseBody {
	return &DailyGetResponseBody{}
}

type DailyGetResponseBody struct {
	Daily []Daily `json:"DAILY"`
}

func (r *DailyGetRequest) Do() (DailyGetResponseBody, error) {
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
