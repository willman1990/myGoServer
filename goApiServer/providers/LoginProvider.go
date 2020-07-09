package providers

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"myGoWeb/goApiServer/dto"
)

type LoginProvider struct {
	myDb *sql.DB
}

func (this *LoginProvider) MakeCheckForLogin(login *dto.Users) (*dto.LoginToken, int, error) {
	//_, err := this.myDb.Exec("insert into users (username, password) values(?, ?)", "jinqiu", "loveyou")
	//if err != nil {
	//	fmt.Println("insert into database failed with errMsg ==> " + err.Error())
	//}

	users := make([]dto.Users,0)
	raws, err := this.myDb.Query("select * from users")
	if err != nil {
		fmt.Println("Query users from database failed , errMsg: " + err.Error())
		return nil, 500, errors.New("visit database failed")
	}
	for raws.Next() {
		var user dto.Users
		serr := raws.Scan(&user.UserName, &user.PassWord)
		if serr != nil {
			fmt.Println("scan failed , errMsg: " + serr.Error())
			return nil, 500, errors.New("visit database failed")
		} else {
			users = append(users, user)
		}

	}
	fmt.Println("users detail : %+v", users)
	for _, user := range users {
		if user.UserName == login.UserName && user.PassWord == login.PassWord {
			return &dto.LoginToken{Token:this.MD5Hash(user.UserName + user.PassWord)}, 200, nil
		}
	}
	errStr := fmt.Sprintf("User %s not exist, please regist first!", login.UserName)
	return nil, 403, errors.New(errStr)
}

func (this *LoginProvider) MD5Hash(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)
	md5Resp := hex.EncodeToString(cipherStr)
	fmt.Printf("%s\n", md5Resp) // 输出加密结果
	return md5Resp
}

func NewLoginProvider(mysqlDb *sql.DB) *LoginProvider {
	rsp := &LoginProvider{myDb: mysqlDb}
	return rsp
}
