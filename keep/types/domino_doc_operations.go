/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package types

type RichTextRepresentation struct {
	HTML     string `json:"html" default:"html"`
	MIME     string `json:"mime" default:"mime"`
	MARKDOWN string `json:"markdown" default:"markdown"`
	PLAIN    string `json:"plain" default:"plain"`
}
type QueryActions struct {
	Execute string `json:"execute" default:"execute"`
	Explain string `json:"explain" default:"explain"`
	Parse   string `json:"parse" default:"parse"`
}
type DocumentOptions struct {
	ComputeWithForm bool   `json:"computeWithForm"`
	Meta            bool   `json:"meta"`
	RichTextAs      string `json:"richTextAs"`
	MarkRead        bool   `json:"markRead"`
	Mode            string `json:"mode"`
	ParentUnid      string `json:"parentUnid"`
	Revision        string `json:"revision"`
}

type GetDocumentOptions struct {
	ComputeWithForm bool   `json:"computeWithForm"`
	Meta            bool   `json:"meta"`
	RichTextAs      string `json:"richTextAs"`
	MarkRead        bool   `json:"markRead"`
	Mode            string `json:"mode"`
	ParentUnid      string `json:"parentUnid"`
}

type CreateDocumentOptions struct {
	RichTextAs string `json:"richTextAs"`
	ParentUnid string `json:"parentUnid"`
}

type UpdateDocumentOptions struct {
	RichTextAs string `json:"richTextAs"`
	Mode       string `json:"mode"`
	ParentUnid string `json:"parentUnid"`
	Revision   string `json:"revision"`
}

type PatchDocumentOptions struct {
	RichTextAs string `json:"richTextAs"`
	Mode       string `json:"mode"`
	ParentUnid string `json:"parentUnid"`
	Revision   string `json:"revision"`
}

type BulkGetDocumentsOptions struct {
	Meta       bool   `json:"meta"`
	RichTextAs string `json:"richTextAs"`
}

type BulkCreateDocumentsOptions struct {
	RichTextAs string `json:"richTextAs"`
}

type QueryOptions struct {
	MaxScanDocs           int                    `json:"maxScanDocs"`
	MaxScanEntries        int                    `json:"maxScanEntries"`
	Mode                  string                 `json:"mode"`
	NoViews               bool                   `json:"noViews"`
	Query                 string                 `json:"query"`
	ReplaceItems          map[string]interface{} `json:"replaceItems"`
	Variables             map[string]interface{} `json:"variables"`
	ReturnUpdatedDocument bool                   `json:"returnUpdatedDocument"`
	TimeoutSecs           int                    `json:"timeoutSecs"`
	ViewRefresh           bool                   `json:"viewRefresh"`
	Start                 int                    `json:"start"`
	Count                 int                    `json:"count"`
	Forms                 []string               `json:"forms"`
	IncludeFormAlias      bool                   `json:"includeFormAlias"`
	MarkRead              bool                   `json:"markRead"`
}

type BulkUpdateDocumentsByQueryRequest struct {
	MaxScanDocs           int                    `json:"maxScanDocs"`
	MaxScanEntries        int                    `json:"maxScanEntries"`
	Mode                  string                 `json:"mode"`
	NoViews               bool                   `json:"noViews"`
	Query                 string                 `json:"query"`
	ReplaceItems          map[string]interface{} `json:"replaceItems"`
	Variables             map[string]interface{} `json:"variables"`
	ReturnUpdatedDocument bool                   `json:"returnUpdatedDocument"`
	TimeoutSecs           int                    `json:"timeoutSecs"`
	ViewRefresh           bool                   `json:"viewRefresh"`
	Start                 int                    `json:"start"`
	Count                 int                    `json:"count"`
}

type GetDocumentsByQueryRequest struct {
	MaxScanDocs      int                    `json:"maxScanDocs"`
	MaxScanEntries   int                    `json:"maxScanEntries"`
	Mode             string                 `json:"mode"`
	NoViews          bool                   `json:"noViews"`
	Query            string                 `json:"query"`
	ReplaceItems     map[string]interface{} `json:"replaceItems"`
	Variables        map[string]interface{} `json:"variables"`
	TimeoutSecs      int                    `json:"timeoutSecs"`
	ViewRefresh      bool                   `json:"viewRefresh"`
	Start            int                    `json:"start"`
	Count            int                    `json:"count"`
	Forms            []string               `json:"forms"`
	IncludeFormAlias bool                   `json:"includeFormAlias"`
	MarkRead         bool                   `json:"markRead"`
}

type GetDocumentsByQueryOptions struct {
	Count      int    `json:"count"`
	RichTextAs string `json:"richTextAs"`
	Start      int    `json:"start"`
}

type DocumentStatusResponse struct {
	StatusText string `json:"statusText"`
	Status     int    `json:"status"`
	Message    string `json:"message"`
	UNID       string `json:"unid"`
}

type QueryDocumentExplainResponse struct {
	ExplainResult string `json:"explainResult"`
}

type QueryDocumentParseResponse struct {
	ParseResult string `json:"parseResult"`
}
