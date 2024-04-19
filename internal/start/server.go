package start

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oneness/erc-4337-api/config"
	"github.com/oneness/erc-4337-api/erc4337"
	"github.com/oneness/erc-4337-api/util"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

var db = make(map[string]string)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func setupRouter(hc *erc4337.HandlerContext) *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	// health test
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	erc4337Group := r.Group("erc4337")
	erc4337Group.GET("sender-info", hc.HandleGetSenderInfo)
	erc4337Group.GET("sender-address", hc.HandleGetSenderAddress)

	erc4337Group.GET("userop/approve", hc.HandleUserOpApprove)
	erc4337Group.GET("userop/withdrawto", hc.HandleUserOpWithdrawTo)
	erc4337Group.GET("userop/transfer", hc.HandleUserOpTransfer)

	erc4337Group.POST("userop/send", hc.HandleUserOpSend)

	return r
}

func Server() {
	// Read in from .env file if available
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			// Can ignore
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	// Read in from environment variables
	_ = viper.BindEnv("ERC4337_API_BUNDLER_URL")
	_ = viper.BindEnv("ERC4337_API_PAYMASTER_URL")
	_ = viper.BindEnv("ERC4337_API_ETH_CLIENT_SK")
	_ = viper.BindEnv("ERC4337_API_ETH_CLIENT_URL")

	maybeEnvUrl := viper.GetString("ERC4337_API_ETH_CLIENT_URL")

	// TODO: some API's will fail without these url's - should we just fail here...?
	maybeSUNodeUrl := viper.GetString("ERC4337_API_BUNDLER_URL")
	maybeSUPaymasterUrl := viper.GetString("ERC4337_API_PAYMASTER_URL")
	cfg := config.Config{
		ChainSKHex:     viper.GetString("ERC4337_API_ETH_CLIENT_SK"),
		ChainRpcUrl:    maybeEnvUrl,
		SUNodeUrl:      maybeSUNodeUrl,
		SUPayMasterUrl: maybeSUPaymasterUrl,
	}
	hc, err := erc4337.MakeContext(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := setupRouter(hc)
	localIp := util.GetOutboundIP()
	println(fmt.Sprintf("server starting at local IP %v", localIp.String())) // TODO: logging...
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
