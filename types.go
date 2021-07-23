package bexio

import (
	"strings"
)

type Fields []Field

type Field string

func (ff Fields) MarshalSchema() string {
	fields := make([]string, len(ff))
	for i, f := range ff {
		fields[i] = string(f)
	}
	return strings.Join(fields, ",")
}

type BusinessObjects []BusinessObject

type BusinessObject struct {
	Code          string `json:"Code"`
	Description   string `json:"Description"`
	InsertAllowed bool   `json:"InsertAllowed"`
	UpdateAllowed bool   `json:"UpdateAllowed"`
	DeleteAllowed bool   `json:"DeleteAllowed"`
}

type Companies []Company

type Company struct {
	ID          string `json:"Id"`
	Description string `json:"Description"`
	Code        string `json:"Code"`
}

type VAT struct {
	AcctMemo  string   `json:"ACCT_MEMO"`  // Rekeningnummer BTW verlegd
	CalcCum   bool     `json:"CALC_CUM"`   // BTW berekenen over BTW
	CalcType  float64  `json:"CALC_TYPE"`  // BTW-rekenwijze
	CngDate   DateTime `json:"CNG_DATE"`   // Wijzigingsdatum
	CngNr     int      `json:"CNG_NR"`     // Wijzignummer
	CngUsr    string   `json:"CNG_USR"`    // Medewerker laatste wijziging
	CogsType  bool     `json:"COGS_TYPE"`  // BTW toevoegen aan kostprijs
	FrmCode   string   `json:"FRM_CODE"`   // Rubr.code
	FrmMemo   string   `json:"FRM_MEMO"`   // Gekoppelde rubriek BTW verlegd
	InpDate   DateTime `json:"INP_DATE"`   // Invoerdatum
	InpUsr    string   `json:"INP_USR"`    // Invoermedewerker
	IsAbc     bool     `json:"IS_ABC"`     // ABC-transactie
	IsReverse bool     `json:"IS_REVERSE"` // BTW verlegd
	VATAcct   string   `json:"VAT_ACCT"`   // Rekeningnr
	VATCode   string   `json:"VAT_CODE"`   // BTW-code
	VATLink   string   `json:"VAT_LINK"`   // Gekoppelde code
	VATName   string   `json:"VAT_NAME"`   // Omschrijving BTW-code
	VATPct    float64  `json:"VAT_PCT"`    // BTW-percentage
	VATType   float64  `json:"VAT_TYPE"`   // BTW-type (numeriek)

	ExtKey int `json:"EXT_KEY"`
}

