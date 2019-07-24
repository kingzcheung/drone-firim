package drone_firim

import (
	"bytes"
	"encoding/json"
	"errors"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const host = "http://api.fir.im"

type Firim struct {
	AppType  string
	BundleId string
	ApiToken string
	host     string
	upload   *Upload
}

type Upload struct {
	File        string
	Name        string
	Version     string
	Build       string
	ReleaseType string
	Changelog   string
}

func NewUpload(file string, name string, version string, build string, releaseType string, changelog string) *Upload {
	return &Upload{File: file, Name: name, Version: version, Build: build, ReleaseType: releaseType, Changelog: changelog}
}

type AppsResponse struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Short string `json:"short"`
	Cert  struct {
		Icon struct {
			Key       string `json:"key"`
			Token     string `json:"token"`
			UploadURL string `json:"upload_url"`
		} `json:"icon"`
		Binary struct {
			Key       string `json:"key"`
			Token     string `json:"token"`
			UploadURL string `json:"upload_url"`
		} `json:"binary"`
	} `json:"cert"`
}

func NewFirim(appType string, bundleId string, apiToken string, file string, name string, version string, build string, releaseType string, changelog string) *Firim {
	appType = strings.ToLower(appType)

	if appType != "ios" && appType != "android" {
		appType = "android"
	}

	return &Firim{
		AppType:  appType,
		BundleId: bundleId,
		ApiToken: apiToken,
		host:     host,
		upload:   NewUpload(file, name, version, build, releaseType, changelog),
	}
}

func (f *Firim) getToken() (AppsResponse, error) {
	cli := gentleman.New()
	cli.URL(f.host)
	req := cli.Request()
	req.Path("/apps")
	data := map[string]string{
		"type":      f.AppType,
		"bundle_id": f.BundleId,
		"api_token": f.ApiToken,
	}
	req.Method("POST")
	req.Use(body.JSON(data))
	res, err := req.Send()
	if err != nil {
		return AppsResponse{}, err
	}
	var apps AppsResponse
	err = json.Unmarshal(res.Bytes(), &apps)
	if err != nil {
		return AppsResponse{}, err
	}
	return apps, nil
}

func (f *Firim) check() error {
	if f.ApiToken == "" {
		return errors.New("api_token can not nil")
	}
	if f.BundleId == "" {
		return errors.New("bundle_id can not nil")
	}
	if f.upload.Name == "" {
		return errors.New("name can not nil")
	}

	if f.upload.Build == "" {
		return errors.New("build cant nil")
	}
	if f.upload.Version == "" {
		return errors.New("version cant nil")
	}
	if f.upload.File == "" {
		return errors.New("file cant nil")
	}
	return nil
}

func (f *Firim) Exec() error {

	if err := f.check(); nil != err {
		return err
	}

	apps, err := f.getToken()
	if err != nil {
		return err
	}

	extraParams := map[string]string{
		"key":       apps.Cert.Binary.Key,
		"token":     apps.Cert.Binary.Token,
		"x:name":    f.upload.Name,
		"x:version": f.upload.Version,
		"x:build":   f.upload.Build,
	}

	if f.upload.Changelog != "" {
		extraParams["x:changelog"] = f.upload.Changelog
	}

	if f.upload.ReleaseType != "" {
		extraParams["x:release_type"] = f.upload.ReleaseType
	}

	request, err := fileUploadRequest(apps.Cert.Binary.UploadURL, extraParams, "file", f.upload.File)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func fileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, buffer)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
