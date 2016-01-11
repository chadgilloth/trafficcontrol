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

type Type struct {
	Id          int64       `db:"id" json:"id"`
	Name        string      `db:"name" json:"name"`
	Description null.String `db:"description" json:"description"`
	UseInTable  null.String `db:"use_in_table" json:"useInTable"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleType(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getType(id)
	} else if method == "POST" {
		return postType(payload)
	} else if method == "PUT" {
		return putType(id, payload)
	} else if method == "DELETE" {
		return delType(id)
	}
	return nil, nil
}

func getType(id int) (interface{}, error) {
	if id >= 0 {
		return getTypeById(id)
	} else {
		return getTypes()
	}
}

// @Title getTypeById
// @Description retrieves the type information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Type
// @Resource /api/2.0
// @Router /api/2.0/type/{id} [get]
func getTypeById(id int) (interface{}, error) {
	ret := []Type{}
	arg := Type{Id: int64(id)}
	nstmt, err := db.GlobalDB.PrepareNamed(`select * from type where id=:id`)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getTypes
// @Description retrieves the type information for a certain id
// @Accept  application/json
// @Success 200 {array}    Type
// @Resource /api/2.0
// @Router /api/2.0/type [get]
func getTypes() (interface{}, error) {
	ret := []Type{}
	queryStr := "select * from type"
	err := db.GlobalDB.Select(&ret, queryStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postType
// @Description enter a new type
// @Accept  application/json
// @Param                 Body body     Type   true "Type object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/type [post]
func postType(payload []byte) (interface{}, error) {
	var v Type
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO type("
	sqlString += "name"
	sqlString += ",description"
	sqlString += ",use_in_table"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:description"
	sqlString += ",:use_in_table"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putType
// @Description modify an existing typeentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     Type   true "Type object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/type/{id}  [put]
func putType(id int, payload []byte) (interface{}, error) {
	var v Type
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE type SET "
	sqlString += "name = :name"
	sqlString += ",description = :description"
	sqlString += ",use_in_table = :use_in_table"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delTypeById
// @Description deletes type information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Type
// @Resource /api/2.0
// @Router /api/2.0/type/{id} [delete]
func delType(id int) (interface{}, error) {
	arg := Type{Id: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM type WHERE id=:id", arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}
