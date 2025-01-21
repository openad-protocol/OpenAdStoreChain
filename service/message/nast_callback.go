package message

import (
	"AdServerCollector/conf"
	"AdServerCollector/constants"
	"AdServerCollector/core/attestation"
	AdErrors "AdServerCollector/core/errors"
	"AdServerCollector/core/mysql"
	"AdServerCollector/logger"
	"AdServerCollector/models"
	"AdServerCollector/utils"
	"encoding/json"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/nats-io/nats.go"
	"golang.org/x/crypto/sha3"
	"strconv"
	"time"
)

type MProcess struct {
	dbCon mysql.DBInterface
	sess  *xorm.Session
}

func NewMProcess() *MProcess {
	dbInterface := mysql.GetMySqlDB()
	if dbInterface == nil {
		logger.Errorf("NewMProcess: get dbInterface error: nil")
		return nil
	}
	return &MProcess{
		dbCon: dbInterface,
	}
}

// ProcessTracerHash 得到订阅消息后的处理
func (pt *MProcess) ProcessTracerHash(msg *nats.Msg) {
	var payload TraceInfo
	var traceInfoMessage models.AdDataRawTraceInfo
	sess := mysql.GetMySqlDB().NewSession()
	defer sess.Close()
	logger.Info("Received message: ", string(msg.Data))
	err := json.Unmarshal(msg.Data, &payload)
	if err != nil {
		logger.Errorf("Unmarshal failed: %v", err)
		return
	}
	traceInfoMessage.TraceId = payload.TraceId
	traceInfoMessage.LoginfoHash = payload.LogInfoHash
	traceInfoMessage.ClickinfoHash = payload.ClickInfoHash
	traceInfoMessage.CbHash = payload.CbHash
	_eventId := utils.CheckPointer[string](payload.EventId, "0")
	traceInfoMessage.EventId = _eventId
	traceInfoMessage.CreateAt = utils.ValueToPoint(time.Now().Unix())
	tableName := traceInfoMessage.GetSliceDateDayTable()
	if b, err := sess.IsTableExist(tableName); err != nil || !b {
		logger.Errorf("Table %s not exist: %v", tableName, err)
		if err = sess.Table(tableName).CreateTable(traceInfoMessage); err != nil {
			logger.Errorf("CreateTable failed: %v", err)
			return
		}
	}
	AdErrors.Try(func() {
		var _str []byte
		if _str, err = json.Marshal(traceInfoMessage); err != nil {
			logger.Errorf("Marshal failed: %s", err.Error())
		} else {
			logger.Infof("Received message traceInfoMessage string: %s", string(_str))
		}
		if _, err = sess.Table(tableName).InsertOne(&traceInfoMessage); err != nil {
			logger.Errorf("insert traceinfo error,%s table name :%s,data:%s", err.Error(), tableName, string(_str))
		}
	}).Catch(constants.ErrRuntimePanic, func(err error) {
		logger.Errorf("InsertOne failed: %s", err.Error())
	}).DefaultCatch(func(err error) {
		logger.Errorf("InsertOne failed: %s", err.Error())
	}).Do()
	traceMessageStr, _ := json.Marshal(traceInfoMessage)
	logger.Infof("InsertOne success: %s", traceMessageStr)
}

