package models

import (
	"AdServerCollector/core/mysql"
	"encoding/json"
	"time"
)

type AdResources struct {
	Id             *int       `json:"id" name:"" xorm:"id pk autoincr MEDIUMINT(9)"`
	AdvertiserId   *int       `json:"advertiser_id" name:"广告主id" xorm:"advertiser_id not null default 0 comment('广告主id') index MEDIUMINT(9)"`
	BannerType     *int       `json:"banner_type" name:"广告类型：与版位的类型匹配" xorm:"banner_type not null default 0 comment('广告类型：与版位的类型匹配') TINYINT(4)"`
	ContentType    *string    `json:"content_type" name:"广告内容的类型，默认为 'gif'" xorm:"content_type not null default ''gif'' comment('广告内容的类型，默认为 'gif'') VARCHAR(8)"`
	Url            *string    `json:"url" name:"" xorm:"url not null TEXT"`
	ImageUrl       *string    `json:"image_url" name:"图片 URL，默认为空字符串" xorm:"image_url not null default '''' comment('图片 URL，默认为空字符串') VARCHAR(255)"`
	IconUrl        *string    `json:"icon_url" name:"icon的url" xorm:"icon_url default 'NULL' comment('icon的url') VARCHAR(255)"`
	ClickUrl       *string    `json:"click_url" name:"点击后的url" xorm:"click_url default 'NULL' comment('点击后的url') VARCHAR(255)"`
	Metadata       *string    `json:"metadata" name:"元数据" xorm:"metadata default 'NULL' comment('元数据') VARCHAR(255)"`
	StorageType    *string    `json:"storage_type" name:"存储类型，默认为 'sql'" xorm:"storage_type not null default ''sql'' comment('存储类型，默认为 'sql'') VARCHAR(16)"`
	FileName       *string    `json:"file_name" name:"文件名，默认为空字符串" xorm:"file_name not null default '''' comment('文件名，默认为空字符串') VARCHAR(255)"`
	HtmlTemplate   *string    `json:"html_template" name:"HTML 模板" xorm:"html_template not null comment('HTML 模板') MEDIUMTEXT"`
	HtmlCache      *string    `json:"html_cache" name:"HTML 缓存" xorm:"html_cache not null comment('HTML 缓存') MEDIUMTEXT"`
	Width          *int       `json:"width" name:"宽度" xorm:"width not null default 0 comment('宽度') SMALLINT(6)"`
	Height         *int       `json:"height" name:"高度" xorm:"height not null default 0 comment('高度') SMALLINT(6)"`
	Weight         *int       `json:"weight" name:"权重，默认为 '1'" xorm:"weight not null default 1 comment('权重，默认为 '1'') TINYINT(4)"`
	Seq            *int       `json:"seq" name:"序列" xorm:"seq not null default 0 comment('序列') TINYINT(4)"`
	Alt            *string    `json:"alt" name:"替代文本" xorm:"alt not null default '''' comment('替代文本') VARCHAR(255)"`
	HoverText      *string    `json:"hover_text" name:"悬停文本" xorm:"hover_text default 'NULL' comment('悬停文本') VARCHAR(255)"`
	StatusText     *string    `json:"status_text" name:"状态文本" xorm:"status_text not null default '''' comment('状态文本') VARCHAR(255)"`
	BannerText     *string    `json:"banner_text" name:"广告文本" xorm:"banner_text not null comment('广告文本') TEXT"`
	Description    *string    `json:"description" name:"描述" xorm:"description not null default '''' comment('描述') VARCHAR(255)"`
	AclPlugins     *string    `json:"acl_plugins" name:"ACL 插件" xorm:"acl_plugins default 'NULL' comment('ACL 插件') TEXT"`
	PluginVersion  *int       `json:"plugin_version" name:"插件版本" xorm:"plugin_version not null default 0 comment('插件版本') MEDIUMINT(9)"`
	Append         *string    `json:"append" name:"附加内容" xorm:"append not null comment('附加内容') TEXT"`
	AltFilename    *string    `json:"alt_filename" name:"替代文件名" xorm:"alt_filename not null default '''' comment('替代文件名') VARCHAR(255)"`
	AltImageurl    *string    `json:"alt_imageurl" name:"替代图片 URL" xorm:"alt_imageurl not null default '''' comment('替代图片 URL') VARCHAR(255)"`
	AltContenttype *string    `json:"alt_contenttype" name:"替代内容类型" xorm:"alt_contenttype not null default ''gif'' comment('替代内容类型') VARCHAR(8)"`
	Comments       *string    `json:"comments" name:"" xorm:"comments default 'NULL' TEXT"`
	Updated        *time.Time `json:"updated" name:"" xorm:"updated not null DATETIME"`
	AclsUpdated    *time.Time `json:"acls_updated" name:"ACL 更新时间" xorm:"acls_updated not null comment('ACL 更新时间') DATETIME"`
	Keyword        *string    `json:"keyword" name:"关键词" xorm:"keyword not null default '''' comment('关键词') VARCHAR(255)"`
	Transparent    *bool      `json:"transparent" name:"是否透明，默认为 '0'" xorm:"transparent not null default 0 comment('是否透明，默认为 '0'') TINYINT(1)"`
	Parameters     *string    `json:"parameters" name:"" xorm:"parameters default 'NULL' TEXT"`
	Status         *int       `json:"status" name:"状态" xorm:"status not null default 0 comment('状态') INT(11)"`
	Prepend        *string    `json:"prepend" name:"前置内容" xorm:"prepend not null comment('前置内容') TEXT"`
	IframeFriendly *bool      `json:"iframe_friendly" name:"是否支持 iframe，默认为 '1'" xorm:"iframe_friendly not null default 1 comment('是否支持 iframe，默认为 '1'') TINYINT(1)"`
	CreateTime     *time.Time `json:"create_time" name:"" xorm:"create_time not null default 'CURRENT_TIMESTAMP' comment('') TIMESTAMP"`
}

func (o *AdResources) TableName() string {
	return "ad_resources"
}
func (o *AdResources) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
func (o *AdResources) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &o)
}
func (o *AdResources) PrimaryKey() interface{} {
	return o.Id
}
func (o *AdResources) NewEntity(dao mysql.BaseDao) mysql.BaseEntity {
	return mysql.NewEntity(dao, o)
}
