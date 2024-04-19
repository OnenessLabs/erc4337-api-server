package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJWTQuery(t *testing.T) {

	// {"aud":["project-test-f999d33e-49de-4889-bec6-8eb28044d047"],"exp":1689275480,"https://stytch.com/session":{"id":"session-test-46b25322-cefb-4695-ace4-a8e001ff2e49","started_at":"2023-07-13T19:06:20Z","last_accessed_at":"2023-07-13T19:06:20Z","expires_at":"2023-07-13T19:16:20Z","attributes":{"user_agent":"","ip_address":""},"authentication_factors":[{"type":"crypto","delivery_method":"crypto_wallet","last_authenticated_at":"2023-07-13T19:05:25Z","crypto_wallet_factor":{"crypto_wallet_id":"crypto-wallet-test-e3b11bab-9126-40c1-b28d-07d1395259f9","crypto_wallet_address":"0xE57bFE9F44b819898F47BF37E5AF72a0783e1141","crypto_wallet_type":"ethereum"}},{"type":"otp","delivery_method":"sms","last_authenticated_at":"2023-07-13T19:06:20Z","phone_number_factor":{"phone_id":"phone-number-test-10807dd7-c1d9-42b6-9c19-bbca053a787f","phone_number":"+15612023748"}}]},"iat":1689275180,"iss":"stytch.com/project-test-f999d33e-49de-4889-bec6-8eb28044d047","monitoring_session_id":"75c2a016-b313-4f21-9993-99b2d8c620ff","nbf":1689275180,"sub":"user-test-117b60f5-9d78-4354-84c9-f26d8a1ed9be"}
	testJwt := "eyJhbGciOiJSUzI1NiIsImtpZCI6Imp3ay10ZXN0LWI5NWU0NTM5LTNlZGMtNGJiNS04MjdmLTViMjAyY2NmNDJlNiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsicHJvamVjdC10ZXN0LWY5OTlkMzNlLTQ5ZGUtNDg4OS1iZWM2LThlYjI4MDQ0ZDA0NyJdLCJleHAiOjE2ODkyNzU0ODAsImh0dHBzOi8vc3R5dGNoLmNvbS9zZXNzaW9uIjp7ImlkIjoic2Vzc2lvbi10ZXN0LTQ2YjI1MzIyLWNlZmItNDY5NS1hY2U0LWE4ZTAwMWZmMmU0OSIsInN0YXJ0ZWRfYXQiOiIyMDIzLTA3LTEzVDE5OjA2OjIwWiIsImxhc3RfYWNjZXNzZWRfYXQiOiIyMDIzLTA3LTEzVDE5OjA2OjIwWiIsImV4cGlyZXNfYXQiOiIyMDIzLTA3LTEzVDE5OjE2OjIwWiIsImF0dHJpYnV0ZXMiOnsidXNlcl9hZ2VudCI6IiIsImlwX2FkZHJlc3MiOiIifSwiYXV0aGVudGljYXRpb25fZmFjdG9ycyI6W3sidHlwZSI6ImNyeXB0byIsImRlbGl2ZXJ5X21ldGhvZCI6ImNyeXB0b193YWxsZXQiLCJsYXN0X2F1dGhlbnRpY2F0ZWRfYXQiOiIyMDIzLTA3LTEzVDE5OjA1OjI1WiIsImNyeXB0b193YWxsZXRfZmFjdG9yIjp7ImNyeXB0b193YWxsZXRfaWQiOiJjcnlwdG8td2FsbGV0LXRlc3QtZTNiMTFiYWItOTEyNi00MGMxLWIyOGQtMDdkMTM5NTI1OWY5IiwiY3J5cHRvX3dhbGxldF9hZGRyZXNzIjoiMHhFNTdiRkU5RjQ0YjgxOTg5OEY0N0JGMzdFNUFGNzJhMDc4M2UxMTQxIiwiY3J5cHRvX3dhbGxldF90eXBlIjoiZXRoZXJldW0ifX0seyJ0eXBlIjoib3RwIiwiZGVsaXZlcnlfbWV0aG9kIjoic21zIiwibGFzdF9hdXRoZW50aWNhdGVkX2F0IjoiMjAyMy0wNy0xM1QxOTowNjoyMFoiLCJwaG9uZV9udW1iZXJfZmFjdG9yIjp7InBob25lX2lkIjoicGhvbmUtbnVtYmVyLXRlc3QtMTA4MDdkZDctYzFkOS00MmI2LTljMTktYmJjYTA1M2E3ODdmIiwicGhvbmVfbnVtYmVyIjoiKzE1NjEyMDIzNzQ4In19XX0sImlhdCI6MTY4OTI3NTE4MCwiaXNzIjoic3R5dGNoLmNvbS9wcm9qZWN0LXRlc3QtZjk5OWQzM2UtNDlkZS00ODg5LWJlYzYtOGViMjgwNDRkMDQ3IiwibW9uaXRvcmluZ19zZXNzaW9uX2lkIjoiNzVjMmEwMTYtYjMxMy00ZjIxLTk5OTMtOTliMmQ4YzYyMGZmIiwibmJmIjoxNjg5Mjc1MTgwLCJzdWIiOiJ1c2VyLXRlc3QtMTE3YjYwZjUtOWQ3OC00MzU0LTg0YzktZjI2ZDhhMWVkOWJlIn0.dG_P4VoWVVLgwrR_HQhSEzCYkxMTDRpB47spdvAyirx1K5EE7PGh0BZn0nIiDIgExFcOLdiedoP7sQoA08qRKPxvNGUj1SMGedOeOPnQxju6qa6zKuB4uNF-UIUn73-ZNuaNZRysPnk6Gp9e7aDm4w5bM2CISfvuDGby1s7ZkscYjARykUkB5fS36jH1OGC60hddID_xI_y8W4R_5wbcq3bObCpA8kKYQOIPY9LcJFx7AlRAzRIIYH8OGOIeyuJj2TLWRogct2LQSNu_OqLVqWjQ2jGsgIEAoE4I0vrJxFMc8t4F4uS1RHmXap5halQGXwB2-r6QmskbnA18v61r_A"

	checkISS := JWTExtractQueryString(testJwt, "iss")
	require.Equal(t, "stytch.com/project-test-f999d33e-49de-4889-bec6-8eb28044d047", checkISS)

	// TODO: better query?
	checkWalletAddr := JWTExtractQueryString(testJwt, "$.*.authentication_factors.*.crypto_wallet_factor.crypto_wallet_address")
	require.Equal(t, "0xE57bFE9F44b819898F47BF37E5AF72a0783e1141", checkWalletAddr)
}
