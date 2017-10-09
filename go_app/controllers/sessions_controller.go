package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goonr/gorails/session"
)

const (
	// here use `development` env secret_key_base from Rails' config/secret.yml
	secretKeyBase = "36bf5d0c4782351d14190f9188037459950778b650bc9efe64902c76c3c1bab1759d1f0e4e1e424e1f4e7a3c9da9687f61ef1bc5280460b4305440a101def62d"
	salt          = "encrypted cookie"        // default value for Rails 4 app
	signSalt      = "signed encrypted cookie" // default value for Rails 4 app
)

func ReadHandler(c *gin.Context) {
	// the session's key format: _<your rails app name>_session
	// in this example our app name is "example_read_rails_session", so the key is "_example_read_rails_session_session"
	sess, err := c.Request.Cookie("_example_read_rails_session_session")
	if err != nil {
		fmt.Printf("read cookie err: %v", err)
	}
	sessData, err := getRailsSessionData(sess.Value)
	if err != nil {
		fmt.Printf("deserialize err: %v", err)
	}
	var jsonData map[string]interface{}
	err = json.Unmarshal(sessData, &jsonData)
	if err != nil {
		fmt.Printf("json unmarshal err: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"data": jsonData})
}

func getRailsSessionData(sessionCookie string) (decryptedCookieData []byte, err error) {
	decryptedCookieData, err = session.DecryptSignedCookie(sessionCookie, secretKeyBase, salt, signSalt)
	return
}
