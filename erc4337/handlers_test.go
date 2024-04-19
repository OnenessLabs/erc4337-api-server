package erc4337

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestContextUserOp(t *testing.T) {

	mc, err := makeTestContext(map[string]string{
		"nonce":  "0",
		"sender": "deadbeef",
	})
	require.NoError(t, err)

	w := httptest.NewRecorder()
	engine := gin.New()
	testContext := gin.CreateTestContextOnly(w, engine)

	testContext.Request, _ = http.NewRequest(http.MethodGet, "/doesntmatter?owner=6D64a4aF99563a82B212124604f6d1759376F37F", nil)
	mc.HandleGetSenderInfo(testContext)

	println(w.Body.String())
}