func (pt *MProcess) ProcessAdMessage(msg *nats.Msg) {
	var payload AdDataRaw
	var adMessage models.DataRawLog
	sess := mysql.GetMySqlDB().NewSession()
	defer sess.Close()
	logger.Info("Received message: ", string(msg.Data))
	if err := json.Unmarshal(msg.Data, &payload); err != nil {
		logger.Errorf("Unmarshal failed: %v", err)
		return
	}
	_zoneId, err := strconv.ParseInt(*utils.CheckPointer[string](payload.ZoneId, "0"), 10, 64)
	if err != nil {
		logger.Errorf("ParseInt failed: %v", err)
		return
	}
	adMessage.ZoneId = &_zoneId
	_publisherId, err := strconv.ParseInt(*utils.CheckPointer[string](payload.PublisherId, "0"), 10, 64)
	if err != nil {
		logger.Errorf("ParseInt failed: %v", err)
		return
	}
	adMessage.PublisherId = &_publisherId
	_eventId, err := strconv.ParseInt(*utils.CheckPointer[string](payload.EventId, "0"), 10, 64)
	if err != nil {
		logger.Errorf("ParseInt failed: %v", err)
		return
	}
	adMessage.EventId = &_eventId

	if payload.TimeStamp == nil || *payload.TimeStamp == "" {
		logger.Errorf("TimeStamp is nil")
		return
	}
	adMessage.FirstName = utils.TruncateString(payload.FirstName, 128)
	adMessage.LastName = utils.TruncateString(payload.LastName, 128)
	adMessage.UserName = utils.TruncateString(payload.UserName, 128)
	adMessage.Hash = utils.TruncateString(payload.Hash, 64)
	adMessage.FromType = utils.TruncateString(payload.FromType, 25)
	adMessage.Language = utils.TruncateString(payload.Language, 8)
	adMessage.Location = payload.Location
	adMessage.Platform = utils.TruncateString(payload.Platform, 12)
	adMessage.Signature = utils.TruncateString(payload.Signature, 64)
	adMessage.TraceId = utils.TruncateString(payload.TraceId, 64)
	adMessage.UserId = utils.TruncateString(payload.UserId, 64)
	adMessage.Version = utils.TruncateString(payload.Version, 64)
	adMessage.IpAddress = utils.TruncateString(payload.IpAddress, 64)
	adMessage.Country = utils.TruncateString(payload.Country, 8)
	adMessage.RequestType = utils.TruncateString(payload.RequestType, 12)
	adMessage.WalletType = utils.TruncateString(payload.WalletType, 64)
	adMessage.WalletAddress = utils.TruncateString(payload.WalletAddress, 64)
	adMessage.IsPremium = utils.TruncateString(payload.IsPremium, 64)
	adMessage.CreateAt = utils.ValueToPoint(time.Now().UnixMilli())
	tm, err := strconv.ParseFloat(*payload.TimeStamp, 64)
	if err != nil {
		logger.Errorf("ParseInt failed: %v", err)
		return
	}
	tmInt64 := int64(tm * 1000)
	adMessage.TimeStamp = &tmInt64
	strMsg, _ := json.Marshal(adMessage)
	logger.Infof("Received message adMessage string: %s", string(strMsg))
	tableName := adMessage.GetSliceDateDayTable()
	_ = AdErrors.TryFunc(func() {
		if b, err := sess.IsTableExist(tableName); err != nil || !b {
			logger.Errorf("Table %s not exist: %v", tableName, err)
			if err = sess.Table(tableName).CreateTable(adMessage); err != nil {
				logger.Errorf("CreateTable failed: %v", err)
				return
			}
		}
	}, func(err interface{}) {
		logger.Errorf("CreateTable failed: %v", err)
		return
	}, func() {
		logger.Debug("CreateTable finally")
		return
	})
	AdErrors.Try(func() {
		if _, err := sess.Table(tableName).InsertOne(&adMessage); err != nil {
			_str, _ := json.Marshal(adMessage)
			logger.Infof("InsertOne adMessage %s,data:%s error: %s", tableName, _str, err)
		}
	}).Catch(constants.ErrRuntimePanic, func(err error) {
		logger.Errorf("InsertOne adMessage failed: %s", err.Error())
	}).DefaultCatch(func(err error) {
		logger.Errorf("InsertOne adMessage failed: %s", err.Error())
	}).Do()
	var payloadStr string
	if payloadHash, err := json.Marshal(adMessage); err != nil {
		logger.Errorf("Marshal failed: %s", err.Error())
		return
	} else {
		payloadStr = string(payloadHash)
	}
	logger.Infof("Received message payload string : %s", payloadStr)
	_hash := sha3.Sum256([]byte(payloadStr))
	adMessage.TraceHash = utils.ValueToPoint(fmt.Sprintf("%x", _hash))
	strMsg, _ = json.Marshal(adMessage)
	logger.Infof("databases adMessage string: %s", string(strMsg))
	// 成功插入数据库，取得id则更新
	if adMessage.Id != nil {
		logger.Infof("adMessage.TraceHash: %s, db id:%d", *adMessage.TraceHash, *adMessage.Id)
		sql := fmt.Sprintf("update %s set trace_hash = '%s' where id = %d", tableName, *adMessage.TraceHash, *adMessage.Id)
		logger.Infof("update sql: %s,table name:%s", sql, tableName)
		AdErrors.Try(func() {
			if _, err := sess.Exec(sql); err != nil {
				logger.Error("Update failed: %s", err.Error())
			} else {
				logger.Infof("Update success: %s", *adMessage.TraceHash)
			}
		}).Catch(constants.ErrRuntimePanic, func(err error) {
			logger.Error("Update failed: ", err.Error())
		}).DefaultCatch(func(err error) {
			logger.Error("Update failed: ", err.Error())
		}).Do()
	}
	switch *adMessage.RequestType {
	case "getAd":
		if conf.Config.Hash.GetAd {
			if err := attestation.GetAdTree.SubmitAttestation(_hash); err != nil {
				logger.Errorf("getAd SubmitAttestation failed: %v", err)
				return
			}
		}
	case "loginfo":
		if conf.Config.Hash.LogInfo {
			if err := attestation.LogTree.SubmitAttestation(_hash); err != nil {
				logger.Errorf("loginfo SubmitAttestation failed: %v", err)
				return
			}
		}
	case "clickinfo":
		if conf.Config.Hash.ClickInfo {
			if err := attestation.ClickTree.SubmitAttestation(_hash); err != nil {
				logger.Errorf("clickinfo SubmitAttestation failed: %v", err)
				return
			}
		}
	case "ad_in_call":
		if conf.Config.Hash.AdInCall {
			if err := attestation.CallBackTree.SubmitAttestation(_hash); err != nil {
				logger.Errorf("callback SaveVersionStateDB failed: %v", err)
				return
			}
		}
	default:
		logger.Errorf("RequestType not found: %s", *adMessage.RequestType)
	}

	return
}

