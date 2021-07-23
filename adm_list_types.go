package bexio

type AdmList struct {
	AdmArch   bool     `json:"ADM_ARCH"`   // Gearchiveerd
	AdmCode   string   `json:"ADM_CODE"`   // Admin.code
	AdmComm   string   `json:"ADM_COMM"`   // Notitie administratie
	AdmDesc   string   `json:"ADM_DESC"`   // Admin.naam
	AdmDir    string   `json:"ADM_DIR"`    // Mapnaam administratie
	AdmHold   string   `json:"ADM_HOLD"`   // Naam holding-administratie
	AdmHold2  string   `json:"ADM_HOLD2"`  // Naam holding-admin intern
	AdmModel  bool     `json:"ADM_MODEL"`  // Voorbeeldadministratie
	AdmPrev   string   `json:"ADM_PREV"`   // Administratie vorig jaar
	AdmRemote bool     `json:"ADM_REMOTE"` // Administratie op andere vestiging
	AdvOrg    bool     `json:"ADV_ORG"`    // Advies: organiseer administratie
	AllTbl    bool     `json:"ALL_TBL"`    // Administratiebestanden niet beperken
	AtsCode   string   `json:"ATS_CODE"`   // Cliëntcode
	AtsDesc   string   `json:"ATS_DESC"`   // Omschr
	AtsExtra  string   `json:"ATS_EXTRA"`  // Extra veld
	AtsLoc    string   `json:"ATS_LOC"`    // Vestigingscode
	AtsYear   string   `json:"ATS_YEAR"`   // Boekjaar
	BlckDate  DateTime `json:"BLCK_DATE"`  // Geblokkeerd op datum/tijd
	BlckUser  string   `json:"BLCK_USER"`  // Geblokkeerd door gebruiker
	BlckUsrv  string   `json:"BLCK_USRV"`  // Geblokkeerd (pictogram)
	BookDate  DateTime `json:"BOOK_DATE"`  // Datum laatste boeking
	CicoDate  DateTime `json:"CICO_DATE"`  // Meegenomen op datum/tijd
	CicoUsr   string   `json:"CICO_USR"`   // Meegenomen door gebruiker
	CngDate   DateTime `json:"CNG_DATE"`   // Wijzigingsdatum
	CngNr     int      `json:"CNG_NR"`     // Wijzignummer
	CngUsr    string   `json:"CNG_USR"`    // Medewerker laatste wijziging
	CnlPct    float64  `json:"CNL_PCT"`    // Deelnamepercentage holding
	CnlPct2   float64  `json:"CNL_PCT2"`   // Deelnamepercentage holding intern
	DbcName   string   `json:"DBC_NAME"`   // Databasenaam administratie
	InpDate   DateTime `json:"INP_DATE"`   // Invoerdatum
	InpUsr    string   `json:"INP_USR"`    // Invoermedewerker
	InstCode  string   `json:"INST_CODE"`  // Installatie
	LicCng    bool     `json:"LIC_CNG"`    // Licentie gewijzigd
	McBlock   bool     `json:"MC_BLOCK"`   // Geblokkeerd voor multicliënt telebankieren
	MsgDest   string   `json:"MSG_DEST"`   // Berichtbestemming
	MsgStgy   string   `json:"MSG_STGY"`   // Berichtklasse
	NetUse    bool     `json:"NET_USE"`    // Netwerkgebruik administratie
	OrgAdm    bool     `json:"ORG_ADM"`    // Administratie moet worden georganiseerd
	SchedBlk  bool     `json:"SCHED_BLK"`  // Geblokkeerd voor alerts
	SyncLog   bool     `json:"SYNC_LOG"`   // Stamgegevens uitwisselen

	ExtKey int `json:"EXT_KEY"`
}
