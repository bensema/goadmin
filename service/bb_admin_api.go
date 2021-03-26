package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bensema/goadmin/model"
	"github.com/bensema/goadmin/utils"
	"github.com/bensema/library/biz"
	"github.com/bensema/library/crypto"
	"github.com/bensema/library/log"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

func (s *Service) FindAdvertisePage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AdvertisePages)
}

func (s *Service) doRequest(c *gin.Context, cmd biz.AdminCmd) (reply *model.AdminApiReply, err error) {
	reply = new(model.AdminApiReply)

	aesKey := s.conf.AdminApi.AesKey
	aesIv := s.conf.AdminApi.AesIv
	method := c.Request.Method

	operatorInfo, err := s.GetAdminFromContext(c)
	if err != nil {
		return nil, err
	}
	var data string

	switch method {
	case "POST":
		d, _ := c.GetRawData()
		data = string(d)
	case "GET":
		data = c.Request.URL.RawQuery
	}

	params := biz.AdminRequest{
		OperatorId: operatorInfo.AdminId,
		Operator:   operatorInfo.Name,
		Ip:         c.ClientIP(),
		T:          time.Now().Unix(),
		R:          utils.RandomString(8),
		Cmd:        cmd,
		Method:     method,
		Data:       data,
	}

	b, err := json.Marshal(params)

	bf, err := crypto.AesCBCEncrypt(b, []byte(aesKey), []byte(aesIv))
	if err != nil {
		log.Errorf("bb AesCBCEncrypt error (%v)", err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, s.conf.AdminApi.Url, bytes.NewReader(bf))

	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("bb response error (%v)", err)
		return nil, err
	}

	bodyEncode, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(aesKey)
	fmt.Println(aesIv)
	body, err := crypto.AesCBCDecrypt(bodyEncode, []byte(aesKey), []byte(aesIv))
	if err != nil {
		log.Errorf("bb AesCBCDecrypt error (%v)", err)
		return nil, err
	}

	res := new(model.AdminApiReply)
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Errorf("bb parse error (%v)", err)
		return nil, err
	}
	return reply, nil
}
