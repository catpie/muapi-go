package musdk

import (
	ss "github.com/orvice/shadowsocks-go/shadowsocks"
	"strings"
)

const (
	StatusEnable  = 1
	StatusDisable = 0
)

type VUser struct {
	Email   string `json:"email"`
	UUID    string `json:"uuid"`
	AlterID uint32 `json:"alter_id"`
	Level   uint32 `json:"level"`
}

func (v *VUser) GetEmail() string {
	return v.Email
}

func (v *VUser) GetUUID() string {
	return v.UUID
}

func (v *VUser) GetAlterID() uint32 {
	return v.AlterID
}

func (v *VUser) GetLevel() uint32 {
	return v.Level
}

type User struct {
	Id             int64  `json:"id"`
	Port           int    `json:"port"`
	Passwd         string `json:"passwd"`
	Method         string `json:"method"`
	Enable         int    `json:"enable"`
	TransferEnable int64  `json:"transfer_enable"`
	U              int64  `json:"u"`
	D              int64  `json:"d"`
	IsAdmin        int    `json:"is_admin"`

	V2rayUser VUser `json:"v2ray_user"`
}

func (u User) GetPort() int {
	return u.Port
}

func (u User) GetId() int64 {
	return u.Id
}

func (u User) GetPasswd() string {
	return u.Passwd
}

func (u User) GetMethod() string {
	return u.Method
}

func (u User) IsEnable() bool {
	if u.Enable == 0 {
		return false
	}
	if u.TransferEnable < (u.U + u.D) {
		return false
	}
	return true
}

func (u User) Admin() bool {
	if u.IsAdmin == 0 {
		return false
	}
	return true
}

func (u User) GetCipher() (*ss.Cipher, error, bool) {
	method := u.Method
	auth := false

	if strings.HasSuffix(method, "-auth") {
		method = method[:len(method)-5]
		auth = true
	}
	s, e := ss.NewCipher(method, u.Passwd)
	return s, e, auth
}