type Ledger struct {
	AcctAlloc string   `json:"ACCT_ALLOC"` // Saldo leegboeken naar
	AcctBal   float64  `json:"ACCT_BAL"`   // Rekeningsaldo
	AcctCost  string   `json:"ACCT_COST"`  // Kostprijsrekening
	AcctLink  float64  `json:"ACCT_LINK"`  // Type verzamelrekening
	AcctName  string   `json:"ACCT_NAME"`  // Omschrijving rekening
	AcctNr    string   `json:"ACCT_NR"`    // Rekeningnr
	AcctProj  bool     `json:"ACCT_PROJ"`  // Projectrekening
	AcctType  float64  `json:"ACCT_TYPE"`  // Rekeningtype
	ActPy     float64  `json:"ACT_PY"`     // Saldo vorig jaar
	BalBal    float64  `json:"BAL_BAL"`    // Balans
	BalPl     float64  `json:"BAL_PL"`     // V&W
	CltBlk    bool     `json:"CLT_BLK"`    // Geblokkeerd voor cliënt
	CngDate   DateTime `json:"CNG_DATE"`   // Wijzigingsdatum
	CngNr     int      `json:"CNG_NR"`     // Wijzignummer
	CngUsr    string   `json:"CNG_USR"`    // Medewerker laatste wijziging
	ConAcct   string   `json:"CON_ACCT"`   // Consolidatierekening statutair
	ConAcct2  string   `json:"CON_ACCT2"`  // Consolidatierekening intern
	CorBal    float64  `json:"COR_BAL"`    // Saldo VJP's
	CorCamt   float64  `json:"COR_CAMT"`   // Creditbedrag correcties
	CorDamt   float64  `json:"COR_DAMT"`   // Debetbedrag correcties
	CostCode  string   `json:"COST_CODE"`  // Kostenplaatscode
	CostMand  bool     `json:"COST_MAND"`  // Kostenplaatscode verplicht
	CoAcct    bool     `json:"CO_ACCT"`    // Tussenrekening
	CrAmt     float64  `json:"CR_AMT"`     // Creditbedrag
	CrAmtmin  float64  `json:"CR_AMTMIN"`  // Negatief creditbedrag
	CurCamt   float64  `json:"CUR_CAMT"`   // Valutabedrag credit
	CurCode   string   `json:"CUR_CODE"`   // Valutacode
	CurDamt   float64  `json:"CUR_DAMT"`   // Valutabedrag debet
	DistCode  string   `json:"DIST_CODE"`  // Verdeelsleutel
	DossDesc  string   `json:"DOSS_DESC"`  // Notitie
	DossRef   string   `json:"DOSS_REF"`   // Dossierreferentie
	DrAmt     float64  `json:"DR_AMT"`     // Debetbedrag
	EfsChk    string   `json:"EFS_CHK"`    // Element van te controleren taxonomie
	EfsLedg   string   `json:"EFS_LEDG"`   // Element van administratietaxonomie
	FaAcct    bool     `json:"FA_ACCT"`    // Activumrekening
	FaSell    bool     `json:"FA_SELL"`    // Activumverk.rek
	InpDate   DateTime `json:"INP_DATE"`   // Invoerdatum
	InpUsr    string   `json:"INP_USR"`    // Invoermedewerker
	JrnlQty   bool     `json:"JRNL_QTY"`   // Aantal journaalregels controleren
	LdgStat   string   `json:"LDG_STAT"`   // Rekeningstatus
	NormStat  float64  `json:"NORM_STAT"`  // Normaal saldo
	NotDash   bool     `json:"NOT_DASH"`   // Niet opnemen in dashboardgrafieken
	PartMand  float64  `json:"PART_MAND"`  // Vennoot verplicht
	PcCode    string   `json:"PC_CODE"`    // Kostensoortcode
	PreCor    float64  `json:"PRE_COR"`    // Voor VJP's
	PreCorc   float64  `json:"PRE_CORC"`   // Voor VJP's credit
	PreCord   float64  `json:"PRE_CORD"`   // Voor VJP's debet
	ProjExp   bool     `json:"PROJ_EXP"`   // Projectkosten/-opbrengsten declarabel
	ProjObli  bool     `json:"PROJ_OBLI"`  // Projectverplichtingen
	ProAttr   string   `json:"PRO_ATTR"`   // Gekoppelde projectdimensies
	QtyMand   float64  `json:"QTY_MAND"`   // Aantal verplicht
	RecBlk    bool     `json:"REC_BLK"`    // Geblokkeerd
	RecSys    bool     `json:"REC_SYS"`    // Systeemrecord
	RevAcct   bool     `json:"REV_ACCT"`   // Valuta herwaarderen
	Rubrieken string   `json:"RUBRIEKEN"`  // Rubrieken
	SlCode1   string   `json:"SL_CODE1"`   // Subadministratie 1
	SlCode2   string   `json:"SL_CODE2"`   // Subadministratie 2
	SlCode3   string   `json:"SL_CODE3"`   // Subadministratie 3
	SplitPa   bool     `json:"SPLIT_PA"`   // Per vennoot uitsplitsen
	StatOrig  string   `json:"STAT_ORIG"`  // Oorspronkelijke rekeningstatus
	StatSell  bool     `json:"STAT_SELL"`  // Omzetrekening
	SumRpt    bool     `json:"SUM_RPT"`    // Op grootboekkrt.rapp samenvatten
	TransNr   int      `json:"TRANS_NR"`   // Aantal mutaties
	VATCode   string   `json:"VAT_CODE"`   // BTW-code
	VATMand   bool     `json:"VAT_MAND"`   // BTW-code verplicht
	WfAllow   bool     `json:"WF_ALLOW"`   // Rekening geschikt voor gebruik in IFA
	WgtMand   float64  `json:"WGT_MAND"`   // Gewicht verplicht
	WkrTrn    bool     `json:"WKR_TRN"`    // Mutatie ten laste van vrije ruimte WKR
	XbrlSpec  bool     `json:"XBRL_SPEC"`  // XBRL-specificatie

	ExtKey int `json:"EXT_KEY"`
}

