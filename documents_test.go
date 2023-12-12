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

func TestDominoDocument(t *testing.T) {
	type args struct {
		doc map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *DocumentMethods
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DominoDocument(tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("DominoDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DominoDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocumentInfo_toDocumentJSON(t *testing.T) {
	type fields struct {
		Meta     DocumentMeta
		Form     string
		Warnings []string
		Fields   map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DocumentInfo{
				Meta:     tt.fields.Meta,
				Form:     tt.fields.Form,
				Warnings: tt.fields.Warnings,
				Fields:   tt.fields.Fields,
			}
			if got := d.toDocumentJSON(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DocumentInfo.toDocumentJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocumentInfo_getUNID(t *testing.T) {
	type fields struct {
		Meta     DocumentMeta
		Form     string
		Warnings []string
		Fields   map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DocumentInfo{
				Meta:     tt.fields.Meta,
				Form:     tt.fields.Form,
				Warnings: tt.fields.Warnings,
				Fields:   tt.fields.Fields,
			}
			if got := d.getUNID(); got != tt.want {
				t.Errorf("DocumentInfo.getUNID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDocumentInfo_setUNID(t *testing.T) {
	type fields struct {
		Meta     DocumentMeta
		Form     string
		Warnings []string
		Fields   map[string]interface{}
	}
	type args struct {
		unid string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DocumentInfo{
				Meta:     tt.fields.Meta,
				Form:     tt.fields.Form,
				Warnings: tt.fields.Warnings,
				Fields:   tt.fields.Fields,
			}
			d.setUNID(tt.args.unid)
		})
	}
}

func TestDocumentInfo_getRevision(t *testing.T) {
	type fields struct {
		Meta     DocumentMeta
		Form     string
		Warnings []string
		Fields   map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DocumentInfo{
				Meta:     tt.fields.Meta,
				Form:     tt.fields.Form,
				Warnings: tt.fields.Warnings,
				Fields:   tt.fields.Fields,
			}
			if got := d.getRevision(); got != tt.want {
				t.Errorf("DocumentInfo.getRevision() = %v, want %v", got, tt.want)
			}
		})
	}
}