func (pt *MProcess) ProcessAdMissing(msg *nats.Msg) {
	var payload AdDataRaw
	var adMessage models.DataRawLogMissing
	sess := mysql.GetMySqlDB().NewSession()
	defer sess.Close()
	logger.Info("Received missing message: ", string(msg.Data))
	if err := json.Unmarshal(msg.Data, &payload); err != nil {
		logger.Errorf("Unmarshal failed: %v", err)
		return
	}
	_zoneId, err := strconv.ParseInt(*utils.CheckPointer[string](payload.ZoneId, "0"), 10, 64)
	if err != nil {
		logger.Errorf("ParseInt failed: %v", err)
		return
	}
	adMessage.ZoneId = &_zoneId
	_publisherId, err := strconv.ParseInt(*utils.CheckPointer[string](payload.PublisherId, "0"), 10, 64)
	if err != nil {
		logger.Errorf("ParseInt failed: %v", err)
		return
	}
	adMessage.PublisherId = &_publisherId
	_eventIdStr := utils.CheckPointer[string](payload.EventId, "0")
	if *_eventIdStr == "nil" {
		*_eventIdStr = "0"
	}
	_eventId, err := strconv.ParseInt(*_eventIdStr, 10, 64)
	if err != nil {
		logger.Errorf("ParseInt failed: %v", err)
		return
	}
	adMessage.EventId = &_eventId
	adMessage.FirstName = utils.TruncateString(payload.FirstName, 128)
	adMessage.LastName = utils.TruncateString(payload.LastName, 128)
	adMessage.UserName = utils.TruncateString(payload.UserName, 128)
	adMessage.Hash = utils.TruncateString(payload.Hash, 64)
	adMessage.FromType = utils.TruncateString(payload.FromType, 25)
	adMessage.Language = utils.TruncateString(payload.Language, 8)
	adMessage.Location = payload.Location
	adMessage.Platform = utils.TruncateString(payload.Platform, 12)
	adMessage.Signature = utils.TruncateString(payload.Signature, 64)
	adMessage.TraceId = utils.TruncateString(payload.TraceId, 128)
	adMessage.UserId = utils.TruncateString(payload.UserId, 64)
	adMessage.Version = utils.TruncateString(payload.Version, 64)
	adMessage.IpAddress = utils.TruncateString(payload.IpAddress, 64)
	adMessage.Country = utils.TruncateString(payload.Country, 8)
	adMessage.RequestType = utils.TruncateString(payload.RequestType, 12)
	adMessage.WalletType = utils.TruncateString(payload.WalletType, 64)
	adMessage.WalletAddress = utils.TruncateString(payload.WalletAddress, 64)
	adMessage.IsPremium = utils.TruncateString(payload.IsPremium, 64)
	adMessage.CreateAt = utils.ValueToPoint(time.Now().UnixMilli())
	if payload.TimeStamp == nil || *payload.TimeStamp == "" {
		logger.Errorf("TimeStamp is nil")
		return
	}
	tm, err := strconv.ParseFloat(*payload.TimeStamp, 64)
	if err != nil {
		logger.Errorf("ParseInt failed: %v", err)
		return
	}
	tmInt64 := int64(tm * 1000)
	adMessage.TimeStamp = &tmInt64

	strMsg, _ := json.Marshal(adMessage)
	logger.Infof("Received message adMessage string: %s", string(strMsg))
	tableName := adMessage.GetSliceDateDayTable()
	_ = AdErrors.TryFunc(func() {
		if b, err := sess.IsTableExist(tableName); err != nil || !b {
			logger.Errorf("Table %s not exist: %v", tableName, err)
			if err = sess.Table(tableName).CreateTable(adMessage); err != nil {
				logger.Errorf("CreateTable failed: %v", err)
				return
			}
		}
	}, func(err interface{}) {
		logger.Errorf("CreateTable failed: %v", err)
		return
	}, func() {
		logger.Debug("CreateTable finally")
		return
	})
	logger.Infof("InsertOne missing adMessage table: %v", tableName)
	AdErrors.Try(func() {
		if _, err := sess.Table(tableName).InsertOne(&adMessage); err != nil {
			logger.Errorf("insert traceinfo error,%s table name :%s", err.Error(), tableName)
		}
	}).Catch(constants.ErrRuntimePanic, func(err error) {
		logger.Errorf("InsertOne failed: %s", err.Error())
	}).DefaultCatch(func(err error) {
		logger.Errorf("InsertOne failed: %s", err.Error())
	}).Do()
}
