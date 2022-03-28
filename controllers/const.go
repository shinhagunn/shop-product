package controllers

const (
	FailedToParseBody = "Failed to parse body"

	ServerInternalError = "Server internal error"

	FailedConnectDataInDatabase = "Failed to connect data in database"

	FailedToParseJWT = "Failed to parse jwt"

	FailedConnectToSessions = "Failed to connect to sessions"

	ServerInvalidQueryErr = "server.invalid_query"

	AuthzInvalidPermissionErr = "authz.invalid_permission"

	AuthzCsrfTokenMismatchErr = "authz.csrf_token_mismatch"

	AuthzMissingCsrfTokenErr = "authz.missing_csrf_token"

	AuthzClientSessionMismatchErr = "authz.client_session_mismatch"

	AuthzUserNotActiveErr = "authz.user_not_active"

	AuthzUserNotExistErr = "authz.user_not_exist"

	AuthzInvalidSessionErr = "authz.invalid_session"

	AuthZPermissionDeniedErr = "authz.permission_denied"

	RecordNotFoundErr = "record.not_found"
)
