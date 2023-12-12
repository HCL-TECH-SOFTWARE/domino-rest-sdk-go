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

func TestDominoListView(t *testing.T) {
	type args struct {
		doc ListViewBody
	}
	tests := []struct {
		name    string
		args    args
		want    *ListViewBody
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DominoListView(tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("DominoListView() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DominoListView() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateDesignColumnSimple(t *testing.T) {
	type args struct {
		obj DesignColumnSimple
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateDesignColumnSimple(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("validateDesignColumnSimple() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListViewBody_ToListViewJson(t *testing.T) {
	type fields struct {
		BaseListView     BaseListView
		Name             string
		SelectionFormula string
		Columns          []DesignColumnSimple
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := &ListViewBody{
				BaseListView:     tt.fields.BaseListView,
				Name:             tt.fields.Name,
				SelectionFormula: tt.fields.SelectionFormula,
				Columns:          tt.fields.Columns,
			}
			got, err := doc.ToListViewJson()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListViewBody.ToListViewJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListViewBody.ToListViewJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
