package sms

import (
	"errors"

	"github.com/goinggo/mapstructure"
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
)

// 发送短信验证码
func SendSMSCode(apiKey, phoneNumber, code, templateID string) error {
	clnt := ypclnt.New(apiKey)
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = phoneNumber
	param[ypclnt.TPL_ID] = templateID
	param[ypclnt.TPL_VALUE] = "#code#=" + code
	r := clnt.Sms().TplSingleSend(param)
	if r.Code == 0 {
		// 发送成功
		return nil
	}
	// 如果发送失败，可以从r.Detail和r.Msg中查看失败原因
	return errors.New(r.Detail)
}

// 获取账号信息
func GetAccount(apiKey string) (*AccountInfo, error) {
	accountInfo := &AccountInfo{}
	clnt := ypclnt.New(apiKey)
	r := clnt.User().Get(nil)
	if r.Code == 0 {
		if err := mapstructure.Decode(r.Data, &accountInfo); err != nil {
			return nil, err
		}
		return accountInfo, nil
	}
	return nil, errors.New(r.Detail)
}

// 余额警告值,余额为0，会发不出来短信
func (a AccountInfo) IsAlarmBalance(alarmBalance float64) bool {
	if alarmBalance != 0 && a.Balance < alarmBalance {
		return true
	}
	return a.Balance < a.AlarmBalance
}

type AccountInfo struct {
	Nick             string  `json:"nick"`
	CreatedTime      string  `json:"gmt_created"`
	Mobile           string  `json:"mobile"`
	ApiVersion       string  `json:"api_version"`
	Email            string  `json:"email"`
	Balance          float64 `json:"balance"`
	AlarmBalance     float64 `json:"alarm_balance"`
	EmergencyContact string  `json:"emergency_contact"`
	EmergencyMobile  string  `json:"emergency_mobile"`
}
