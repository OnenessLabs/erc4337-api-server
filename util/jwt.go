package util

import (
	"encoding/base64"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"strings"
)

func JWTExtractQueryString(jwt, queryPath string) (res string) {
	if jwtSplits := strings.Split(jwt, "."); len(jwtSplits) < 2 {
		return
	} else {
		if jwtBody, err := base64.RawURLEncoding.DecodeString(jwtSplits[1]); err == nil {
			if obj, err := oj.Parse(jwtBody); err == nil {
				if x, _ := jp.ParseString(queryPath); x != nil {
					if ys := x.Get(obj); len(ys) > 0 {
						res, _ = ys[0].(string)
					}
				}
			}
		}
	}
	return
}

func JWTExtractRatioAuthAddress(jwt string) string {
	return JWTExtractQueryString(jwt, "$.*.authentication_factors.*.crypto_wallet_factor.crypto_wallet_address")
}
