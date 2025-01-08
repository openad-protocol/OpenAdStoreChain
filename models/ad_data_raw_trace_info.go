package models

import (
	"AdServerCollector/core/mysql"
	"encoding/json"
	"fmt"
	"time"
)

type AdDataRawTraceInfo struct {
	Id            *int64  `json:"id" xorm:"id pk autoincr BIGINT(8)"`
	TraceId       *string `json:"trace_id"  xorm:"trace_id default 'NULL' comment('对应前端trace_id字段') VARCHAR(64)"`
	EventId       *string `json:"event_id" xorm:"event_id default 'NULL' comment('投流事件ID') VARCHAR(64)"`
	LoginfoHash   *string `json:"loginfo_hash"  orm:"loginfo_hash default 'NULL' comment('对应前端 hash字段') VARCHAR(64)"`
	CbHash        *string `json:"cb_hash"  xorm:"cb_hash default 'NULL' comment('对应前端cb字段') VARCHAR(64)"`
	CreateAt      *int64  `json:"create_at"  xorm:"create_at default current_timestamp(6) comment('记录时间') BIGINT(6)"`
	ClickinfoHash *string `json:"clickinfo_hash"  xorm:"clickinfo_hash default 'NULL' comment('对应前端sig字段') VARCHAR(64)"`
}

func (o *AdDataRawTraceInfo) TableName() string {
	return "ad_data_raw_trace_info"
}

func (o *AdDataRawTraceInfo) GetSliceDateMonthTable() string {
	t := time.Now()
	return fmt.Sprintf("ad_data_raw_trace_info_%d%02d", t.Year(), t.Month())
}

func (o *AdDataRawTraceInfo) GetSliceDateDayTable() string {
	t := time.Now()
	return fmt.Sprintf("ad_data_raw_trace_info_%d%02d%02d", t.Year(), t.Month(), t.Day())
}

func (o *AdDataRawTraceInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
func (o *AdDataRawTraceInfo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &o)
}
func (o *AdDataRawTraceInfo) PrimaryKey() interface{} {
	return o.Id
}
func (o *AdDataRawTraceInfo) NewEntity(dao mysql.BaseDao) mysql.BaseEntity {
	return mysql.NewEntity(dao, o)
}
