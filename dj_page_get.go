package bexio

import "github.com/omniboost/go-bexio/utils"

func (c *Client) NewDjPageGetRequest() DjPageGetRequest {
	r := DjPageGetRequest{
		BexioDataGetRequest: c.NewBexioDataGetRequest(),
	}
	r.BexioDataGetRequest.QueryParams().BusinessObject = "DJ2P"
	return r
}

type DjPageGetRequest struct {
	BexioDataGetRequest
}

func (r *DjPageGetRequest) NewResponseBody() *DjPageGetResponseBody {
	return &DjPageGetResponseBody{}
}

type DjPageGetResponseBody struct {
	DjPage  []DjPage      `json:"DJ_PAGE"`
	DjLine  []DjLine      `json:"DJ_LINE"`
	UsrLink []interface{} `json:"USR_LINK"`
	PageWf  []interface{} `json:"PAGE_WF"`
}

func (r *DjPageGetRequest) Do() (DjPageGetResponseBody, error) {
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
