/* ========================================================================== *
 * Copyright (C) 2023 HCL America Inc.                                        *
 * Apache-2.0 license   https://www.apache.org/licenses/LICENSE-2.0           *
 * ========================================================================== */

// Project : Keep Go SDK
// Author : Patrick Mark Garcia Mazo
// Role : Senior Software Engineer
package gosdk

type SessionConfig struct {
	AccessMethods    *AccessMethods
	ConnectorMethods *ConnectorMethods
}

type SessionMethods struct {
	DocumentOperationsMethods
	ScopeOperationsMethods
	ListViewOperationsMethods
}

func (s *SessionConfig) DominoUserSession() *SessionMethods {

	ac := new(AccessConnectorConfig)
	ac.AccessMethods = *s.AccessMethods
	ac.ConnectorMethods = *s.ConnectorMethods
	docOperations := ac.DocumentOperations()
	scopeOperations := ac.ScopeOperations()
	listViewOperations := ac.DominoListViewOperations()

	session := new(SessionMethods)
	session.DocumentOperationsMethods = *docOperations
	session.ScopeOperationsMethods = *scopeOperations
	session.ListViewOperationsMethods = *listViewOperations

	return session
}
