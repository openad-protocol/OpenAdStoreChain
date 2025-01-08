package models

import (
	"AdServerCollector/core/mysql"
	"encoding/json"
	"time"
)

type StatHourEvent struct {
	Id         *int       `json:"id" name:"" xorm:"id pk autoincr INT(11)"`
	Hour       *time.Time `json:"hour" name:"这一小时" xorm:"hour not null pk comment('这一小时') DATETIME"`
	EventId    *int       `json:"event_id" name:"投流事件id" xorm:"event_id default NULL comment('投流事件id') INT(11)"`
	ResourceId *int       `json:"resource_id" name:"物料id" xorm:"resource_id not null pk comment('物料id') MEDIUMINT(9)"`
	ZoneId     *int       `json:"zone_id" name:"版位id" xorm:"zone_id not null pk comment('版位id') MEDIUMINT(9)"`
	Action     *string    `json:"action" name:"操作：VIEW,CLICK" xorm:"action default 'NULL' comment('操作：VIEW,CLICK') VARCHAR(255)"`
	Count      *int64     `json:"count" name:"" xorm:"count not null default 0 INT(11)"`
}

func (o *StatHourEvent) TableName() string {
	return "stat_hour_event"
}
func (o *StatHourEvent) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
func (o *StatHourEvent) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &o)
}
func (o *StatHourEvent) PrimaryKey() interface{} {
	return o.Id
}
func (o *StatHourEvent) NewEntity(dao mysql.BaseDao) mysql.BaseEntity {
	return mysql.NewEntity(dao, o)
}
