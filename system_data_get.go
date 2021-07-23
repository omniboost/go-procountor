package bexio

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-bexio/utils"
)

func (c *Client) NewSystemDataGetRequest() SystemDataGetRequest {
	r := SystemDataGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SystemDataGetRequest struct {
	client      *Client
	queryParams *SystemDataGetRequestQueryParams
	pathParams  *SystemDataGetRequestPathParams
	method      string
	headers     http.Header
	requestBody SystemDataGetRequestBody
}

func (r SystemDataGetRequest) NewQueryParams() *SystemDataGetRequestQueryParams {
	return &SystemDataGetRequestQueryParams{}
}

type SystemDataGetRequestQueryParams struct {
	BusinessObject       string `schema:"businessobject"`
	PageSize             int    `schema:"pagesize"`
	FilterControlSource1 string `schema:"FilterControlSource1,omitempty"`
	FilterOperator1      string `schema:"FilterOperator1,omitempty"`
	FilterValue1         string `schema:"FilterValue1,omitempty"`
	FilterValueType1     string `schema:"FilterValueType1,omitempty"`
	SortFields           string `schema:"SortFields,omitempty"`
	SortOrder            string `schema:"SortOrder,omitempty"`
	Fields               Fields `schema:"Fields"`
}

func (p SystemDataGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(Fields{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SystemDataGetRequest) QueryParams() *SystemDataGetRequestQueryParams {
	return r.queryParams
}

func (r SystemDataGetRequest) NewPathParams() *SystemDataGetRequestPathParams {
	return &SystemDataGetRequestPathParams{}
}

type SystemDataGetRequestPathParams struct {
}

func (p *SystemDataGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SystemDataGetRequest) PathParams() *SystemDataGetRequestPathParams {
	return r.pathParams
}

func (r *SystemDataGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SystemDataGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *SystemDataGetRequest) Method() string {
	return r.method
}

func (r SystemDataGetRequest) NewRequestBody() SystemDataGetRequestBody {
	return SystemDataGetRequestBody{}
}

type SystemDataGetRequestBody struct {
}

func (r *SystemDataGetRequest) RequestBody() *SystemDataGetRequestBody {
	return nil
}

func (r *SystemDataGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *SystemDataGetRequest) SetRequestBody(body SystemDataGetRequestBody) {
	r.requestBody = body
}

func (r *SystemDataGetRequest) NewResponseBody() *SystemDataGetResponseBody {
	return &SystemDataGetResponseBody{}
}

type SystemDataGetResponseBody struct {
	DCTFILE []struct {
		DBFNAME   string  `json:"DBF_NAME"`   // Tabelnaam
		DBFDESC   string  `json:"DBF_DESC"`   // Omschrijving tabel
		DBFDESCT  string  `json:"DBF_DESCT"`  // Omschrijving tabel (vertaald)
		DBFTYPE   int     `json:"DBF_TYPE"`   // Tabeltype
		OBJGRP    string  `json:"OBJ_GRP"`    // Explorer-subgroepen
		DBFTYPED  string  `json:"DBF_TYPED"`  // Omschrijving tabeltype
		DBFVIRT   bool    `json:"DBF_VIRT"`   // Virtuele tabel
		DBFBASE   string  `json:"DBF_BASE"`   // Verwijzing naar basistabel
		ADDBASE   bool    `json:"ADD_BASE"`   // Echte velden en indexen aan basistabel toevoegen
		DBFCATEG  int     `json:"DBF_CATEG"`  // Categorie
		DBFCATEGD string  `json:"DBF_CATEGD"` // Omschrijving categorie
		FILEEXTRA string  `json:"FILE_EXTRA"` // Extra tabelinformatie
		RECDESC   string  `json:"REC_DESC"`   // Recordomschrijving
		OBJCODE   string  `json:"OBJ_CODE"`   // Object Code
		STDVIEW   bool    `json:"STD_VIEW"`   // Standaardobject gebruiken
		TBLCHILD  string  `json:"TBL_CHILD"`  // Regeltabel
		SORTEXPN  string  `json:"SORT_EXPN"`  // Sorteeruitdrukking
		WHENEXPN  string  `json:"WHEN_EXPN"`  // When-uitdrukking
		OBJWHEN   string  `json:"OBJ_WHEN"`   // Object zichtbaar als
		DESCEXPN  string  `json:"DESC_EXPN"`  // Uitdrukking omschrijving sleutelveld
		RECCUST   bool    `json:"REC_CUST"`   // Aangepast
		RECBLK    bool    `json:"REC_BLK"`    // Geblokkeerd
		CNGNR     int     `json:"CNG_NR"`     // Wijzignummer
		OBJLIST   bool    `json:"OBJ_LIST"`   // Object in opvraaglijst
		OBJIMP    bool    `json:"OBJ_IMP"`    // Object bruikbaar bij import
		OBJTASK   bool    `json:"OBJ_TASK"`   // Object geschikt voor taakgerichte interface
		TBLACTIVE bool    `json:"TBL_ACTIVE"` // Tabel actief
		OBJACTIVE bool    `json:"OBJ_ACTIVE"` // Object actief
		OBJTYPE   float64 `json:"OBJ_TYPE"`   // Objecttype
		NOUPDATE  bool    `json:"NO_UPDATE"`  // Wijzigen niet mogelijk
		NOINSERT  bool    `json:"NO_INSERT"`  // Toevoegen niet mogelijk
		NODELETE  bool    `json:"NO_DELETE"`  // Verwijderen niet mogelijk
		NOCOPY    bool    `json:"NO_COPY"`    // Kopiëren niet mogelijk
		DBFEXTRA  string  `json:"DBF_EXTRA"`  // Notitie
		LOGINS    bool    `json:"LOG_INS"`    // Toevoegingen registreren
		LOGDEL    bool    `json:"LOG_DEL"`    // Verwijderingen registreren
		LOGUPD    bool    `json:"LOG_UPD"`    // Wijzigingen registreren
		SETUPDATA bool    `json:"SETUP_DATA"` // In nieuw jaar overnemen
		TXTID     string  `json:"TXT_ID"`     // Tekstidentificatie
		TBLHDR    string  `json:"TBL_HDR"`    // Headertabel
		INPDATE   string  `json:"INP_DATE"`   // Invoerdatum
		CNGDATE   string  `json:"CNG_DATE"`   // Wijzigingsdatum
		CNGUSR    string  `json:"CNG_USR"`    // Medewerker laatste wijziging
		INPUSR    string  `json:"INP_USR"`    // Invoermedewerker
		NOXMLUPD  bool    `json:"NO_XML_UPD"` // Mutaties niet via Web- API
	} `json:"DCT_FILE"`
	DCTFLD []struct {
		FLD_DESC  string  `json:"FLD_DESC"`   // Source Field Description
		FIELDNAME string  `json:"FIELD_NAME"` // Veldnaam
		DBFNAME   string  `json:"DBF_NAME"`   // Tabelnaam
		FIELDLEN  float64 `json:"FIELD_LEN"`  // Breedte
		FIELDTYPE string  `json:"FIELD_TYPE"` // Veldtype
		INPMAND   bool    `json:"INP_MAND"`   // Invoer verplicht
		FLDFMT    string  `json:"FLD_FMT"`    // Formaat
		FLDMASK   string  `json:"FLD_MASK"`   // Invoermasker
		COMBLK    bool    `json:"COM_BLK"`    // Blokkeren voor COM
		CNGDATE   string  `json:"CNG_DATE"`   // Wijzigingsdatum
		CNGNR     int     `json:"CNG_NR"`     // Wijzignummer
		CNGUSR    string  `json:"CNG_USR"`    // Medewerker laatste wijziging
		INPDATE   string  `json:"INP_DATE"`   // Invoerdatum
		PKYTBL    string  `json:"PKY_TBL"`    // Naam hoofdtabel
		RECCUST   bool    `json:"REC_CUST"`   // Aangepast
		DCTFLDID  int     `json:"DCT_FLD_Id"`
	} `json:"DCT_FLD"`
	DCTTAG []interface{} `json:"DCT_TAG"`
}

func (r *SystemDataGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("accountviewsystemdata", r.PathParams())
	return &u
}

func (r *SystemDataGetRequest) Do() (SystemDataGetResponseBody, error) {
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
