package bexio

type DjPage struct {
	AcctNr    string   `json:"ACCT_NR" field_type:"C"`    // Rekeningnr
	AcctPost  bool     `json:"ACCT_POST" field_type:"L"`  // Direct op grootboekrek boeken
	AcctUpd   bool     `json:"ACCT_UPD" field_type:"L"`   // Wijziging accountant
	ApprSeq   float64  `json:"APPR_SEQ" field_type:"B"`   // Approval-volgummer
	ApprState int      `json:"APPR_STATE" field_type:"I"` // Workflowstatus
	AuthDate  DateTime `json:"AUTH_DATE" field_type:"T"`  // Laatste controledatum
	AuthNr    string   `json:"AUTH_NR" field_type:"C"`    // Autorisatienr
	AuthUsr   string   `json:"AUTH_USR" field_type:"C"`   // Laatste controlegebruiker
	BlkAmt    float64  `json:"BLK_AMT" field_type:"Y"`    // Bedrag G-rekening
	BlkCur    float64  `json:"BLK_CUR" field_type:"Y"`    // Bedrag G-rekening VV
	BvAmt     float64  `json:"BV_AMT" field_type:"Y"`     // Bedrag UWV
	BvCur     float64  `json:"BV_CUR" field_type:"Y"`     // Bedrag UWV VV
	CardName  string   `json:"CARD_NAME" field_type:"C"`  // Creditcardnaam
	CardNr    string   `json:"CARD_NR" field_type:"C"`    // Creditcardnr
	CheckAmt  float64  `json:"CHECK_AMT" field_type:"Y"`  // Cheque-bedrag
	CheckPrn  bool     `json:"CHECK_PRN" field_type:"L"`  // Cheque afgedrukt
	CheckWd   string   `json:"CHECK_WD" field_type:"C"`   // Cheque-bedrag voluit
	CltUpd    bool     `json:"CLT_UPD" field_type:"L"`    // Wijziging cliënt
	CngDate   DateTime `json:"CNG_DATE" field_type:"T"`   // Wijzigingsdatum
	CngNr     int      `json:"CNG_NR" field_type:"I"`     // Wijzignummer
	CngUsr    string   `json:"CNG_USR" field_type:"C"`    // Medewerker laatste wijziging
	Comment1  string   `json:"COMMENT1" field_type:"M"`   // Notitie
	CostCode  string   `json:"COST_CODE" field_type:"C"`  // Kostenplaatscode
	CredDays  float64  `json:"CRED_DAYS" field_type:"N"`  // Betalingstermijn in dagen
	CurAmt    float64  `json:"CUR_AMT" field_type:"Y"`    // Totaal valuta
	CurCode   string   `json:"CUR_CODE" field_type:"C"`   // Valutacode
	CurVAT    float64  `json:"CUR_VAT" field_type:"Y"`    // BTW-valutatotaal
	DdAct     bool     `json:"DD_ACT" field_type:"L"`     // Automatische incasso
	DiscCode  string   `json:"DISC_CODE" field_type:"C"`  // Betalingsconditie
	DiscDays  float64  `json:"DISC_DAYS" field_type:"N"`  // Kortingstermijn in dagen
	DjlRate   float64  `json:"DJL_RATE" field_type:"B"`   // Valutakoers vreemde valutadagboek
	DjCode    string   `json:"DJ_CODE" field_type:"C"`    // Dagboekcode
	DocID     string   `json:"DOC_ID" field_type:"C"`     // Documentcode
	DocType   int      `json:"DOC_TYPE" field_type:"I"`   // Documenttype
	EaVoucher string   `json:"EA_VOUCHER" field_type:"C"` // ID van het boekstuk in eAccounting
	EbankID   string   `json:"EBANK_ID" field_type:"C"`   // Bankkoppeling-identificatie
	EbankStat int      `json:"EBANK_STAT" field_type:"I"` // Status bankkoppeling
	ExchRate  float64  `json:"EXCH_RATE" field_type:"B"`  // Valutakoers
	ExpCli    bool     `json:"EXP_CLI" field_type:"L"`    // Export naar cliënt
	ExpDate   Date     `json:"EXP_DATE" field_type:"D"`   // Geldig tot
	ExpDjp    bool     `json:"EXP_DJP" field_type:"L"`    // Geëxporteerde dagboekbladzijden
	ExpSys    bool     `json:"EXP_SYS" field_type:"L"`    // Oude waarde van sys_gen
	ExtDesc   string   `json:"EXT_DESC" field_type:"C"`   // Betalingsreferentie
	FirstWf   string   `json:"FIRST_WF" field_type:"C"`   // Eerste goedkeurder in de workflow
	GrossFlg  bool     `json:"GROSS_FLG" field_type:"L"`  // Basisbedrag incl BTW weergeven
	HdrDesc   string   `json:"HDR_DESC" field_type:"C"`   // Omschrijving bladzijde
	ImpCode   string   `json:"IMP_CODE" field_type:"C"`   // Dagboekbericht
	InpDate   DateTime `json:"INP_DATE" field_type:"T"`   // Invoerdatum
	InpUsr    string   `json:"INP_USR" field_type:"C"`    // Invoermedewerker
	InvNr     string   `json:"INV_NR" field_type:"C"`     // Document-/factuurnummer
	LbAmt     float64  `json:"LB_AMT" field_type:"Y"`     // Bedrag loonbelasting
	LbCur     float64  `json:"LB_CUR" field_type:"Y"`     // Bedrag loonbelasting VV
	MandB2b   bool     `json:"MAND_B2B" field_type:"L"`   // Incassovorm: zakelijk
	MandDate  Date     `json:"MAND_DATE" field_type:"D"`  // Datum ondertekening
	MandID    string   `json:"MAND_ID" field_type:"C"`    // Kenmerk incassomachtiging
	MandSeq   int      `json:"MAND_SEQ" field_type:"I"`   // Incassomachtiging
	NettAmt   float64  `json:"NETT_AMT" field_type:"Y"`   // Netto bladsaldo
	NoPost    bool     `json:"NO_POST" field_type:"L"`    // Bladzijde niet boeken
	OwnCngDt  DateTime `json:"OWN_CNG_DT" field_type:"T"` // Wijzigingsdatum eigenaar
	PageAmt   float64  `json:"PAGE_AMT" field_type:"Y"`   // Bladsaldo
	PageAmtt  float64  `json:"PAGE_AMTT" field_type:"Y"`  // Eindsaldo
	PageEamt  float64  `json:"PAGE_EAMT" field_type:"Y"`  // Eindsaldo
	PageNr    string   `json:"PAGE_NR" field_type:"C"`    // Bladzijdenummer (teken)
	PageNrd   string   `json:"PAGE_NRD" field_type:"C"`   // Bladzijdenummer (omschrijving)
	PageNrv   float64  `json:"PAGE_NRV" field_type:"N"`   // Bladzijdenummer
	PageVAT   float64  `json:"PAGE_VAT" field_type:"Y"`   // BTW-bedrag bladzijde
	PayinID   string   `json:"PAYIN_ID" field_type:"C"`   // Pay-in ID
	PayID     string   `json:"PAY_ID" field_type:"C"`     // Betalingscode
	PayType   string   `json:"PAY_TYPE" field_type:"C"`   // Betaalwijze
	PdAmt     float64  `json:"PD_AMT" field_type:"Y"`     // Bedrag betalingskorting
	PdPct     float64  `json:"PD_PCT" field_type:"N"`     // Percentage betalingskorting
	Period    float64  `json:"PERIOD" field_type:"N"`     // Periode
	PrevAmt   float64  `json:"PREV_AMT" field_type:"Y"`   // Beginsaldo bladzijde
	PrnCheck  bool     `json:"PRN_CHECK" field_type:"L"`  // Cheque afdrukken
	PrnRa     bool     `json:"PRN_RA" field_type:"L"`     // Cheque-begeleiding afdrukken
	ProjCode  string   `json:"PROJ_CODE" field_type:"C"`  // Projectcode
	RecBlk    bool     `json:"REC_BLK" field_type:"L"`    // Geblokkeerd
	StmtDate  Date     `json:"STMT_DATE" field_type:"D"`  // Datum afschrift
	StmtLine  bool     `json:"STMT_LINE" field_type:"L"`  // Bankafschriftregel
	StmtNr    int      `json:"STMT_NR" field_type:"I"`    // Bankafschrift
	SubNr     string   `json:"SUB_NR" field_type:"C"`     // Debiteur/Crediteur
	SysGen    bool     `json:"SYS_GEN" field_type:"L"`    // Bladzijde gegenereerd door systeem
	SysVal    bool     `json:"SYS_VAL" field_type:"L"`    // Systeemvalidaties
	TgtBal    float64  `json:"TGT_BAL" field_type:"Y"`    // Controlesaldo
	TrnDate   Date     `json:"TRN_DATE" field_type:"D"`   // Mutatiedatum
	UsrName   string   `json:"USR_NAME" field_type:"C"`   // Laatst gewijzigd door
	VexchRate float64  `json:"VEXCH_RATE" field_type:"B"` // Valutakoers (import)

	ExtKey   int `json:"EXT_KEY"`
	DjPageID int `json:"DJ_PAGE_Id"`
}

func (djPage DjPage) BusinessObject() string {
	return "DJ2"
}

func (djPage DjPage) Table() string {
	return "DJ_PAGE"
}

func (djPage DjPage) Fields() []string {
	return []string{
		"PageNr",
		"DjCode",
		"Period",
		"SubNr",
		"InvNr",
		"CurCode",
		"HdrDesc",
	}
}

func (djPage DjPage) Values() ([]interface{}, error) {
	return FieldsToValues(djPage, djPage.Fields())
}

func (djPage DjPage) ToBexioDataPostRequest(client *Client, lines []DjLine) (BexioDataPostRequest, error) {
	children := make([]BusinessObjectInterface, len(lines))
	for i, v := range lines {
		children[i] = BusinessObjectInterface(v)
	}
	return BusinessObjectToBexioDataPostRequest(client, djPage, children)
}
