package models

import (
	"AdServerCollector/core/mysql"
	"encoding/json"
	"time"
)

type AdEventResourceAssoc struct {
	Id         *int       `json:"id" name:"" xorm:"id pk autoincr INT(11)"`
	EventId    *int       `json:"event_id" name:"" xorm:"event_id default NULL unique(bk) INT(11)"`
	ResourceId *int       `json:"resource_id" name:"" xorm:"resource_id default NULL unique(bk) INT(11)"`
	Updated    *time.Time `json:"updated" name:"" xorm:"updated default 'NULL' DATETIME"`
}

func (o *AdEventResourceAssoc) TableName() string {
	return "ad_event_resource_assoc"
}
func (o *AdEventResourceAssoc) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
func (o *AdEventResourceAssoc) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &o)
}
func (o *AdEventResourceAssoc) PrimaryKey() interface{} {
	return o.Id
}
func (o *AdEventResourceAssoc) NewEntity(dao mysql.BaseDao) mysql.BaseEntity {
	return mysql.NewEntity(dao, o)
}