type CompHdr struct {
	AiAllowpr bool     `json:"AI_ALLOWPR"` // Facturen via printservice AutoInvoice mogelijk
	AiApi     string   `json:"AI_API"`     // UUID-sleutel onderneming AutoInvoice
	CngDate   DateTime `json:"CNG_DATE"`   // Wijzigingsdatum
	CngNr     int      `json:"CNG_NR"`     // Wijzignummer
	CngUsr    string   `json:"CNG_USR"`    // Medewerker laatste wijziging
	CompDesc  string   `json:"COMP_DESC"`  // Omschrijving onderneming
	CompID    string   `json:"COMP_ID"`    // Ondernemingcode
	CMdmStrt  DateTime `json:"C_MDM_STRT"` // Startmoment synchronisatie met Visma.net
	DimFo     string   `json:"DIM_FO"`     // Dimensies voor Financial Overview
	DjisMsg   string   `json:"DJIS_MSG"`   // Meldingen dagboekberichten
	DjisProcd DateTime `json:"DJIS_PROCD"` // Laatste verwerking Dagboekberichten
	EaVar     float64  `json:"EA_VAR"`     // eAccounting-uitvoering
	EvCkey    string   `json:"EV_CKEY"`    // Klant key (eVerbinding)
	EvCsec    string   `json:"EV_CSEC"`    // Klant geheim (eVerbinding)
	EvEntity  string   `json:"EV_ENTITY"`  // Bedrijfs-ID (eVerbinding)
	EvIntreq  string   `json:"EV_INTREQ"`  // Toegangsverzoekcode
	EvReqid   string   `json:"EV_REQID"`   // E-mailadres (eVerbinding)
	ExpDate   Date     `json:"EXP_DATE"`   // Verloopdatum
	FoID      string   `json:"FO_ID"`      // ID van de administratie in Visma Online
	GocompCng int      `json:"GOCOMP_CNG"` // Wijzignummer onderneming voor AV Go
	InpDate   DateTime `json:"INP_DATE"`   // Invoerdatum
	InpUsr    string   `json:"INP_USR"`    // Invoermedewerker
	IsAutoi   bool     `json:"IS_AUTOI"`   // AutoInvoice-onderneming
	IsConApp  bool     `json:"IS_CON_APP"` // AccountView Contact (alleen-lezen)
	IsEverb   bool     `json:"IS_EVERB"`   // eVerbinding-onderneming
	IsFo      bool     `json:"IS_FO"`      // Financial Overview-onderneming
	IsGo      bool     `json:"IS_GO"`      // AccountView Go
	IsPayl    bool     `json:"IS_PAYL"`    // Betaallinkdienst-onderneming
	IsScan    bool     `json:"IS_SCAN"`    // Scanboeken-onderneming
	IsVdc     bool     `json:"IS_VDC"`     // Beschikbaar in IFA
	KvkNr     string   `json:"KVK_NR"`     // Kamer van Koophandel-nummer
	LicCheck  string   `json:"LIC_CHECK"`  // Controlenummer cliëntlicentie
	LicCode   string   `json:"LIC_CODE"`   // Modulegegevens cliëntlicentie
	LicName   string   `json:"LIC_NAME"`   // Licentienaam cliëntlicentie
	LicSer    string   `json:"LIC_SER"`    // Serienummer cliëntlicentie
	LMdmStrt  DateTime `json:"L_MDM_STRT"` // Laatste correcte synchronisatie met Visma.net
	NmbrsID   string   `json:"NMBRS_ID"`   // Nmbrs-bedrijfscd
	OdpID     int      `json:"ODP_ID"`     // Visma-ID onderneming
	PaylSec   string   `json:"PAYL_SEC"`   // Betaallink-sleutel
	ResMdm    bool     `json:"RES_MDM"`    // Stamgegevens herstellen in MDM bij volgende sync
	SbDeptcd  string   `json:"SB_DEPTCD"`  // Vestigingscode
	SbSbsnr   string   `json:"SB_SBSNR"`   // YSB-nummer
	SbSpecs   string   `json:"SB_SPECS"`   // Specificatie
	SyncEnd   DateTime `json:"SYNC_END"`   // Einde laatste synchronisatie
	SyncMsg   string   `json:"SYNC_MSG"`   // Melding laatste synchronisatie
	SyncStat  int      `json:"SYNC_STAT"`  // Status laatste synchronisatie
	SyncStrt  DateTime `json:"SYNC_STRT"`  // Start laatste synchronisatie
	VATNr     string   `json:"VAT_NR"`     // Omzetbelastingnummer

	ExtKey    int      `json:"EXT_KEY"`
	CompHdrID int      `json:"COMP_HDR_Id"`
	CompLN    []CompLn `json:"COMP_LN"`
}

