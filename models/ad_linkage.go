package models

import (
	"AdServerCollector/core/mysql"
	"encoding/json"
	"time"
)

type AdLinkage struct {
	Id             *int       `json:"id" name:"" xorm:"id pk autoincr MEDIUMINT(9)"`
	ZoneId         *int       `json:"zone_id" name:"版位id" xorm:"zone_id default NULL comment('版位id') unique(bk) MEDIUMINT(9)"`
	EventId        *int       `json:"event_id" name:"投流事件id" xorm:"event_id default NULL comment('投流事件id') unique(bk) INT(11)"`
	ResourceId     *int       `json:"resource_id" name:"素材id" xorm:"resource_id default NULL comment('素材id') unique(bk) MEDIUMINT(9)"`
	PublisherId    *int       `json:"publisher_id" name:"流量主id" xorm:"publisher_id default NULL comment('流量主id') MEDIUMINT(9)"`
	AdvertiserId   *int       `json:"advertiser_id" name:"广告主id" xorm:"advertiser_id default NULL comment('广告主id') MEDIUMINT(9)"`
	Priority       *float64   `json:"priority" name:"优先级" xorm:"priority default 0 comment('优先级') DOUBLE"`
	LinkType       *int       `json:"link_type" name:"关联类型？" xorm:"link_type not null default 1 comment('关联类型？') SMALLINT(6)"`
	PriorityFactor *float64   `json:"priority_factor" name:"" xorm:"priority_factor default 0 DOUBLE"`
	ToBeDelivered  *bool      `json:"to_be_delivered" name:"" xorm:"to_be_delivered not null default 1 TINYINT(1)"`
	Updated        *time.Time `json:"updated" name:"" xorm:"updated not null DATETIME"`
	Deleted        *time.Time `json:"deleted" name:"" xorm:"deleted DATETIME"`
}

func (o *AdLinkage) TableName() string {
	return "ad_linkage"
}
func (o *AdLinkage) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
func (o *AdLinkage) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &o)
}
func (o *AdLinkage) PrimaryKey() interface{} {
	return o.Id
}
func (o *AdLinkage) NewEntity(dao mysql.BaseDao) mysql.BaseEntity {
	return mysql.NewEntity(dao, o)
}
