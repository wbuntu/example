package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Client contains jwt and http.client
type Client struct {
	client http.Client
}

// SharedClient is public object
var SharedClient = NewClient()

// NewClient return an Client
func NewClient() *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 15 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		MaxConnsPerHost:       100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}
	c := &Client{
		client: client,
	}
	return c
}

// PushData is common http data push method
func (c *Client) PushData(object interface{}, url string, headers []byte) error {
	data, err := json.Marshal(object)
	if err != nil {
		return errors.Wrap(err, "Marshal object")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(data))
	if err != nil {
		return errors.Wrap(err, "Build Request")
	}
	req.Header.Set("Content-Type", "application/json")
	// json 括号 {} length == 2
	if len(headers) > 2 {
		var headerMap map[string]string
		if err := json.Unmarshal(headers, &headerMap); err == nil {
			for k, v := range headerMap {
				req.Header.Set(k, v)
			}
		}
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "Do Request")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "Read Resp Body")
		}
		// 限制返回的错误信息长度，避免大量打印
		if len(body) > 128 {
			body = body[:128]
		}
		return errors.Errorf("Code: %d Body: %s URL: %s", resp.StatusCode, string(body), url)
	}
	return nil
}

// DeviceDownlinkReq obj
type DeviceDownlinkReq struct {
	DevEUI    string `json:"devEUI"`
	Data      []byte `json:"data"`
	Confirmed bool   `json:"confirmed"`
	FPort     uint8  `json:"fPort"`
}

func main() {
	nowTime := time.Now()
	data := DeviceDownlinkReq{
		DevEUI:    "23336666",
		Data:      []byte{0x23, 0x23, 0x23, 0x23},
		Confirmed: true,
		FPort:     8,
	}
	err := SharedClient.PushData(data, "https://cloud.iotsquare.xyz/api/login", nil)
	log.Infof("Error: %s Duration: %s", err, time.Since(nowTime))
}
