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
	"time"
)

type FederationFederationResolver struct {
	Federation         int64     `db:"federation" json:"federation"`
	FederationResolver int64     `db:"federation_resolver" json:"federationResolver"`
	LastUpdated        time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleFederationFederationResolver(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getFederationFederationResolver(id)
	} else if method == "POST" {
		return postFederationFederationResolver(payload)
	} else if method == "PUT" {
		return putFederationFederationResolver(id, payload)
	} else if method == "DELETE" {
		return delFederationFederationResolver(id)
	}
	return nil, nil
}

func getFederationFederationResolver(id int) (interface{}, error) {
	if id >= 0 {
		return getFederationFederationResolverById(id)
	} else {
		return getFederationFederationResolvers()
	}
}

// @Title getFederationFederationResolverById
// @Description retrieves the federation_federation_resolver information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    FederationFederationResolver
// @Resource /api/2.0
// @Router /api/2.0/federation_federation_resolver/{id} [get]
func getFederationFederationResolverById(id int) (interface{}, error) {
	ret := []FederationFederationResolver{}
	arg := FederationFederationResolver{Federation: int64(id)}
	nstmt, err := db.GlobalDB.PrepareNamed(`select * from federation_federation_resolver where federation=:federation`)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getFederationFederationResolvers
// @Description retrieves the federation_federation_resolver information for a certain id
// @Accept  application/json
// @Success 200 {array}    FederationFederationResolver
// @Resource /api/2.0
// @Router /api/2.0/federation_federation_resolver [get]
func getFederationFederationResolvers() (interface{}, error) {
	ret := []FederationFederationResolver{}
	queryStr := "select * from federation_federation_resolver"
	err := db.GlobalDB.Select(&ret, queryStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postFederationFederationResolver
// @Description enter a new federation_federation_resolver
// @Accept  application/json
// @Param                 Body body     FederationFederationResolver   true "FederationFederationResolver object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/federation_federation_resolver [post]
func postFederationFederationResolver(payload []byte) (interface{}, error) {
	var v FederationFederationResolver
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO federation_federation_resolver("
	sqlString += "federation"
	sqlString += ",federation_resolver"
	sqlString += ") VALUES ("
	sqlString += ":federation"
	sqlString += ",:federation_resolver"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putFederationFederationResolver
// @Description modify an existing federation_federation_resolverentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     FederationFederationResolver   true "FederationFederationResolver object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/federation_federation_resolver/{id}  [put]
func putFederationFederationResolver(id int, payload []byte) (interface{}, error) {
	var v FederationFederationResolver
	err := json.Unmarshal(payload, &v)
	v.Federation = int64(id) // overwrite the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE federation_federation_resolver SET "
	sqlString += "federation = :federation"
	sqlString += ",federation_resolver = :federation_resolver"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE federation=:federation"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delFederationFederationResolverById
// @Description deletes federation_federation_resolver information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    FederationFederationResolver
// @Resource /api/2.0
// @Router /api/2.0/federation_federation_resolver/{id} [delete]
func delFederationFederationResolver(id int) (interface{}, error) {
	arg := FederationFederationResolver{Federation: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM federation_federation_resolver WHERE id=:id", arg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}
