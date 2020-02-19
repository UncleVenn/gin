package model

import (
	"config"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

type User struct {
	UId      int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	password string `json:"password" form:"password"`
	AuthCode string `json:"authCode" form:"authCode"`
}

var UserInfo User

func LoadData(c *gin.Context) (err error) {
	u := make(map[string]string)
	body, _ := c.GetRawData()
	err = json.Unmarshal(body, &u)
	UserInfo.Username = u["username"]
	UserInfo.password = u["password"]
	return
}
func (u *User) GetUserByUsername() (err error) {
	err = config.Db.QueryRow("select username,password from user where username=?", u.Username).Scan(&u.Username, &u.password)
	return
}
func (u *User) GetUserByAuthCode() (err error) {
	err = config.Db.QueryRow("select id,username from user where authCode=?", u.AuthCode).Scan(&u.UId, &u.Username)
	return
}
func GetPassword() (pass string) {
	md5str := md5.New()
	md5str.Write([]byte(UserInfo.password))
	pass = hex.EncodeToString(md5str.Sum(nil))
	return
}

func CheckPassword(pass string) bool {
	return pass == UserInfo.password
}

func SetAuthCode() {
	UserInfo.AuthCode = getAuthCode()
	config.Db.Exec("update user set authCode = ? where username=?", UserInfo.AuthCode, UserInfo.Username)
}

func getAuthCode() string {
	// 获取当前时间的时间戳
	t := time.Now().Unix()
	// 生成一个MD5的哈希
	h := md5.New()
	// 将时间戳转换为byte，并写入哈希
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(t))
	h.Write([]byte(b))
	// 将字节流转化为16进制的字符串
	return hex.EncodeToString(h.Sum(nil))
}
