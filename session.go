package gosdk

type SessionConfig struct {
	AccessMethods    *AccessMethods
	ConnectorMethods *ConnectorMethods
}

type SessionMethods struct {
	DocumentOperationsMethods
	ScopeOperationsMethods
}

func (s *SessionConfig) UserSession() *SessionMethods {

	ac := new(AccessConnectorConfig)
	ac.AccessMethods = *s.AccessMethods
	ac.ConnectorMethods = *s.ConnectorMethods
	docOperations := ac.DocumentOperations()
	scopeOperations := ac.ScopeOperations()

	session := new(SessionMethods)
	session.DocumentOperationsMethods = *docOperations
	session.ScopeOperationsMethods = *scopeOperations

	return session
}
