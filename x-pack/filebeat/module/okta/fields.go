// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package okta

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "okta", asset.ModuleFieldsPri, AssetOkta); err != nil {
		panic(err)
	}
}

// AssetOkta returns asset data.
// This is the base64 encoded gzipped contents of module/okta.
func AssetOkta() string {
	return "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"
}