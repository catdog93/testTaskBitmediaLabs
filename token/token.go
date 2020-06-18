package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"os"
	"time"
)

func CreateToken(userid uint64) (*TokenDetails, error) {
	tokenDetails := &TokenDetails{}
	tokenDetails.AtExpires = time.Now().Add(time.Minute * 5).Unix()
	tokenDetails.AccessUuid = uuid.NewV4().String()

	tokenDetails.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.RefreshUuid = uuid.NewV4().String()

	//Creating Access Token
	err := os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	if err != nil {
		return nil, err
	}
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = tokenDetails.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = tokenDetails.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenDetails.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	err = os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	if err != nil {
		return nil, err
	}
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = tokenDetails.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = tokenDetails.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	tokenDetails.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return tokenDetails, nil
}
