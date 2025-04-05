package message

type BaseMsg struct {
	MsgId    string      `json:"msg_id"`
	MainCode int         `json:"main_code"`
	SubCode  int         `json:"sub_code"`
	Payload  interface{} `json:"payload"`
}

type TraceInfo struct {
	TraceId       *string `json:"trace_id"`
	EventId       *string `json:"event_id"`
	LogInfoHash   *string `json:"loginfo_hash"`
	ClickInfoHash *string `json:"clickinfo_hash"`
	CbHash        *string `json:"cb_hash"`
}

type AdDataRaw struct {
	// telegram
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	UserName  *string `json:"userName"`
	// line
	ChanneId    *string `json:"channeId"`
	LiffId      *string `json:"liffId"`
	DisplayName *string `json:"displayName"`

	Hash          *string `json:"hash"`
	FromType      *string `json:"fromType"`
	Language      *string `json:"language"`
	Location      *string `json:"location"`
	Platform      *string `json:"platform"`
	ZoneId        *string `json:"zoneId"`
	EventId       *string `jons:"eventId"`
	PublisherId   *string `json:"publisherId"`
	Signature     *string `json:"signature"`
	TimeStamp     *string `json:"timeStamp"`
	TraceId       *string `json:"traceId"`
	UserId        *string `json:"userId"`
	Version       *string `json:"version"`
	IpAddress     *string `json:"ip_address"`
	Country       *string `json:"country"`
	RequestType   *string `json:"requestType"`
	TraceHash     *string `json:"traceHash"`
	WalletType    *string `json:"walletType" `
	WalletAddress *string `json:"walletAddress"`
	IsPremium     *string `json:"isPremium" `
}
