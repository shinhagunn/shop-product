package controllers

import "github.com/gofiber/fiber/v2"

var (
	ErrServerInternal = fiber.NewError(fiber.StatusInternalServerError, "server.internal_error")

	ErrServerInvalidQuery = fiber.NewError(fiber.StatusBadRequest, "server.method.invalid_message_query")

	ErrServerInvalidBody = fiber.NewError(fiber.StatusBadRequest, "server.method.invalid_message_body")

	ErrAuthzInvalidPermission = fiber.NewError(fiber.StatusUnauthorized, "authz.invalid_permission")

	ErrAuthzCsrfTokenMismatch = fiber.NewError(fiber.StatusUnauthorized, "authz.csrf_token_mismatch")

	ErrAuthzMissingCsrfToken = fiber.NewError(fiber.StatusUnauthorized, "authz.missing_csrf_token")

	ErrAuthzClientSessionMismatch = fiber.NewError(fiber.StatusUnauthorized, "authz.client_session_mismatch")

	ErrAuthzUserNotActive = fiber.NewError(fiber.StatusUnauthorized, "authz.user_not_active")

	ErrAuthzUserNotPending = fiber.NewError(fiber.StatusUnauthorized, "authz.user_not_pending")

	ErrAuthzUserNotGuest = fiber.NewError(fiber.StatusUnauthorized, "authz.user_not_guest")

	ErrAuthzUserNotExist = fiber.NewError(fiber.StatusUnauthorized, "authz.user_not_exist")

	ErrUnprocessableEntity = fiber.NewError(fiber.StatusUnprocessableEntity, "authz.unprocessable_entity")

	ErrAuthzInvalidSession = fiber.NewError(fiber.StatusUnauthorized, "authz.invalid_session")

	ErrAuthzPermissionDenied = fiber.NewError(fiber.StatusUnauthorized, "authz.permission_denied")

	ErrJWTDecodeAndVerify = fiber.NewError(fiber.StatusUnauthorized, "jwt.decode_and_verify")

	ErrMethodNotAllowed = fiber.NewError(fiber.StatusMethodNotAllowed, "server.method.not_allowed")

	ErrRecordNotFound = fiber.NewError(fiber.StatusNotFound, "record.not_found")
)
