// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package executor

import (
	"errors"
	"strconv"

	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func Warning(row Row) error {
	if len(row) != 3 {
		return errors.New("warning table should have 3 columns")
	}

	code, err := strconv.Atoi(row[1])
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"code": code,
		"msg":  row[2],
	}).Debug("sql warning")

	return &mysql.MySQLError{Number: uint16(code), Message: row[2]}
}
