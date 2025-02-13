/* ========================================================================== *
 * Copyright (C) 2023, 2025 HCL America Inc.                                  *
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

func TestAccessConnectorConfig_DominoListViewOperations(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	tests := []struct {
		name   string
		fields fields
		want   *ListViewOperationsMethods
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AccessConnectorConfig{
				AccessMethods:    tt.fields.AccessMethods,
				ConnectorMethods: tt.fields.ConnectorMethods,
			}
			if got := ac.DominoListViewOperations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.DominoListViewOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_getListViewEntry(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource   string
		listViewName string
		options      GetListViewEntryOptions
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
			got, err := ac.getListViewEntry(tt.args.dataSource, tt.args.listViewName, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.getListViewEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.getListViewEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_getListViewPivotEntry(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource   string
		listViewName string
		pivotColumn  string
		options      GetListPivotViewEntryOptions
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
			got, err := ac.getListViewPivotEntry(tt.args.dataSource, tt.args.listViewName, tt.args.pivotColumn, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.getListViewPivotEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.getListViewPivotEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_getListViews(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource string
		options    GetListViewOptions
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
			got, err := ac.getListViews(tt.args.dataSource, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.getListViews() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.getListViews() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_createUpdateListView(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource string
		listView   ListViewBody
		designName string
		options    CreateUpdateDesignOptions
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
			got, err := ac.createUpdateListView(tt.args.dataSource, tt.args.listView, tt.args.designName, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.createUpdateListView() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.createUpdateListView() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccessConnectorConfig_getListView(t *testing.T) {
	type fields struct {
		AccessMethods    AccessMethods
		ConnectorMethods ConnectorMethods
	}
	type args struct {
		dataSource string
		designName string
		options    GetDesignOptions
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
			got, err := ac.getListView(tt.args.dataSource, tt.args.designName, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccessConnectorConfig.getListView() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccessConnectorConfig.getListView() = %v, want %v", got, tt.want)
			}
		})
	}
}
