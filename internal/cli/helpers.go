// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package cli

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/10gen/mcli/internal/convert"
	"github.com/mongodb-labs/pcgc/cloudmanager"
)

func validURL(val interface{}) error {
	_, err := url.ParseRequestURI(val.(string))
	if err != nil {
		return errors.New("the value is not a valid URL")
	}
	return nil
}

func deploymentStatusMessage(baseURL, projectID string) string {
	return fmt.Sprintf("Changes are being applied, please check %s/v2/%s#deployment/topology for status\n", baseURL, projectID)
}

func clusterExists(c *cloudmanager.AutomationConfig, name string) bool {
	i := convert.SearchReplicaSets(c.ReplicaSets, func(r *cloudmanager.ReplicaSet) bool {
		return r.ID == name
	})

	return i >= 0
}
