package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	STATUS_INITIALIZING = "Initializing"
	STATUS_START        = "Starting"
	STATUS_DOWNLOADING  = "Downloading from S3"
	STATUS_EXTRACTING   = "Extracting"
	STATUS_EXECUTING    = "Executing deployment"
	STATUS_FAILED       = "Failed - "
	STATUS_BOOTING      = "Booting"
	STATUS_PRE_RUNNING  = "Pre Running"
	STATUS_RUNNING      = "Running"
	STATUS_DELETING     = "Deleting"
	STATUS_DELETED      = "Deleted"
	STATUS_DOWN         = "Down"
)

const (
	STATUS_RETRY           = -1
	STATUS_INITIALIZING_ID = 0
	STATUS_START_ID        = 1
	STATUS_DOWNLOADING_ID  = 2
	STATUS_EXTRACTING_ID   = 3
	STATUS_EXECUTING_ID    = 4
	STATUS_BOOTING_ID      = 5
	STATUS_RUNNING_ID      = 6
	STATUS_FAILED_ID       = 7
	STATUS_DELETING_ID     = 8
	STATUS_DELETED_ID      = 9
	STATUS_DOWN_ID         = 10
)

func StrPtr(s string) *string {
	return &s
}

func Int64Ptr(i int) *int64 {
	r := int64(i)
	return &r
}

func IntPtr(i int) *int {
	return &i
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func BoolPtr(f bool) *bool {
	return &f
}

func AppendQueryToURL(url string, querys map[string]string) string {
	var sb strings.Builder
	for k, v := range querys {
		sb.WriteString(k + "=" + v + "&")
	}
	return url + "?" + sb.String()
}

func GetRequest(url string, authHeader string, accessToken string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(authHeader, accessToken)
	req.Header.Set("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func PostRequest(url string, b interface{}, authHeader string, accessToken string) (*http.Response, error) {
	jsonByte, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonByte))
	if err != nil {
		return nil, err
	}
	if accessToken != "" {
		req.Header.Set(authHeader, accessToken)
	}
	req.Header.Set("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 201 && res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, errors.New(fmt.Sprint(string(body)))
	}
	return res, nil
}

func DeleteRequest(url string, authHeader string, accessToken string) error {
	var req *http.Request
	var err error
	req, err = http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	req.Header.Set(authHeader, accessToken)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 204 && res.StatusCode != 404 {
		body, _ := ioutil.ReadAll(res.Body)
		return errors.New(fmt.Sprint(string(body)))
	}
	return nil
}

func PutRequest(url string, b interface{}, authHeader string, accessToken string) (*http.Response, error) {
	jsonByte, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonByte))
	if err != nil {
		return nil, err
	}
	req.Header.Set(authHeader, accessToken)
	req.Header.Set("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, errors.New(fmt.Sprint(string(body)))
	}
	return res, nil
}

func CheckFileExist(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) || info.IsDir() {
		return false
	}
	return true
}

func ReadTextFile(file string) (*string, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	res := string(dat)
	return &res, nil
}

func GetPublicKeyName(keyContent string) (*string, error) {
	sections := strings.Split(keyContent, " ")
	if len(sections) != 3 {
		return nil, errors.New("Given public key has no name section.")
	}
	return &sections[2], nil
}
