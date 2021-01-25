package util

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/url"
	"sort"
	"time"

	"strings"
)

func Md5(content string) string {
	log.Debug("signContent:", content)
	w := md5.New()
	io.WriteString(w, content)
	return fmt.Sprintf("%x", w.Sum(nil))
}

/**
参数签名算法
*/
func GenSign(params map[string]interface{}, signKey string) (string, string) {
	params["sign_key"] = signKey
	keys := []string{}
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	kvPairs := []string{}
	kvPairsQuery := []string{}

	for _, v := range keys {
		if v == "sign_key" {
			kvPairs = append(kvPairs, fmt.Sprintf("%s=%s", v, strings.TrimSpace(fmt.Sprintf("%v", params[v]))))
		} else {
			kvPairsQuery = append(kvPairsQuery, fmt.Sprintf("%s=%s", v, url.QueryEscape(strings.TrimSpace(fmt.Sprintf("%v", params[v])))))
			kvPairs = append(kvPairs, fmt.Sprintf("%s=%s", v, strings.TrimSpace(fmt.Sprintf("%v", params[v]))))

		}
	}
	sign := Md5(strings.Join(kvPairs, "&"))
	kvPairsQuery = append(kvPairsQuery, "sign="+sign)
	delete(params, "sign_key")
	return sign, strings.Join(kvPairsQuery, "&")
}

func SignRequest(data interface{}, signKey string) []byte {
	b, _ := json.Marshal(data)
	params := map[string]interface{}{}
	json.Unmarshal(b, &params)
	params["timestamp"] = time.Now().Unix()
	params["sign"], _ = GenSign(params, signKey)
	r, _ := json.Marshal(params)
	log.Debug("签名结果:", string(r))
	return r
}

func SignGetRequest(data interface{}, signKey string) string {
	b, _ := json.Marshal(data)
	params := map[string]interface{}{}
	json.Unmarshal(b, &params)
	params["timestamp"] = time.Now().Unix()
	_, query := GenSign(params, signKey)
	return query
}
