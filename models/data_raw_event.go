package models

import (
	"AdServerCollector/core/mysql"
	"encoding/json"
	"time"
)

type DataRawEvent struct {
	Id         *int       `json:"id" name:"" xorm:"id pk INT(11)"`
	UserId     *string    `json:"user_id" name:"telegram的id" xorm:"user_id default 'NULL' comment('telegram的id') index VARCHAR(32)"`
	EventId    *int       `json:"event_id" name:"投流事件id" xorm:"event_id not null comment('投流事件id') index INT(10)"`
	ZoneId     *int       `json:"zone_id" name:"版位id" xorm:"zone_id not null comment('版位id') index INT(10)"`
	ResourceId *int       `json:"resource_id" name:"物料id" xorm:"resource_id not null comment('物料id') INT(10)"`
	DateTime   *time.Time `json:"date_time" name:"精确时间，到毫秒" xorm:"date_time not null comment('精确时间，到毫秒') index DATETIME"`
	Language   *string    `json:"language" name:"" xorm:"language default 'NULL' VARCHAR(32)"`
	IpAddress  *string    `json:"ip_address" name:"" xorm:"ip_address default 'NULL' VARCHAR(16)"`
	HostName   *string    `json:"host_name" name:"" xorm:"host_name default 'NULL' VARCHAR(255)"`
	Country    *string    `json:"country" name:"" xorm:"country default 'NULL' CHAR(2)"`
	Https      *bool      `json:"https" name:"" xorm:"https default NULL TINYINT(1)"`
	Domain     *string    `json:"domain" name:"" xorm:"domain default 'NULL' VARCHAR(255)"`
	Page       *string    `json:"page" name:"" xorm:"page default 'NULL' VARCHAR(255)"`
	Query      *string    `json:"query" name:"" xorm:"query default 'NULL' VARCHAR(255)"`
	Referer    *string    `json:"referer" name:"" xorm:"referer default 'NULL' VARCHAR(255)"`
	SearchTerm *string    `json:"search_term" name:"" xorm:"search_term default 'NULL' VARCHAR(255)"`
	UserAgent  *string    `json:"user_agent" name:"" xorm:"user_agent default 'NULL' VARCHAR(255)"`
	Os         *string    `json:"os" name:"" xorm:"os default 'NULL' VARCHAR(32)"`
	Browser    *string    `json:"browser" name:"" xorm:"browser default 'NULL' VARCHAR(32)"`
	MaxHttps   *bool      `json:"max_https" name:"" xorm:"max_https default NULL TINYINT(1)"`
}

func (o *DataRawEvent) TableName() string {
	return "data_raw_event"
}
func (o *DataRawEvent) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
func (o *DataRawEvent) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &o)
}
func (o *DataRawEvent) PrimaryKey() interface{} {
	return o.Id
}
func (o *DataRawEvent) NewEntity(dao mysql.BaseDao) mysql.BaseEntity {
	return mysql.NewEntity(dao, o)
}
