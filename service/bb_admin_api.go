package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bensema/goadmin/model"
	"github.com/bensema/goadmin/utils"
	"github.com/bensema/library/biz"
	"github.com/bensema/library/crypto"
	"github.com/bensema/library/log"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func (s *Service) FindAdvertisePage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AdvertisePages)
}

func (s *Service) AdvertiseAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AdvertiseAdd)
}

func (s *Service) AdvertiseDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AdvertiseDel)
}

func (s *Service) AdvertiseQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AdvertiseQuery)
}

func (s *Service) AdvertiseUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AdvertiseUpdate)
}

func (s *Service) FindAnnouncementPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AnnouncementPages)
}

func (s *Service) AnnouncementAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AnnouncementAdd)
}

func (s *Service) AnnouncementDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AnnouncementDel)
}

func (s *Service) AnnouncementQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AnnouncementQuery)
}

func (s *Service) AnnouncementUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AnnouncementUpdate)
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

	fmt.Println(method)

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
	if err != nil {
		log.Errorf("bb error (%v)", err)
		return nil, err
	}
	switch method {
	case "POST":
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	case "GET":
		req.Header.Set("Content-Type", "application/json")
	}
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("bb response error (%v)", err)
		return nil, err
	}

	bodyEncode, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	body, err := crypto.AesCBCDecrypt(bodyEncode, []byte(aesKey), []byte(aesIv))
	if err != nil {
		log.Errorf("bb AesCBCDecrypt error (%v)", err)
		return nil, err
	}
	err = json.Unmarshal(body, &reply)
	if err != nil {
		log.Errorf("bb parse error (%v)", err)
		return nil, err
	}
	return reply, nil
}

func (s *Service) UploadImg(c *gin.Context) (fName string, err error) {
	f, err := c.FormFile("file")
	if err != nil {
		return
	}

	fileExt := strings.ToLower(path.Ext(f.Filename))
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
		err = errors.New("上传失败!只允许png,jpg,gif,jpeg文件")
		return
	}

	fileName := fmt.Sprintf("%s%s", utils.RandomString(10), time.Now().String())
	fildDir := fmt.Sprintf("%s%d%s/", s.conf.Upload.Path, time.Now().Year(), time.Now().Month().String())
	isExist, _ := utils.PathExists(fildDir)
	if !isExist {
		err = os.Mkdir(fildDir, os.ModePerm)
	}
	filepath := fmt.Sprintf("%s%s%s", fildDir, fileName, fileExt)
	err = c.SaveUploadedFile(f, filepath)
	fName = fildDir + fileName
	return
}