type CompLn struct {
}

type Transact struct {
	AcctNr    string  `json:"ACCT_NR"`    // Rekeningnr
	ActCost   float64 `json:"ACT_COST"`   // Werkelijke kosten
	ActProc   float64 `json:"ACT_PROC"`   // Werkelijke opbrengsten
	AdmCode   string  `json:"ADM_CODE"`   // Admin.code
	Amount    float64 `json:"AMOUNT"`     // Mutatiebedrag
	BaseAmt   float64 `json:"BASE_AMT"`   // Basisbedrag BTW
	ChkID     string  `json:"CHK_ID"`     // Afletternummer
	CnID      string  `json:"CN_ID"`      // Unieke consolidatiecode
	CogsAmt   float64 `json:"COGS_AMT"`   // Kostprijs
	CogsQty   float64 `json:"COGS_QTY"`   // Kostprijs aantal
	CostCode  string  `json:"COST_CODE"`  // Kostenplaatscode
	CstAmt    float64 `json:"CST_AMT"`    // Bedr geboekt
	CurAmt    float64 `json:"CUR_AMT"`    // Valutabedrag
	CurCode   string  `json:"CUR_CODE"`   // Valutacode
	DeclAmt   float64 `json:"DECL_AMT"`   // Bedrag declarabel
	DistCode  string  `json:"DIST_CODE"`  // Verdeelsleutel
	DistQty   float64 `json:"DIST_QTY"`   // Aantal
	DjCode    string  `json:"DJ_CODE"`    // Dagboekcode
	ExchRate  float64 `json:"EXCH_RATE"`  // Valutakoers
	GrekType  string  `json:"GREK_TYPE"`  // Transactietype G-rekening
	HasDoc    bool    `json:"HAS_DOC"`    // Documentkoppeling aanwezig
	InvDesc   string  `json:"INV_DESC"`   // Factuuromschrijving transactie
	InvNr     string  `json:"INV_NR"`     // Document-/factuurnummer
	ItemDesc  string  `json:"ITEM_DESC"`  // Omschrijving transactie
	ItemQty   float64 `json:"ITEM_QTY"`   // Aantal (agr)
	LinkID    string  `json:"LINK_ID"`    // Koppeling projectjournaal
	MndPer    int     `json:"MND_PER"`    // Maanden
	ObliNr    string  `json:"OBLI_NR"`    // Verplichtingnummer
	PageNr    string  `json:"PAGE_NR"`    // Bladzijdenummer (teken)
	PageNrv   float64 `json:"PAGE_NRV"`   // Bladzijdenummer
	PartCode  string  `json:"PART_CODE"`  // Vennoot
	PartRng   string  `json:"PART_RNG"`   // Vennoot als dimensie
	PaCode    string  `json:"PA_CODE"`    // Vennoot als dimensie
	PcCode    string  `json:"PC_CODE"`    // Kostensoortcode
	PcDim     string  `json:"PC_DIM"`     // Kostensoort als dimensie
	Period    int     `json:"PERIOD"`     // Periode
	PerCalc   int     `json:"PER_CALC"`   // Periode berekend
	PerProj   float64 `json:"PER_PROJ"`   // Projectperiode
	ProjCode  string  `json:"PROJ_CODE"`  // Projectcode
	QtrPer    int     `json:"QTR_PER"`    // Kwartalen
	RecID     string  `json:"REC_ID"`     // Unieke code
	ResID     string  `json:"RES_ID"`     // Bronidentificatie
	RevAmt    float64 `json:"REV_AMT"`    // Omgedraaid bedrag
	RngCode   string  `json:"RNG_CODE"`   // Vennotenadministratie
	SellAmt   float64 `json:"SELL_AMT"`   // Bedrag te factureren
	SellQty   float64 `json:"SELL_QTY"`   // Aant te fact
	SrcID     string  `json:"SRC_ID"`     // Identificatie bron
	SrcType   float64 `json:"SRC_TYPE"`   // Brontype
	SubNr     string  `json:"SUB_NR"`     // Debiteur-/crediteurnummer
	TrnDate   Date    `json:"TRN_DATE"`   // Mutatiedatum
	TrnDesc   string  `json:"TRN_DESC"`   // Omschrijving mutatie
	TrnQty    float64 `json:"TRN_QTY"`    // Aantal (agr)
	TrnWgt    float64 `json:"TRN_WGT"`    // Gewicht
	VATAmt    float64 `json:"VAT_AMT"`    // BTW-bedrag
	VATCode   string  `json:"VAT_CODE"`   // BTW-code
	VoucherNr string  `json:"VOUCHER_NR"` // Boekstuknummer
	WkrTrn    bool    `json:"WKR_TRN"`    // Mutatie ten laste van vrije ruimte WKR
	XmlData   string  `json:"XML_DATA"`   // Gegevens in XML-formaat

	ExtKey     int `json:"EXT_KEY"`
	TransactID int `json:"TRANSACT_Id"`
}

