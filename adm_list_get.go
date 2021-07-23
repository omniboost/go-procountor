package bexio

import "github.com/omniboost/go-bexio/utils"

func (c *Client) NewAdmListGetRequest() AdmListGetRequest {
	r := AdmListGetRequest{
		SystemDataGetRequest: c.NewSystemDataGetRequest(),
	}
	r.SystemDataGetRequest.QueryParams().BusinessObject = "ADL"
	return r
}

type AdmListGetRequest struct {
	SystemDataGetRequest
}

func (r *AdmListGetRequest) NewResponseBody() *AdmListGetResponseBody {
	return &AdmListGetResponseBody{}
}

type AdmListGetResponseBody struct {
	AdmList []AdmList `json:"ADM_LIST"`
}

func (r *AdmListGetRequest) Do() (AdmListGetResponseBody, error) {
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
