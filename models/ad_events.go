package models

import (
	"AdServerCollector/core/mysql"
	"encoding/json"
	"time"
)

type AdEvents struct {
	Id                 *int       `json:"id" name:"" xorm:"id pk autoincr MEDIUMINT(9)"`
	Name               *string    `json:"name" name:"" xorm:"name not null default '''' VARCHAR(255)"`
	Pricing            *int       `json:"pricing" name:"定价方式，默认为 'CPM'" xorm:"pricing not null default 0 comment('定价方式，默认为 'CPM'') INT(11)"`
	AdvertiserId       *int       `json:"advertiser_id" name:"广告主id" xorm:"advertiser_id not null default 0 comment('广告主id') index MEDIUMINT(9)"`
	Restriction        *string    `json:"restriction" name:"约束（标的tag意向）" xorm:"restriction default 'NULL' comment('约束（标的tag意向）') VARCHAR(255)"`
	ActivateTime       *time.Time `json:"activate_time" name:"激活时间" xorm:"activate_time default 'NULL' comment('激活时间') DATETIME"`
	ExpireTime         *time.Time `json:"expire_time" name:"过期时间" xorm:"expire_time default 'NULL' comment('过期时间') DATETIME"`
	Price              *float64   `json:"price" name:"单价" xorm:"price default NULL comment('单价') DECIMAL(10,2)"`
	Budget             *float64   `json:"budget" name:"预算" xorm:"budget default NULL comment('预算') DECIMAL(10,2)"`
	CapDailyPerUser    *int       `json:"cap_daily_per_user" name:"每用户每天展示上限" xorm:"cap_daily_per_user default NULL comment('每用户每天展示上限') INT(10)"`
	Views              *int       `json:"views" name:"展示次数，默认为 '-1'" xorm:"views default -1 comment('展示次数，默认为 '-1'') INT(11)"`
	Clicks             *int       `json:"clicks" name:"点击次数，默认为 '-1'" xorm:"clicks default -1 comment('点击次数，默认为 '-1'') INT(11)"`
	Conversions        *int       `json:"conversions" name:"转化次数，默认为 '-1'" xorm:"conversions default -1 comment('转化次数，默认为 '-1'') INT(11)"`
	Priority           *int       `json:"priority" name:"优先级，默认为 '0'" xorm:"priority not null default 0 comment('优先级，默认为 '0'') INT(11)"`
	Weight             *int       `json:"weight" name:"权重，默认为 '1'" xorm:"weight not null default 1 comment('权重，默认为 '1'') TINYINT(4)"`
	Anonymous          *string    `json:"anonymous" name:"是否匿名，'t' 表示是，'f' 表示否，默认为 'f'" xorm:"anonymous not null default ''f'' comment('是否匿名，'t' 表示是，'f' 表示否，默认为 'f'') ENUM('f','t')"`
	Companion          *int       `json:"companion" name:"伴随广告，默认为 '0'" xorm:"companion default 0 comment('伴随广告，默认为 '0'') SMALLINT(6)"`
	Comments           *string    `json:"comments" name:"" xorm:"comments default 'NULL' TEXT"`
	Updated            *time.Time `json:"updated" name:"" xorm:"updated not null DATETIME"`
	Status             *int       `json:"status" name:"状态" xorm:"status not null default 0 comment('状态') INT(11)"`
	HostedViews        *int       `json:"hosted_views" name:"托管展示次数" xorm:"hosted_views not null default 0 comment('托管展示次数') INT(11)"`
	HostedClicks       *int       `json:"hosted_clicks" name:"托管点击次数" xorm:"hosted_clicks not null default 0 comment('托管点击次数') INT(11)"`
	ViewWindow         *int       `json:"view_window" name:"展示窗口期" xorm:"view_window not null default 0 comment('展示窗口期') MEDIUMINT(9)"`
	ClickWindow        *int       `json:"click_window" name:"点击窗口期" xorm:"click_window not null default 0 comment('点击窗口期') MEDIUMINT(9)"`
	Ecpm               *float64   `json:"ecpm" name:"每千次展示的有效成本" xorm:"ecpm default NULL comment('每千次展示的有效成本') DECIMAL(10,4)"`
	MinImpressions     *int       `json:"min_impressions" name:"最小展示次数" xorm:"min_impressions not null default 0 comment('最小展示次数') INT(11)"`
	EcpmEnabled        *int       `json:"ecpm_enabled" name:"是否启用 eCPM" xorm:"ecpm_enabled not null default 0 comment('是否启用 eCPM') TINYINT(4)"`
	ShowCappedNoCookie *int       `json:"show_capped_no_cookie" name:"是否在无 Cookie 时显示封顶广告" xorm:"show_capped_no_cookie not null default 0 comment('是否在无 Cookie 时显示封顶广告') TINYINT(4)"`
}

func (o *AdEvents) TableName() string {
	return "ad_events"
}
func (o *AdEvents) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
func (o *AdEvents) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &o)
}
func (o *AdEvents) PrimaryKey() interface{} {
	return o.Id
}
func (o *AdEvents) NewEntity(dao mysql.BaseDao) mysql.BaseEntity {
	return mysql.NewEntity(dao, o)
}