type Daily struct {
	AcctPuw  string   `json:"ACCT_PUW"`  // Rekening betalingen onderweg
	AcctRuw  string   `json:"ACCT_RUW"`  // Rekening ontvangsten onderweg
	BnkAcct  string   `json:"BNK_ACCT"`  // Bankrekening dagboek
	BnkAcctv string   `json:"BNK_ACCTV"` // Bankrekening dagboek
	BnkCode  string   `json:"BNK_CODE"`  // Bankcode
	CheckNr  string   `json:"CHECK_NR"`  // Laatst gebruikte cheque-nummer
	ChkBal   bool     `json:"CHK_BAL"`   // Bladzijdesaldi controleren
	ChkDif   bool     `json:"CHK_DIF"`   // Verschillen controleren
	CltBlk   bool     `json:"CLT_BLK"`   // Geblokkeerd voor cliënt
	CngBlk   bool     `json:"CNG_BLK"`   // Handmatige wijzigingen geblokkeerd
	CngDate  DateTime `json:"CNG_DATE"`  // Wijzigingsdatum
	CngNr    int      `json:"CNG_NR"`    // Wijzignummer
	CngUsr   string   `json:"CNG_USR"`   // Medewerker laatste wijziging
	CntCode  string   `json:"CNT_CODE"`  // ISO-landcode
	CredLim  float64  `json:"CRED_LIM"`  // Kredietlimiet bankrekening
	CurCode  string   `json:"CUR_CODE"`  // Valutacode
	DjpBlk   bool     `json:"DJP_BLK"`   // Bladzijde automatisch blokkeren
	DjAcct   string   `json:"DJ_ACCT"`   // Gekoppelde rekening
	DjCode   string   `json:"DJ_CODE"`   // Dagboekcode
	DjComb   bool     `json:"DJ_COMB"`   // Verdicht boeken
	DjExt    bool     `json:"DJ_EXT"`    // Extern dagboek
	DjName   string   `json:"DJ_NAME"`   // Dagboeknaam
	DjOwner  float64  `json:"DJ_OWNER"`  // Dagboekeigenaar
	DjRecon  bool     `json:"DJ_RECON"`  // Bankaflettering
	DjType   float64  `json:"DJ_TYPE"`   // Type dagboek (numeriek)
	DjTypev  string   `json:"DJ_TYPEV"`  // Type dagboek
	ExchStat float64  `json:"EXCH_STAT"` // Uitwisselingsstatus
	GrekBank bool     `json:"GREK_BANK"` // Geblokkeerde bankrekening
	GrossFlg bool     `json:"GROSS_FLG"` // Bruto-invoer
	HdrDesc  int      `json:"HDR_DESC"`  // Omschrijving dagboekbladzijde
	IbanNr   string   `json:"IBAN_NR"`   // IBAN
	InpDate  DateTime `json:"INP_DATE"`  // Invoerdatum
	InpUsr   string   `json:"INP_USR"`   // Invoermedewerker
	LayCheck string   `json:"LAY_CHECK"` // Layout cheque
	LayRemit string   `json:"LAY_REMIT"` // Layout cheque-begeleiding
	LineDesc int      `json:"LINE_DESC"` // Omschrijving dagboekregels
	LstStmt  string   `json:"LST_STMT"`  // Laatst ingelezen bankafschrift
	NumAuto  bool     `json:"NUM_AUTO"`  // Document-/factuurnummers automatisch uitdelen
	NumGrp   float64  `json:"NUM_GRP"`   // Document-/factuurnummergroep
	PaCode   string   `json:"PA_CODE"`   // Vennoot
	PdDesc   int      `json:"PD_DESC"`   // Betalingsreferentie
	SendAppr bool     `json:"SEND_APPR"` // Nieuwe bladzijden verplicht in de workflow sturen
	SysVal   bool     `json:"SYS_VAL"`   // Systeemvalidaties

	ExtKey int `json:"EXT_KEY"`
}

