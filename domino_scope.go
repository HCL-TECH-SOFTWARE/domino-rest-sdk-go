/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package keep_sdk

import (
	"errors"
	"strings"

	h "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/helpers"
	t "github.com/HCL-TECH-SOFTWARE/domino-rest-sdk-go/keep/types"
)

var baseScope = new(BaseDominoScope)

type BaseDominoScope struct {
	Meta               t.DominoDocumentMeta `json:"@meta"`
	ApiName            string               `json:"apiName"`
	Description        string               `json:"description"`
	Icon               string               `json:"icon" default:"PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4NCjxzdmcgdmVyc2lvbj0iMS4xIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB4PSIwcHgiIHk9IjBweCIgdmlld0JveD0iMCAwIDUxMiA1MTIiIHN0eWxlPSJlbmFibGUtYmFja2dyb3VuZDpuZXcgMCAwIDUxMiA1MTI7IiB4bWw6c3BhY2U9InByZXNlcnZlIj4NCjxwb2x5Z29uIHN0eWxlPSJmaWxsOiM0Mzk4RDE7IiBwb2ludHM9IjQ2OS4zMzMsMzI0LjI2NyA0Mi42NjcsMzI0LjI2NyA0Mi42NjcsNTkuNzMzIDIzMC40LDU5LjczMyAyMzAuNCwyNS42IDQ2OS4zMzMsMjUuNiAiLz4NCjxyZWN0IHg9IjQyLjY2NyIgeT0iMTI4IiBzdHlsZT0iZmlsbDojREU0QzNDOyIgd2lkdGg9IjQyNi42NjciIGhlaWdodD0iMTQ1LjA2NyIvPg0KPHJlY3QgeD0iNTkuNzMzIiB5PSIyNS42IiBzdHlsZT0iZmlsbDojRkRCNjJGOyIgd2lkdGg9IjUxLjIiIGhlaWdodD0iMzQuMTMzIiAvPg0KPHBhdGggc3R5bGU9ImZpbGw6IzQzOThEMTsiIGQ9Ik04LjUzMywxNjIuMTMzYzAsOS40MjksNy42MzcsMTcuMDY3LDE3LjA2NywxNy4wNjdoMTcuMDY3di0zNC4xMzNIMjUuNiBDMTYuMTcxLDE0NS4wNjcsOC41MzMsMTUyLjcwNCw4LjUzMywxNjIuMTMzeiIgLz4NCjxwYXRoIHN0eWxlPSJmaWxsOiM0Mzk4RDE7IiBkPSJNNDg2LjQsMTQ1LjA2N2gtMTcuMDY3VjE3OS4ySDQ4Ni40YzkuNDI5LDAsMTcuMDY3LTcuNjM3LDE3LjA2Ny0xNy4wNjcgQzUwMy40NjcsMTUyLjcwNCw0OTUuODI5LDE0NS4wNjcsNDg2LjQsMTQ1LjA2N3oiIC8+DQo8cmVjdCB4PSIzODQiIHk9IjU5LjczMyIgc3R5bGU9ImZpbGw6Izg3Q0VEOTsiIHdpZHRoPSI1MS4yIiBoZWlnaHQ9IjM0LjEzMyIvPg0KPGNpcmNsZSBzdHlsZT0iZmlsbDojRTVFNUU1OyIgY3g9IjExOS40NjciIGN5PSIxMjgiIHI9IjE3LjA2NyIvPg0KPHJlY3QgeD0iMTQ1LjA2NyIgeT0iMjUuNiIgc3R5bGU9ImZpbGw6I0ZEQjYyRjsiIHdpZHRoPSI1MS4yIiBoZWlnaHQ9IjM0LjEzMyIvPg0KPHBhdGggc3R5bGU9ImZpbGw6IzQzOThEMTsiIGQ9Ik00MDkuNiwxODAuNzd2LTE4LjYzN2MwLTQuNzEtMy44MjMtOC41MzMtOC41MzMtOC41MzNzLTguNTMzLDMuODIzLTguNTMzLDguNTMzdjE4LjYzNyBjLTEzLjMyOSw0LjcxLTIwLjMxOCwxOS4zMzctMTUuNTk5LDMyLjY2NmM0LjcxLDEzLjMyOSwxOS4zMzcsMjAuMzE4LDMyLjY2NiwxNS41OTljMTMuMzI5LTQuNzE5LDIwLjMxOC0xOS4zMzcsMTUuNTk5LTMyLjY2NiBDNDIyLjYyMiwxODkuMDgyLDQxNi44ODcsMTgzLjM0Nyw0MDkuNiwxODAuNzd6IE00MDEuMDY3LDIxMy4zMzNjLTQuNzEsMC04LjUzMy0zLjgyMy04LjUzMy04LjUzM2MwLTQuNzEsMy44MjMtOC41MzMsOC41MzMtOC41MzMgczguNTMzLDMuODIzLDguNTMzLDguNTMzQzQwOS42LDIwOS41MSw0MDUuNzc3LDIxMy4zMzMsNDAxLjA2NywyMTMuMzMzeiIvPg0KPHJlY3QgeD0iMTYyLjEzMyIgeT0iODUuMzMzIiBzdHlsZT0iZmlsbDojODdDRUQ5OyIgd2lkdGg9IjI1LjYiIGhlaWdodD0iMTcuMDY3Ii8+DQo8cGF0aCBzdHlsZT0iZmlsbDojQ0ZDRkNGOyIgZD0iTTQwMS4wNjcsNDg2LjRIMzQuMTMzQzE1LjI4Myw0ODYuNCwwLDQ3MS4xMTcsMCw0NTIuMjY3VjE2Mi4xMzNjMC00LjcxLDMuODIzLTguNTMzLDguNTMzLTguNTMzIHM4LjUzMywzLjgyMyw4LjUzMyw4LjUzM3YyOTAuMTMzYzAsOS40MjksNy42MzcsMTcuMDY3LDE3LjA2NywxNy4wNjdoMzY2LjkzM2M5LjQyOSwwLDE3LjA2Ny03LjYzNywxNy4wNjctMTcuMDY3IHMtNy42MzctMTcuMDY3LTE3LjA2Ny0xNy4wNjdoLTIwNC44Yy0xOC44NSwwLTM0LjEzMy0xNS4yODMtMzQuMTMzLTM0LjEzM3MxNS4yODMtMzQuMTMzLDM0LjEzMy0zNC4xMzNoMjgxLjYgYzkuNDI5LDAsMTcuMDY3LTcuNjM3LDE3LjA2Ny0xNy4wNjdWMTYyLjEzM2MwLTQuNzEsMy44MjMtOC41MzMsOC41MzMtOC41MzNjNC43MSwwLDguNTMzLDMuODIzLDguNTMzLDguNTMzdjE4Ny43MzMgYzAsMTguODUtMTUuMjgzLDM0LjEzMy0zNC4xMzMsMzQuMTMzaC0yODEuNmMtOS40MjksMC0xNy4wNjcsNy42MzctMTcuMDY3LDE3LjA2N2MwLDkuNDI5LDcuNjM3LDE3LjA2NywxNy4wNjcsMTcuMDY3aDIwNC44IGMxOC44NSwwLDM0LjEzMywxNS4yODMsMzQuMTMzLDM0LjEzM1M0MTkuOTE3LDQ4Ni40LDQwMS4wNjcsNDg2LjR6Ii8+DQo8cGF0aCBzdHlsZT0iZmlsbDojNDM5OEQxOyIgZD0iTTE1My42LDE3MC42NjdjLTE0LjE0LDAtMjUuNiwxMS40Ni0yNS42LDI1LjZzMTEuNDYsMjUuNiwyNS42LDI1LjZoMjUuNnYtNTEuMkgxNTMuNnoiLz4NCjxjaXJjbGUgc3R5bGU9ImZpbGw6I0U1RTVFNTsiIGN4PSIyNTYiIGN5PSIxOTYuMjY3IiByPSI5My44NjciLz4NCjxjaXJjbGUgc3R5bGU9ImZpbGw6Izg3Q0VEOTsiIGN4PSIyNTYiIGN5PSIxOTYuMjY3IiByPSI1OS43MzMiLz4NCjxwYXRoIHN0eWxlPSJmaWxsOiM3MUM0RDE7IiBkPSJNMjgzLjQyNiwxNDMuMjQxbC04MC40NTIsODAuNDUyYzE1LjEwNCwyOS4zMjksNTEuMTIzLDQwLjg2Niw4MC40NTIsMjUuNzYyIHM0MC44NjYtNTEuMTIzLDI1Ljc2Mi04MC40NTJDMzAzLjQ4OCwxNTcuOTQzLDI5NC40ODUsMTQ4LjkzMiwyODMuNDI2LDE0My4yNDF6Ii8+DQo8L3N2Zz4NCg=="`
	IconName           string               `json:"iconName" default:"beach"`
	IsActive           bool                 `json:"isActive" default:"true"`
	MaximumAccessLevel string               `json:"maximumAccessLevel"`
	NSFPath            string               `json:"nsfPath"`
	SchemaName         string               `json:"schemaName"`
	Server             string               `json:"server"`
	Form               string               `json:"Form"`
	Type               string               `json:"Type"`
	UpdatedBy          []string             `json:"$UpdatedBy"`
	Revisions          string               `json:"$Revisions"`
}

