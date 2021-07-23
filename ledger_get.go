package bexio

import "github.com/omniboost/go-bexio/utils"

func (c *Client) NewLedgerGetRequest() LedgerGetRequest {
	r := LedgerGetRequest{
		BexioDataGetRequest: c.NewBexioDataGetRequest(),
	}
	r.BexioDataGetRequest.QueryParams().BusinessObject = "GL1"
	return r
}

type LedgerGetRequest struct {
	BexioDataGetRequest
}

func (r *LedgerGetRequest) NewResponseBody() *LedgerGetResponseBody {
	return &LedgerGetResponseBody{}
}

type LedgerGetResponseBody struct {
	Ledger []Ledger `json:"LEDGER"`
}

func (r *LedgerGetRequest) Do() (LedgerGetResponseBody, error) {
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