type Cost struct {
	CngDate   DateTime `json:"CNG_DATE"`   // Wijzigingsdatum
	CngNr     int      `json:"CNG_NR"`     // Wijzignummer
	CngUsr    string   `json:"CNG_USR"`    // Medewerker laatste wijziging
	CostCode  string   `json:"COST_CODE"`  // Kostenplaatscode
	CostDesc  string   `json:"COST_DESC"`  // Omschrijving kostenplaats
	CostType  float64  `json:"COST_TYPE"`  // Type kostenplaats
	DifAcct   string   `json:"DIF_ACCT"`   // Verschillenrek
	DistCode  string   `json:"DIST_CODE"`  // Verdeelsleutel
	InpDate   DateTime `json:"INP_DATE"`   // Invoerdatum
	InpUsr    string   `json:"INP_USR"`    // Invoermedewerker
	RepLevel  float64  `json:"REP_LEVEL"`  // Niveau
	WfAuthemp string   `json:"WF_AUTHEMP"` // IFA-medewerker

	ExtKey int `json:"EXT_KEY"`
}

type Table struct {
	Definition        TableDefinition        `json:"Definition"`
	DetailDefinitions TableDetailDefinitions `json:"DetailDefinitions"`
}

type TableDefinition struct {
	Name   string                `json:"Name"`
	Fields TableDefinitionFields `json:"Fields"`
}

type TableDetailDefinitions []TableDetailDefinition

type TableDetailDefinition struct {
	Name   string                `json:"Name"`
	Fields TableDefinitionFields `json:"Fields"`
}

type TableDefinitionFields []TableDefinitionField

type TableDefinitionField struct {
	Name      string `json:"Name"`
	FieldType string `json:"FieldType"`
}

type TableData struct {
	Data       TableDataData `json:"Data"`
	DetailData DetailData    `json:"DetailData"`
}

type TableDataData struct {
	Rows Rows `json:"Rows"`
}

type DetailData []DetailDataEntry

type DetailDataEntry struct {
	Rows Rows `json:"Rows"`
}

type Rows []Row

type Row struct {
	Values []interface{} `json:"Values"`
}
