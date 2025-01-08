package models

import (
	"AdServerCollector/core/mysql"
	"encoding/json"
	"time"
)

type AdZones struct {
	Id                 *int       `json:"id" name:"" xorm:"id pk autoincr index(rv_zones_zonenameid) MEDIUMINT(9)"`
	Name               *string    `json:"name" name:"" xorm:"name not null default '''' index(rv_zones_zonenameid) VARCHAR(245)"`
	Type               *int       `json:"type" name:"Banner,FullscreenVideo,MetaTask" xorm:"type not null default 0 comment('Banner,FullscreenVideo,MetaTask') SMALLINT(6)"`
	PublisherId        *int       `json:"publisher_id" name:"流量主id" xorm:"publisher_id default NULL comment('流量主id') index MEDIUMINT(9)"`
	Restriction        *string    `json:"restriction" name:"约束（标的tag意向）" xorm:"restriction default 'NULL' comment('约束（标的tag意向）') VARCHAR(255)"`
	Description        *string    `json:"description" name:"" xorm:"description not null default '''' VARCHAR(255)"`
	Delivery           *int       `json:"delivery" name:"广告投放方式?" xorm:"delivery not null default 0 comment('广告投放方式?') SMALLINT(6)"`
	Category           *string    `json:"category" name:"广告类别" xorm:"category not null comment('广告类别') TEXT"`
	Width              *int       `json:"width" name:"广告版位的宽度" xorm:"width not null default 0 comment('广告版位的宽度') SMALLINT(6)"`
	Height             *int       `json:"height" name:"广告版位的高度" xorm:"height not null default 0 comment('广告版位的高度') SMALLINT(6)"`
	AdSelection        *string    `json:"ad_selection" name:"广告选择方式" xorm:"ad_selection not null comment('广告选择方式') TEXT"`
	Chain              *string    `json:"chain" name:"广告链" xorm:"chain not null comment('广告链') TEXT"`
	Prepend            *string    `json:"prepend" name:"前置内容" xorm:"prepend not null comment('前置内容') TEXT"`
	Append             *string    `json:"append" name:"附加内容" xorm:"append not null comment('附加内容') TEXT"`
	Appendtype         *int       `json:"appendtype" name:"附加内容类型" xorm:"appendtype not null default 0 comment('附加内容类型') TINYINT(4)"`
	Forceappend        *string    `json:"forceappend" name:"是否强制附加" xorm:"forceappend default ''f'' comment('是否强制附加') ENUM('f','t')"`
	Comments           *string    `json:"comments" name:"" xorm:"comments default 'NULL' TEXT"`
	Updated            *time.Time `json:"updated" name:"" xorm:"updated not null DATETIME"`
	Rate               *float64   `json:"rate" name:"费率" xorm:"rate default NULL comment('费率') DECIMAL(19,2)"`
	Pricing            *int       `json:"pricing" name:"计费方式" xorm:"pricing not null default 0 comment('计费方式') TINYINT(4)"`
	ShowCappedNoCookie *int       `json:"show_capped_no_cookie" name:"是否在无 Cookie 时显示封顶广告，默认为 '0'" xorm:"show_capped_no_cookie not null default 0 comment('是否在无 Cookie 时显示封顶广告，默认为 '0'') TINYINT(4)"`
	Status             *int       `json:"status" name:"状态 0 暂停 1 开始" xorm:"status not null default 0 comment('状态') TINYINT(4)"`
}

func (o *AdZones) TableName() string {
	return "ad_zones"
}
func (o *AdZones) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
func (o *AdZones) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &o)
}
func (o *AdZones) PrimaryKey() interface{} {
	return o.Id
}
func (o *AdZones) NewEntity(dao mysql.BaseDao) mysql.BaseEntity {
	return mysql.NewEntity(dao, o)
}
