/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package gosdk

import (
	"reflect"
	"testing"
)

func TestAccessConnectorConfig_DocumentOperations(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	tests := []struct {
		name   string
		fields fields
		want   *DocumentOperationsMethods
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			if got := ac.DocumentOperations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.DocumentOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_createDocument(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource string
		doc        DocumentJSON
		options    CreateDocumentOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ac.createDocument(tt.args.dataSource, tt.args.doc, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.createDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.createDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_getDocument(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource string
		unid       string
		options    GetDocumentOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ac.getDocument(tt.args.dataSource, tt.args.unid, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.getDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.getDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_updateDocument(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource string
		doc        DocumentInfo
		options    UpdateDocumentOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ac.updateDocument(tt.args.dataSource, tt.args.doc, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.updateDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.updateDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_patchDocument(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource string
		unid       string
		doc        DocumentJSON
		options    UpdateDocumentOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ac.patchDocument(tt.args.dataSource, tt.args.unid, tt.args.doc, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.patchDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.patchDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_deleteDocument(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource string
		doc        DocumentInfo
		mode       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ac.deleteDocument(tt.args.dataSource, tt.args.doc, tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.deleteDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.deleteDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_deleteDocumentByUnid(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource string
		unid       string
		mode       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ac.deleteDocumentByUnid(tt.args.dataSource, tt.args.unid, tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.deleteDocumentByUnid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.deleteDocumentByUnid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_bulkGetDocument(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource string
		unids      []string
		options    BulkGetDocumentOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ac.bulkGetDocument(tt.args.dataSource, tt.args.unids, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.bulkGetDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.bulkGetDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_getDocumentByQuery(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource string
		getRequest GetDocumentByQueryRequest
		actions    string
		options    GetDocumentByQueryOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ac.getDocumentByQuery(tt.args.dataSource, tt.args.getRequest, tt.args.actions, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.getDocumentByQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.getDocumentByQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_bulkCreateDocument(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource string
		docs       []DocumentJSON
		richTextAs RichTextRepresentation
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ac.bulkCreateDocument(tt.args.dataSource, tt.args.docs, tt.args.richTextAs)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.bulkCreateDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.bulkCreateDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_bulkUpdateDocumentByQuery(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource    string
		bulkUpdateReq BulkUpdateDocumentsByQueryRequest
		richTextAs    RichTextRepresentation
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ac.bulkUpdateDocumentByQuery(tt.args.dataSource, tt.args.bulkUpdateReq, tt.args.richTextAs)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.bulkUpdateDocumentByQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.bulkUpdateDocumentByQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_bulkDeleteDocumentsByQuery(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource string
		docs       []DocumentInfo
		mode       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ac.bulkDeleteDocumentsByQuery(tt.args.dataSource, tt.args.docs, tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.bulkDeleteDocumentsByQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.bulkDeleteDocumentsByQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_bulkDeleteDocumentByUnid(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource string
		unids      []string
		mode       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			got, err := ac.bulkDeleteDocumentByUnid(tt.args.dataSource, tt.args.unids, tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.bulkDeleteDocumentByUnid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.bulkDeleteDocumentByUnid() = %v, want %v", got, tt.want)
			}
		})
	}
}
