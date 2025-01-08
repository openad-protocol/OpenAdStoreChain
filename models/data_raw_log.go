package models

import (
	"AdServerCollector/core/mysql"
	"encoding/json"
	"fmt"
	"time"
)

type DataRawLog struct {
	Id          *int64  `json:"id" xorm:"id pk autoincr BIGINT(8)"`
	FirstName   *string `json:"first_name" xorm:"first_name VARCHAR(256)"`
	LastName    *string `json:"last_name" xorm:"last_name text"`
	UserName    *string `json:"user_name" xorm:"user_name text"`
	Hash        *string `json:"hash" xorm:"hash VARCHAR(64)"`
	FromType    *string `json:"from_type" xorm:"from_type comment('ajax, JsSDK, script,') VARCHAR(25)"`
	Language    *string `json:"language" xorm:"language comment('语言') VARCHAR(8)"`
	Location    *string `json:"location" xorm:"location text comment('位置')"`
	Platform    *string `json:"platform" xorm:"platform VARCHAR(12)"`
	ZoneId      *int64  `json:"zone_id" xorm:"zone_id not null comment('区域ID') BIGINT(64)"`
	PublisherId *int64  `json:"publisher_id" xorm:"publisher_id not null comment('流量主ID') BIGINT(64)"`
	EventId     *int64  `json:"event_id" xorm:"event_id not null comment('投流事件ID') BIGINT(64)"`
	Signature   *string `json:"signature" xorm:"signature VARCHAR(255)"`
	TimeStamp   *int64  `json:"time_stamp" xorm:"time_stamp BIGINT(13)"`
	TraceId     *string `json:"trace_id" xorm:"trace_id VARCHAR(64)"`
	UserId      *string `json:"user_id" xorm:"user_id comment('telegram id') VARCHAR(64)"`
	Version     *string `json:"version"  xorm:"version comment('版本') VARCHAR(64)"`
	IpAddress   *string `json:"ip_address" xorm:"ip_address comment('ip地址') VARCHAR(64)"`
	Country     *string `json:"country" xorm:"country comment('国家') VARCHAR(8)"`
	CreateAt    *int64  `json:"create_at" xorm:"create_at comment('创建时间') BIGINT(8)"`
	RequestType *string `json:"request_type" xorm:"request_type default 'NULL' comment('请求类型:getAd,logInfo,clickinfo,cb') VARCHAR(12)"`
	TraceHash   *string `json:"trace_hash" xorm:"trace_hash default 'NULL' comment('跟踪hash，指向表 ad_data_raw_traceInfo') VARCHAR(64)"`
	Cb          *string `json:"cb" xorm:"cb default 'NULL' VARCHAR(255)"`
}

func (o *DataRawLog) TableName() string {
	return "data_raw_log"
}

func (o *DataRawLog) GetSliceName(slice string) string {
	return fmt.Sprintf("data_raw_log_%s", slice)
}

func (o *DataRawLog) GetSliceDateMonthTable() string {
	t := time.Now()
	return fmt.Sprintf("data_raw_log_%d%02d", t.Year(), t.Month())
}

func (o *DataRawLog) GetSliceDateDayTable() string {
	t := time.Now()
	return fmt.Sprintf("data_raw_log_%d%02d%02d", t.Year(), t.Month(), t.Day())
}

func (o *DataRawLog) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
func (o *DataRawLog) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &o)
}
func (o *DataRawLog) PrimaryKey() interface{} {
	return o.Id
}
func (o *DataRawLog) NewEntity(dao mysql.BaseDao) mysql.BaseEntity {
	return mysql.NewEntity(dao, o)
}
