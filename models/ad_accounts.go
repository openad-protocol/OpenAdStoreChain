package models

import (
	"AdServerCollector/core/mysql"
	"encoding/json"
	"time"
)

type AdAccounts struct {
	Id               *int       `json:"id" name:"" xorm:"id pk autoincr MEDIUMINT(9)"`
	Username         *string    `json:"username" name:"" xorm:"username not null default '''' unique VARCHAR(64)"`
	Password         *string    `json:"password" name:"" xorm:"password not null default '''' VARCHAR(64)"`
	Type             *string    `json:"type" name:"" xorm:"type not null VARCHAR(255)"`
	ContactName      *string    `json:"contact_name" name:"" xorm:"contact_name default '''' VARCHAR(255)"`
	EmailAddress     *string    `json:"email_address" name:"" xorm:"email_address default '''' VARCHAR(64)"`
	Language         *string    `json:"language" name:"" xorm:"language default 'NULL' VARCHAR(5)"`
	Comments         *string    `json:"comments" name:"" xorm:"comments default 'NULL' TEXT"`
	Active           *bool      `json:"active" name:"" xorm:"active not null default 1 TINYINT(1)"`
	DateCreated      *time.Time `json:"date_created" name:"" xorm:"date_created default 'NULL' DATETIME"`
	DateLastLogin    *time.Time `json:"date_last_login" name:"" xorm:"date_last_login default 'NULL' DATETIME"`
	EmailUpdated     *time.Time `json:"email_updated" name:"" xorm:"email_updated default 'NULL' DATETIME"`
	Balance          *float64   `json:"balance" name:"余额" xorm:"balance default NULL comment('余额') DECIMAL(8,6)"`
	Salt             *string    `json:"salt" name:"" xorm:"salt default 'NULL' VARCHAR(255)"`
	LoginIp          *string    `json:"login_ip" name:"" xorm:"login_ip default 'NULL' VARCHAR(25)"`
	CompanyName      *string    `json:"company_name" name:"" xorm:"company_name default 'NULL' VARCHAR(255)"`
	LockedBalance    *float64   `json:"locked_balance" name:"锁定余额" xorm:"locked_balance comment('锁定余额') DECIMAL(8,6)"`
	AvailableBalance *float64   `json:"available_balance" name:"可用余额=总余额-已锁定余额" xorm:"available_balance comment('可用余额=总余额-已锁定余额') DECIMAL(8,6)"`
}

func (o *AdAccounts) TableName() string {
	return "ad_accounts"
}
func (o *AdAccounts) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
func (o *AdAccounts) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &o)
}
func (o *AdAccounts) PrimaryKey() interface{} {
	return o.Id
}
func (o *AdAccounts) NewEntity(dao mysql.BaseDao) mysql.BaseEntity {
	return mysql.NewEntity(dao, o)
}
