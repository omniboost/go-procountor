package bexio

import "github.com/omniboost/go-bexio/utils"

func (c *Client) NewTransactGetRequest() TransactGetRequest {
	r := TransactGetRequest{
		BexioDataGetRequest: c.NewBexioDataGetRequest(),
	}
	r.BexioDataGetRequest.QueryParams().BusinessObject = "GJ1"
	return r
}

type TransactGetRequest struct {
	BexioDataGetRequest
}

func (r *TransactGetRequest) NewResponseBody() *TransactGetResponseBody {
	return &TransactGetResponseBody{}
}

type TransactGetResponseBody struct {
	Transact []Transact    `json:"TRANSACT"`
	TrnHdr   []interface{} `json:"TRN_HDR"`
}

func (r *TransactGetRequest) Do() (TransactGetResponseBody, error) {
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
