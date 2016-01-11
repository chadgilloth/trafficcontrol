// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_goto2.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	"fmt"
	"github.com/Comcast/traffic_control/traffic_ops/goto2/db"
	_ "github.com/Comcast/traffic_control/traffic_ops/goto2/output_format" // needed for swagger
	null "gopkg.in/guregu/null.v3"
	"time"
)

type TmUser struct {
	Id                 int64       `db:"id" json:"id"`
	Username           null.String `db:"username" json:"username"`
	Role               null.Int    `db:"role" json:"role"`
	Uid                null.Int    `db:"uid" json:"uid"`
	Gid                null.Int    `db:"gid" json:"gid"`
	LocalPasswd        null.String `db:"local_passwd" json:"localPasswd"`
	ConfirmLocalPasswd null.String `db:"confirm_local_passwd" json:"confirmLocalPasswd"`
	LastUpdated        time.Time   `db:"last_updated" json:"lastUpdated"`
	Company            null.String `db:"company" json:"company"`
	Email              null.String `db:"email" json:"email"`
	FullName           null.String `db:"full_name" json:"fullName"`
	NewUser            int64       `db:"new_user" json:"newUser"`
	AddressLine1       null.String `db:"address_line1" json:"addressLine1"`
	AddressLine2       null.String `db:"address_line2" json:"addressLine2"`
	City               null.String `db:"city" json:"city"`
	StateOrProvince    null.String `db:"state_or_province" json:"stateOrProvince"`
	PhoneNumber        null.String `db:"phone_number" json:"phoneNumber"`
	PostalCode         null.String `db:"postal_code" json:"postalCode"`
	Country            null.String `db:"country" json:"country"`
	LocalUser          int64       `db:"local_user" json:"localUser"`
	Token              null.String `db:"token" json:"token"`
	RegistrationSent   time.Time   `db:"registration_sent" json:"registrationSent"`
}

func handleTmUser(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getTmUser(id)
	} else if method == "POST" {
		return postTmUser(payload)
	} else if method == "PUT" {
		return putTmUser(id, payload)
	} else if method == "DELETE" {
		return delTmUser(id)
	}
	return nil, nil
}

func getTmUser(id int) (interface{}, error) {
	if id >= 0 {
		return getTmUserById(id)
	} else {
		return getTmUsers()
	}
}

// @Title getTmUserById
// @Description retrieves the tm_user information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    TmUser
// @Resource /api/2.0
// @Router /api/2.0/tm_user/{id} [get]
func getTmUserById(id int) (interface{}, error) {
	ret := []TmUser{}
	arg := TmUser{Id: int64(id)}
	nstmt, err := db.GlobalDB.PrepareNamed(`select * from tm_user where id=:id`)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getTmUsers
// @Description retrieves the tm_user information for a certain id
// @Accept  application/json
// @Success 200 {array}    TmUser
// @Resource /api/2.0
// @Router /api/2.0/tm_user [get]
func getTmUsers() (interface{}, error) {
	ret := []TmUser{}
	queryStr := "select * from tm_user"
	err := db.GlobalDB.Select(&ret, queryStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postTmUser
// @Description enter a new tm_user
// @Accept  application/json
// @Param                 Body body     TmUser   true "TmUser object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/tm_user [post]
func postTmUser(payload []byte) (interface{}, error) {
	var v TmUser
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO tm_user("
	sqlString += "username"
	sqlString += ",role"
	sqlString += ",uid"
	sqlString += ",gid"
	sqlString += ",local_passwd"
	sqlString += ",confirm_local_passwd"
	sqlString += ",company"
	sqlString += ",email"
	sqlString += ",full_name"
	sqlString += ",new_user"
	sqlString += ",address_line1"
	sqlString += ",address_line2"
	sqlString += ",city"
	sqlString += ",state_or_province"
	sqlString += ",phone_number"
	sqlString += ",postal_code"
	sqlString += ",country"
	sqlString += ",local_user"
	sqlString += ",token"
	sqlString += ",registration_sent"
	sqlString += ") VALUES ("
	sqlString += ":username"
	sqlString += ",:role"
	sqlString += ",:uid"
	sqlString += ",:gid"
	sqlString += ",:local_passwd"
	sqlString += ",:confirm_local_passwd"
	sqlString += ",:company"
	sqlString += ",:email"
	sqlString += ",:full_name"
	sqlString += ",:new_user"
	sqlString += ",:address_line1"
	sqlString += ",:address_line2"
	sqlString += ",:city"
	sqlString += ",:state_or_province"
	sqlString += ",:phone_number"
	sqlString += ",:postal_code"
	sqlString += ",:country"
	sqlString += ",:local_user"
	sqlString += ",:token"
	sqlString += ",:registration_sent"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putTmUser
// @Description modify an existing tm_userentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     TmUser   true "TmUser object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/tm_user/{id}  [put]
func putTmUser(id int, payload []byte) (interface{}, error) {
	var v TmUser
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE tm_user SET "
	sqlString += "username = :username"
	sqlString += ",role = :role"
	sqlString += ",uid = :uid"
	sqlString += ",gid = :gid"
	sqlString += ",local_passwd = :local_passwd"
	sqlString += ",confirm_local_passwd = :confirm_local_passwd"
	sqlString += ",last_updated = :last_updated"
	sqlString += ",company = :company"
	sqlString += ",email = :email"
	sqlString += ",full_name = :full_name"
	sqlString += ",new_user = :new_user"
	sqlString += ",address_line1 = :address_line1"
	sqlString += ",address_line2 = :address_line2"
	sqlString += ",city = :city"
	sqlString += ",state_or_province = :state_or_province"
	sqlString += ",phone_number = :phone_number"
	sqlString += ",postal_code = :postal_code"
	sqlString += ",country = :country"
	sqlString += ",local_user = :local_user"
	sqlString += ",token = :token"
	sqlString += ",registration_sent = :registration_sent"
	sqlString += " WHERE id=:id"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delTmUserById
// @Description deletes tm_user information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    TmUser
// @Resource /api/2.0
// @Router /api/2.0/tm_user/{id} [delete]
func delTmUser(id int) (interface{}, error) {
	arg := TmUser{Id: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM tm_user WHERE id=:id", arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}
