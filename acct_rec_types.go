package bexio

type AcctRec struct {
	AcctNr string `json:"ACCT_NR"` // Verzamelrekening
	SubNr  string `json:"SUB_NR"`  // Debiteurnummer
	Contact

	ExtKey    int `json:"EXT_KEY"`
	ContactID int `json:"CONTACT_Id"`
}

func (acctRec AcctRec) BusinessObject() string {
	return "AR1"
}

func (acctRec AcctRec) Table() string {
	return "CONTACT"
}

func (acctRec AcctRec) Fields() []string {
	return []string{
		"SrcCode",
		"AcctName",
		"AcctNr",
		"Address1",
		"CntCode",
		"DebNr",
		"MailBus",
		"PostCode",
		"City",
		"SubNr",
		"VATNr",
	}
}

func (acctRec AcctRec) Values() ([]interface{}, error) {
	return FieldsToValues(acctRec, acctRec.Fields())
}

func (acctRec AcctRec) ToBexioDataPostRequest(client *Client) (BexioDataPostRequest, error) {
	children := make([]BusinessObjectInterface, 0)
	return BusinessObjectToBexioDataPostRequest(client, acctRec, children)
}

type Contact struct {
	AcctBal   float64  `json:"ACCT_BAL"`   // Saldo
	AcctBal2  float64  `json:"ACCT_BAL2"`  // Saldo te betalen
	AcctLink  int      `json:"ACCT_LINK"`  // Type verzamelrekening
	AcctName  string   `json:"ACCT_NAME"`  // Bedrijfsnaam
	AcctNr    string   `json:"ACCT_NR"`    // Verzamelrekening
	AcctSnd   string   `json:"ACCT_SND"`   // Klank bedrijfsnaam
	Address1  string   `json:"ADDRESS1"`   // Adres
	AddrLine1 string   `json:"ADDR_LINE1"` // Adresregel 1
	AddrLine2 string   `json:"ADDR_LINE2"` // Adresregel 2
	AdminCont string   `json:"ADMIN_CONT"` // Administratiecontactpersoon
	AdminName string   `json:"ADMIN_NAME"` // Administratiecontactpersoon
	AdminPp   string   `json:"ADMIN_PP"`   // Contactpersooncode
	AdmDear   string   `json:"ADM_DEAR"`   // Aanhef administratiecontactpersoon
	AieiMand  bool     `json:"AIEI_MAND"`  // Altijd system-to-system
	AiemMand  bool     `json:"AIEM_MAND"`  // Altijd via e-mail AutoInvoice
	AiprMand  bool     `json:"AIPR_MAND"`  // Altijd printservice AutoInvoice
	AiNowf    bool     `json:"AI_NOWF"`    // Ink.fact boeken in dagb 'Ink.f, geen workflow'
	AmtDd     float64  `json:"AMT_DD"`     // Geselecteerd bedrag te ontvangen
	AmtDdmax  float64  `json:"AMT_DDMAX"`  // Maximumincassobedrag
	AmtPay    float64  `json:"AMT_PAY"`    // Geselecteerd bedrag te betalen
	AmtRec    float64  `json:"AMT_REC"`    // Te ontvangen
	AmtRest   float64  `json:"AMT_REST"`   // Restbedrag
	AmtRest2  float64  `json:"AMT_REST2"`  // Restbedrag te betalen
	ApxList   string   `json:"APX_LIST"`   // Prijslijst art.prijzen
	AuthEmp   string   `json:"AUTH_EMP"`   // Fiatteringsmedewerker
	AutoInv   bool     `json:"AUTO_INV"`   // Facturering via AutoInvoice
	BankAcc2  string   `json:"BANK_ACC2"`  // Bankrekening 2
	BankAcct  string   `json:"BANK_ACCT"`  // Bankrekening 1
	BankAcN1  string   `json:"BANK_AC_N1"` // Bankrekening 1 relatie
	BankAcN2  string   `json:"BANK_AC_N2"` // Bankrekening 2 relatie
	BankSplit bool     `json:"BANK_SPLIT"` // Gecombineerde betaling verdelen
	BlkAcct   string   `json:"BLK_ACCT"`   // Rekeningnummer G-rekening
	BlkBnk    string   `json:"BLK_BNK"`    // Bankcode G-rekening
	BlkDate   Date     `json:"BLK_DATE"`   // Datum bewijs betalingsgedrag onderaannemer
	BlkEinv   bool     `json:"BLK_EINV"`   // G-rekeninginfo in elektronische facturen toestaan
	BlkIban   string   `json:"BLK_IBAN"`   // IBAN G-rekening
	BlkPct    float64  `json:"BLK_PCT"`    // Percentage G-rekening
	BlkSp     bool     `json:"BLK_SP"`     // Geblokkeerd voor orders
	BnkClear  bool     `json:"BNK_CLEAR"`  // Postbankrekening zuiver
	BnkCode1  string   `json:"BNK_CODE1"`  // Bankcode, bankrekening 1
	BnkCode2  string   `json:"BNK_CODE2"`  // Bankcode, bankrekening 2
	BosCode   string   `json:"BOS_CODE"`   // AccountView Server-code
	BoxCity   string   `json:"BOX_CITY"`   // Plaats of plaats postadres
	BvNr      string   `json:"BV_NR"`      // Aansluitnummer UWV
	BvPct     float64  `json:"BV_PCT"`     // Percentage UWV
	BvSub     string   `json:"BV_SUB"`     // Crediteur UWV
	ChnAdress string   `json:"CHN_ADRESS"` // Postadres dealer
	ChnCity   string   `json:"CHN_CITY"`   // Plaats dealer
	ChnCmail  string   `json:"CHN_CMAIL"`  // E-mail hoofdcontactpersoon dealer
	ChnCont   string   `json:"CHN_CONT"`   // Hoofdcontactpersoon dealer
	ChnFax    string   `json:"CHN_FAX"`    // Fax dealer
	ChnMail   string   `json:"CHN_MAIL"`   // E-mail dealer
	ChnMemb   bool     `json:"CHN_MEMB"`   // Relatie is dealer
	ChnName   string   `json:"CHN_NAME"`   // Bedrijfsnaam dealer
	ChnTel    string   `json:"CHN_TEL"`    // Telefoon dealer
	ChnType   string   `json:"CHN_TYPE"`   // Dealerkanaal
	City      string   `json:"CITY"`       // Plaats
	CngDate   DateTime `json:"CNG_DATE"`   // Datum laatste wijz
	CngNr     int      `json:"CNG_NR"`     // Wijzignummer
	CngUsr    string   `json:"CNG_USR"`    // Gewijzigd door
	CntCode   string   `json:"CNT_CODE"`   // ISO-landcode
	CocCode   string   `json:"COC_CODE"`   // KvK-nummer
	Comment1  string   `json:"COMMENT1"`   // Notitie deb/cred
	Country   string   `json:"COUNTRY"`    // Land
	CpaCode   string   `json:"CPA_CODE"`   // Accountant
	CredBal   float64  `json:"CRED_BAL"`   // Kredietruimte
	CredLim   int      `json:"CRED_LIM"`   // Kredietlimiet
	CredShow  bool     `json:"CRED_SHOW"`  // Ook als crediteur tonen
	CrAmt     float64  `json:"CR_AMT"`     // Creditbedrag
	CurCode   string   `json:"CUR_CODE"`   // Valutacode
	CurSign   string   `json:"CUR_SIGN"`   // Valutasymbool
	DdAct     bool     `json:"DD_ACT"`     // Automatische incasso toegestaan
	DebNr     string   `json:"DEB_NR"`     // Extern debiteurnr
	DebNrv    string   `json:"DEB_NRV"`    // Extern debiteur-/crediteurnummer
	DebShow   bool     `json:"DEB_SHOW"`   // Ook als debiteur tonen
	DiscCode  string   `json:"DISC_CODE"`  // Code betalingsconditie
	DiscPct   float64  `json:"DISC_PCT"`   // Factuurkorting
	DjipAcor  bool     `json:"DJIP_ACOR"`  // Boekingsvoorstellen automatisch accepteren
	DjipAMax  float64  `json:"DJIP_A_MAX"` // Maximumacceptatiebedrag
	DjipAUm   bool     `json:"DJIP_A_UM"`  // Maximumacceptatiebedrag gebruiken
	DjipGrp   int      `json:"DJIP_GRP"`   // Voorstelregels groeperen
	DjipToPi  bool     `json:"DJIP_TO_PI"` // Boekingsvoorstellen als inkoopfacturen verwerken
	DjipWnote bool     `json:"DJIP_WNOTE"` // Geen waarschuwing voor notitie in boekingsvoorstel
	DrAmt     float64  `json:"DR_AMT"`     // Debetbedrag
	Eia       string   `json:"EIA"`        // Electronic Invoice Address
	ElCi      bool     `json:"EL_CI"`      // Baliefacturen per e-mail versturen
	ElDp      bool     `json:"EL_DP"`      // Aanbetalingen per e-mail versturen
	ElPo      bool     `json:"EL_PO"`      // Bestelbonnen per e-mail versturen
	ElPs      bool     `json:"EL_PS"`      // Pakbonnen per e-mail versturen
	ElRe      bool     `json:"EL_RE"`      // Aanmaningen per e-mail versturen
	ElSbs     bool     `json:"EL_SBS"`     // Abonnementen per e-mail versturen
	ElSi      bool     `json:"EL_SI"`      // Facturen per e-mail versturen
	ElSo      bool     `json:"EL_SO"`      // Orderbevestigingen per e-mail versturen
	EmpCat    string   `json:"EMP_CAT"`    // Medewerkerklasse
	EmpNr     string   `json:"EMP_NR"`     // Verkoper/inkoper
	FaxBus    string   `json:"FAX_BUS"`    // Fax zakelijk
	FirstWf   string   `json:"FIRST_WF"`   // Eerste goedkeurder in de workflow
	Giro      string   `json:"GIRO"`       // Postbankrekening
	GrayDate  DateTime `json:"GRAY_DATE"`  // Graydon-importdatum
	GrayInfo  string   `json:"GRAY_INFO"`  // Graydon-kredietinformatie
	GrayLim   int      `json:"GRAY_LIM"`   // Graydon-kredietlimiet
	GrayScore float64  `json:"GRAY_SCORE"` // Graydon-score
	HdqCode   string   `json:"HDQ_CODE"`   // Hoofdvestiging
	HdqMemb   bool     `json:"HDQ_MEMB"`   // Nevenvestiging
	IbanNr    string   `json:"IBAN_NR"`    // IBAN 1
	IbanNr2   string   `json:"IBAN_NR2"`   // IBAN 2
	IcatCode  string   `json:"ICAT_CODE"`  // IS categorie goederenstroom
	IcntCode  string   `json:"ICNT_CODE"`  // IS landcode
	ImmWf     bool     `json:"IMM_WF"`     // Facturen direct naar workflow sturen
	InpDate   DateTime `json:"INP_DATE"`   // Invoerdatum
	InpUsr    string   `json:"INP_USR"`    // Invoermedewerker
	InvAddres string   `json:"INV_ADDRES"` // Adres in landafhankelijke indeling (memo)
	InvAnnex  bool     `json:"INV_ANNEX"`  // Factuur en bijlagen samenvoegen
	IshpCode  string   `json:"ISHP_CODE"`  // IS vervoerswijze
	IsxfCode  string   `json:"ISXF_CODE"`  // IS statistische factorcode
	IsysCode  string   `json:"ISYS_CODE"`  // IS stelsel
	IsComp    bool     `json:"IS_COMP"`    // Relatie is mijn concurrent
	IsThird   bool     `json:"IS_THIRD"`   // Relatie is derde partij
	ItypeCodv string   `json:"ITYPE_CODV"` // IS transactiecode (lang)
	ItypCode  string   `json:"ITYP_CODE"`  // IS transactiecode
	KValDate  Date     `json:"K_VAL_DATE"` // Verificatiedatum KvK-nummer
	LbNr      string   `json:"LB_NR"`      // Loonbelastingnummer
	LbPct     float64  `json:"LB_PCT"`     // Percentage loonbelasting
	LbSub     string   `json:"LB_SUB"`     // Crediteur loonbelasting
	LngCode   string   `json:"LNG_CODE"`   // Taalcode
	LnkAcct   string   `json:"LNK_ACCT"`   // Standaardrekening
	MailBus   string   `json:"MAIL_BUS"`   // E-mail zakelijk
	MailOpt   int      `json:"MAIL_OPT"`   // Mailing
	MainCont  string   `json:"MAIN_CONT"`  // Contactpersoon
	MainPp    string   `json:"MAIN_PP"`    // Hoofdcontactpersoon
	MandB2b   bool     `json:"MAND_B2B"`   // Incassovorm: zakelijk
	MandDate  Date     `json:"MAND_DATE"`  // Datum ondertekening
	MandID    string   `json:"MAND_ID"`    // Kenmerk incassomachtiging
	McDear    string   `json:"MC_DEAR"`    // Aanhef hoofdcontactpersoon
	McMailb   string   `json:"MC_MAILB"`   // E-mail hoofdcontactpersoon
	MrnList   string   `json:"MRN_LIST"`   // Prijslijst partnerkorting
	NoAppr    bool     `json:"NO_APPR"`    // Uitsluiten van workflow
	OppCursol string   `json:"OPP_CURSOL"` // Pakketkeuze
	OppSrc    string   `json:"OPP_SRC"`    // Bron
	PayName   string   `json:"PAY_NAME"`   // Korte bedrijfsnaam
	PctList   string   `json:"PCT_LIST"`   // Prijslijst kortingspercentages
	PdPfx     string   `json:"PD_PFX"`     // Betalingskenmerk
	PgBlk     bool     `json:"PG_BLK"`     // Geblokkeerd voor autom betaling
	PolPrio   float64  `json:"POL_PRIO"`   // Uitleverprioriteit
	PostCode  string   `json:"POST_CODE"`  // Postcode
	PoBox     string   `json:"PO_BOX"`     // Postadres
	PoCity    string   `json:"PO_CITY"`    // Plaats postadres
	PoCode    string   `json:"PO_CODE"`    // Postcd postadres
	ProjCred  bool     `json:"PROJ_CRED"`  // Projectcrediteur
	QdCode    string   `json:"QD_CODE"`    // Prijslijstcode
	RecBlk    bool     `json:"REC_BLK"`    // Geblokkeerd
	RemAddr   string   `json:"REM_ADDR"`   // Adres crediteurenadministratie in landafh indeling
	RemBlk    bool     `json:"REM_BLK"`    // Geblokkeerd voor aanmaningen
	RevCat    string   `json:"REV_CAT"`    // Omzetcategorie
	RplChn    string   `json:"RPL_CHN"`    // Dealer van relatie
	RplGrp    string   `json:"RPL_GRP"`    // Relatiegroep
	RplInv    string   `json:"RPL_INV"`    // Factuurdebiteur/crediteur
	RplType   string   `json:"RPL_TYPE"`   // Relatietype
	RplUser   bool     `json:"RPL_USER"`   // Relatie is klant
	SaleAmt   float64  `json:"SALE_AMT"`   // Omzet huidig jaar
	SalePy    float64  `json:"SALE_PY"`    // Omzet vorig jaar
	ScanID    string   `json:"SCAN_ID"`    // Scanb-relatienr
	SgCode    string   `json:"SG_CODE"`    // Versie
	ShipComp  bool     `json:"SHIP_COMP"`  // Geen deelleveringen toegestaan
	ShipMth   string   `json:"SHIP_MTH"`   // Vervoersmethode
	ShowMempo bool     `json:"SHOW_MEMPO"` // Notitie tonen bij toevoegen inkooporder
	ShowMemso bool     `json:"SHOW_MEMSO"` // Notitie tonen bij toevoegen verkooporder
	SicCode   string   `json:"SIC_CODE"`   // Branchecode
	SoAnnex   bool     `json:"SO_ANNEX"`   // Orderbevestiging en bijlagen samenvoegen
	SpPostadr int      `json:"SP_POSTADR"` // Adres aflevering
	SrcCode   string   `json:"SRC_CODE"`   // Zoekcode bedrijf
	SrcPc     string   `json:"SRC_PC"`     // Postcode (snelzoeken)
	StExtra   string   `json:"ST_EXTRA"`   // Adres-toevoeging
	StNr      string   `json:"ST_NR"`      // Huisnummer
	StProp    string   `json:"ST_PROP"`    // Gebouw
	StRegion  string   `json:"ST_REGION"`  // Adres-regio
	StStreet  string   `json:"ST_STREET"`  // Straatnaam
	SubNr     string   `json:"SUB_NR"`     // Debiteur-/crediteurnummer
	SyncState int      `json:"SYNC_STATE"` // Synchronisatiestatus
	TelBus    string   `json:"TEL_BUS"`    // Telefoon zakelijk
	TelMob    string   `json:"TEL_MOB"`    // Telefoon mobiel
	TransNr   int      `json:"TRANS_NR"`   // Aantal mutaties
	VATCode   string   `json:"VAT_CODE"`   // BTW-code
	VATDate   Date     `json:"VAT_DATE"`   // Verificatiedatum
	VATNr     string   `json:"VAT_NR"`     // BTW-nummer
	VisAddres string   `json:"VIS_ADDRES"` // Bezoekadres in landafhankelijke indeling (memo)
	WfAuthemp string   `json:"WF_AUTHEMP"` // IFA-medewerker
	WfDirect  bool     `json:"WF_DIRECT"`  // Direct boeken op kostenrekeningen (IFA)
	WwwUrl    string   `json:"WWW_URL"`    // Website

	ExtKey int `json:"EXT_KEY"`
}
