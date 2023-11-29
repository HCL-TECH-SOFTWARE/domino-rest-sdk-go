/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package types

type DominoRestServer struct {
	BaseUrl            string
	AvailableAPIs      func() (map[string]interface{}, error)
	GetDominoConnector func(apiName string) (DominoRestConnector, error)
	Basis              func() (DominoRestConnector, error)
	Setup              func() (DominoRestConnector, error)
	Admin              func() (DominoRestConnector, error)
}

type ApiList struct {
	Admin DominoApiMeta `json:"admin"`
	Basis DominoApiMeta `json:"basis"`
	Pim   DominoApiMeta `json:"pim"`
	Poi   DominoApiMeta `json:"poi"`
	Setup DominoApiMeta `json:"setup"`
}

type DominoApiMeta struct {
	FileName  string `json:"fileName"`
	MountPath string `json:"mountPath"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Version   string `json:"version"`
}
