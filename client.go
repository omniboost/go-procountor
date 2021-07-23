package bexio

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"reflect"
	"strconv"
	"strings"
	"text/template"

	"github.com/fatih/structtag"
	"github.com/pkg/errors"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-bexio/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "www.bexio",
		Path:   "/api/v3",
	}
)

// NewClient returns a new Exact Globe Client client
func NewClient(httpClient *http.Client, companyID string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{}

	client.SetHTTPClient(httpClient)
	client.SetCompanyID(companyID)
	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with Exact Globe Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	companyID string

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	// Optional function called after every successful request made to the DO Clients
	beforeRequestDo    BeforeRequestDoCallback
	onRequestCompleted RequestCompletionCallback
}

type BeforeRequestDoCallback func(*http.Client, *http.Request, interface{})

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) SetHTTPClient(client *http.Client) {
	c.http = client
}

func (c Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c Client) CompanyID() string {
	return c.companyID
}

func (c *Client) SetCompanyID(companyID string) {
	c.companyID = companyID
}

func (c Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) SetBeforeRequestDo(fun BeforeRequestDoCallback) {
	c.beforeRequestDo = fun
}

func (c *Client) GetEndpointURL(p string, pathParams PathParams) url.URL {
	clientURL := c.BaseURL()

	parsed, err := url.Parse(p)
	if err != nil {
		log.Fatal(err)
	}
	q := clientURL.Query()
	for k, vv := range parsed.Query() {
		for _, v := range vv {
			q.Add(k, v)
		}
	}
	clientURL.RawQuery = q.Encode()

	clientURL.Path = path.Join(clientURL.Path, parsed.Path)

	tmpl, err := template.New("path").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	// params["administration_id"] = c.Administration()
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = buf.String()
	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, req Request) (*http.Request, error) {
	// convert body struct to json
	buf := new(bytes.Buffer)
	if req.RequestBodyInterface() != nil {
		err := json.NewEncoder(buf).Encode(req.RequestBodyInterface())
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	r, err := http.NewRequest(req.Method(), req.URL().String(), buf)
	if err != nil {
		return nil, err
	}

	// values := url.Values{}
	// err = utils.AddURLValuesToRequest(values, req, true)
	// if err != nil {
	// 	return nil, err
	// }

	// optionally pass along context
	if ctx != nil {
		r = r.WithContext(ctx)
	}

	// set other headers
	r.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	r.Header.Add("Accept", c.MediaType())
	r.Header.Add("User-Agent", c.UserAgent())
	r.Header.Add("X-Company", c.CompanyID())

	return r, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, body interface{}) (*http.Response, error) {
	if c.beforeRequestDo != nil {
		c.beforeRequestDo(c.http, req, body)
	}

	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	if body == nil {
		return httpResp, err
	}

	if httpResp.ContentLength == 0 {
		return httpResp, nil
	}

	errResp := &ErrorResponse{Response: httpResp}
	err = c.Unmarshal(httpResp.Body, body, errResp)
	if err != nil {
		return httpResp, err
	}

	if errResp.Error() != "" {
		return httpResp, errResp
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	errs := []error{}
	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)
		if c.disallowUnknownFields {
			dec.DisallowUnknownFields()
		}

		err := dec.Decode(v)
		if err != nil && err != io.EOF {
			errs = append(errs, err)
		}

	}

	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = fmt.Sprint(e)
		}
		return errors.New(strings.Join(msgs, ", "))
	}

	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	err = checkContentType(r)
	if err != nil {
		return errors.WithStack(err)
	}

	if r.ContentLength == 0 {
		return errors.New("response body is empty")
	}

	// convert json to struct
	if len(data) != 0 {
		err = json.Unmarshal(data, &errorResponse)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	if errorResponse.Error() != "" {
		return errorResponse
	}

	return nil
}

// {
//   "ErrorType": "AccountViewError",
//   "ErrorNumbers": null,
//   "ErrorMessage": ""
// }

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	Type    string      `json:"ErrorType"`
	Numbers interface{} `json:"ErrorNumbers"`
	Message string      `json:"ErrorMessage"`
}

