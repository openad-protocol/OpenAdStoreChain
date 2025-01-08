package models

import (
	"AdServerCollector/core/mysql"
	"encoding/json"
	"time"
)

type AdAdvertisers struct {
	Id               *int       `json:"id" name:"" xorm:"id pk autoincr MEDIUMINT(9)"`
	Name             *string    `json:"name" name:"" xorm:"name not null default '''' VARCHAR(255)"`
	AccountId        *int       `json:"account_id" name:"" xorm:"account_id default NULL unique MEDIUMINT(9)"`
	Contact          *string    `json:"contact" name:"" xorm:"contact default 'NULL' VARCHAR(255)"`
	Email            *string    `json:"email" name:"" xorm:"email not null default '''' VARCHAR(64)"`
	Report           *string    `json:"report" name:"是否生成报告，'t' 表示是，'f' 表示否，默认为 'f'" xorm:"report not null default ''f'' comment('是否生成报告，'t' 表示是，'f' 表示否，默认为 'f'') ENUM('f','t')"`
	Reportinterval   *int       `json:"reportinterval" name:"报告生成的间隔天数，默认为 '7'" xorm:"reportinterval not null default 7 comment('报告生成的间隔天数，默认为 '7'') MEDIUMINT(9)"`
	Reportlastdate   *time.Time `json:"reportlastdate" name:"上次生成报告的日期，默认为 '0000-00-00'" xorm:"reportlastdate not null comment('上次生成报告的日期，默认为 '0000-00-00'') DATE"`
	Reportdeactivate *string    `json:"reportdeactivate" name:"是否停用报告，'t' 表示是，'f' 表示否，默认为 'f'" xorm:"reportdeactivate not null default ''f'' comment('是否停用报告，'t' 表示是，'f' 表示否，默认为 'f'') ENUM('f','t')"`
	Comments         *string    `json:"comments" name:"" xorm:"comments default 'NULL' TEXT"`
	Updated          *time.Time `json:"updated" name:"" xorm:"updated not null DATETIME"`
}

func (o *AdAdvertisers) TableName() string {
	return "ad_advertisers"
}
func (o *AdAdvertisers) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
func (o *AdAdvertisers) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &o)
}
func (o *AdAdvertisers) PrimaryKey() interface{} {
	return o.Id
}
func (o *AdAdvertisers) NewEntity(dao mysql.BaseDao) mysql.BaseEntity {
	return mysql.NewEntity(dao, o)
}