func DominoScope(doc t.ScopeBody) error {

	if !h.HasFieldStructProps(doc, "apiName") && len(strings.Trim(doc.ApiName, "")) == 0 {
		return errors.New("Domino scope needs apiName value.")
	}

	if h.HasFieldStructProps(doc, "schemaName") && len(strings.Trim(doc.SchemaName, "")) == 0 {
		return errors.New("Domino scope needs schemaName value.")
	}

	if h.HasFieldStructProps(doc, "nsfPath") && len(strings.Trim(doc.NSFPath, "")) == 0 {
		return errors.New("Domino scope needs nsfPath value.")
	}

	baseScope.ApiName = doc.ApiName
	baseScope.SchemaName = doc.SchemaName
	baseScope.NSFPath = doc.NSFPath
	baseScope.Description = doc.Description
	baseScope.Icon = doc.Icon
	baseScope.IconName = doc.IconName
	baseScope.IsActive = doc.IsActive
	baseScope.MaximumAccessLevel = doc.MaximumAccessLevel
	baseScope.Server = doc.Server
	baseScope.Meta = doc.Meta
	baseScope.Form = doc.Form
	baseScope.Type = doc.Type
	baseScope.UpdatedBy = doc.UpdatedBy
	baseScope.Revisions = doc.Revisions

	return nil

}

func ToScopeJson() map[string]interface{} {

	scopeJson := map[string]interface{}{
		"apiName":            baseScope.ApiName,
		"description":        baseScope.Description,
		"icon":               baseScope.Icon,
		"iconName":           baseScope.IconName,
		"isActive":           baseScope.IsActive,
		"maximumAccessLevel": baseScope.MaximumAccessLevel,
		"nsfPath":            baseScope.NSFPath,
		"schemaName":         baseScope.SchemaName,
		"server":             baseScope.Server,
	}

	return scopeJson

}
