package models

import (
	"AdServerCollector/core/mysql"
	"encoding/json"
	"time"
)

type AdPublishers struct {
	Id        *int       `json:"id" name:"" xorm:"id pk autoincr MEDIUMINT(9)"`
	Name      *string    `json:"name" name:"" xorm:"name not null default '''' VARCHAR(255)"`
	AccountId *int       `json:"account_id" name:"" xorm:"account_id default NULL unique MEDIUMINT(9)"`
	Comments  *string    `json:"comments" name:"" xorm:"comments default 'NULL' TEXT"`
	Contact   *string    `json:"contact" name:"" xorm:"contact default 'NULL' VARCHAR(255)"`
	Email     *string    `json:"email" name:"" xorm:"email not null default '''' VARCHAR(64)"`
	Website   *string    `json:"website" name:"" xorm:"website default 'NULL' VARCHAR(255)"`
	Updated   *time.Time `json:"updated" name:"" xorm:"updated not null DATETIME"`
}

func (o *AdPublishers) TableName() string {
	return "ad_publishers"
}
func (o *AdPublishers) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
func (o *AdPublishers) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &o)
}
func (o *AdPublishers) PrimaryKey() interface{} {
	return o.Id
}
func (o *AdPublishers) NewEntity(dao mysql.BaseDao) mysql.BaseEntity {
	return mysql.NewEntity(dao, o)
}