func (r *ErrorResponse) Error() string {
	if r.Type == "" && r.Message == "" {
		return ""
	}

	return fmt.Sprintf("%s: %s", r.Type, r.Message)
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

type BusinessObjectInterface interface {
	BusinessObject() string
	Table() string
	Fields() []string
	Values() ([]interface{}, error)
}

func BusinessObjectToBexioDataPostRequest(client *Client, object BusinessObjectInterface, children []BusinessObjectInterface) (BexioDataPostRequest, error) {
	var err error
	req := client.NewBexioDataPostRequest()
	body := req.RequestBody()

	body.BookDate = "2021-07-02T10:39:05.276Z"
	body.BusinessObject = object.BusinessObject()
	body.Table.Definition, err = BusinessObjectToTableDefinition(object)
	if err != nil {
		return req, errors.WithStack(err)
	}
	body.TableData.Data, err = BusinessObjectToTableDataData(object)
	if err != nil {
		return req, errors.WithStack(err)
	}

	if len(children) > 0 {
		detailDefinition, err := BusinessObjectToDetailDefinition(children[0])
		if err != nil {
			return req, errors.WithStack(err)
		}
		body.Table.DetailDefinitions = append(body.Table.DetailDefinitions, detailDefinition)

		body.TableData.DetailData = make(DetailData, 1)
		for _, c := range children {
			rowID := strconv.Itoa(len(body.TableData.DetailData[0].Rows) + 1)
			headerID := "1"
			data, err := BusinessObjectToDetailData(c, rowID, headerID)
			if err != nil {
				return req, errors.WithStack(err)
			}

			body.TableData.DetailData[0].Rows = append(body.TableData.DetailData[0].Rows, data[0].Rows...)
		}
	}

	return req, nil
}

func BusinessObjectToTableDefinition(object BusinessObjectInterface) (TableDefinition, error) {
	definition := TableDefinition{}
	definition.Name = object.Table()

	ff := object.Fields()
	dff, err := FieldsToDefinitionFields(object, ff)
	if err != nil {
		return definition, err
	}

	// add RowId table definition
	dff = append(dff, TableDefinitionField{
		Name:      "RowId",
		FieldType: "C",
	})

	definition.Fields = dff
	return definition, nil
}

func BusinessObjectToTableDataData(object BusinessObjectInterface) (TableDataData, error) {
	tdd := TableDataData{}

	values, err := object.Values()
	if err != nil {
		return tdd, errors.WithStack(err)
	}

	// add RowId value
	values = append(values, []interface{}{"1"}...)

	tdd.Rows = Rows{{values}}
	return tdd, nil
}

func BusinessObjectToDetailDefinition(object BusinessObjectInterface) (TableDetailDefinition, error) {
	definition := TableDetailDefinition{}
	definition.Name = object.Table()

	ff := object.Fields()
	dff, err := FieldsToDefinitionFields(object, ff)
	if err != nil {
		return definition, err
	}

	// add RowId table definition
	dff = append(dff, TableDefinitionField{
		Name:      "RowId",
		FieldType: "C",
	})

	// add HeaderId table definition
	dff = append(dff, TableDefinitionField{
		Name:      "HeaderId",
		FieldType: "C",
	})

	definition.Fields = dff
	return definition, nil
}

func BusinessObjectToDetailData(object BusinessObjectInterface, rowID, headerID string) (DetailData, error) {
	dd := DetailData{
		DetailDataEntry{
			Rows: Rows{},
		},
	}

	values, err := object.Values()
	if err != nil {
		return dd, errors.WithStack(err)
	}

	// add RowId & HeaderId value
	values = append(values, []interface{}{rowID, headerID}...)

	dd[0].Rows = Rows{{values}}
	return dd, nil
}

func FieldsToDefinitionFields(object BusinessObjectInterface, fields []string) (TableDefinitionFields, error) {
	tdf := make(TableDefinitionFields, len(fields))

	for i, f := range fields {
		field, ok := reflect.TypeOf(object).FieldByName(f)
		if !ok {
			return tdf, errors.Errorf("%s is not an existing field", f)
		}

		tags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			return tdf, err
		}

		jsonTag, err := tags.Get("json")
		if err != nil {
			return tdf, err
		}

		// fieldTypeTag, err := tags.Get("field_type")
		// if err != nil {
		// 	return tdf, err
		// }

		value := reflect.ValueOf(object).FieldByName(f).Interface()
		fieldType := ""
		switch t := value.(type) {
		case int:
			fieldType = "I"
		case string:
			fieldType = "C"
		case float64:
			fieldType = "N"
		case Date:
			fieldType = "T"
		default:
			return tdf, errors.Errorf("Don't know how to map type %s", t)
		}

		tdf[i] = TableDefinitionField{Name: jsonTag.Name, FieldType: fieldType}
	}

	return tdf, nil
}

func FieldsToValues(object BusinessObjectInterface, fields []string) ([]interface{}, error) {
	values := make([]interface{}, len(fields))

	for i, f := range fields {
		values[i] = reflect.ValueOf(object).FieldByName(f).Interface()
	}

	return values, nil
}
