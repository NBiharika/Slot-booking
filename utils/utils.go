package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func ReadRequestBody(ctx *gin.Context) (map[string]interface{}, error) {
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(jsonData, &m)
	return m, err
}
