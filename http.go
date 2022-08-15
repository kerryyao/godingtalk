package godingtalk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

// UploadFile is for uploading a single file to DingTalk
type UploadFile struct {
	FieldName string
	FileName  string
	Reader    io.Reader
}

// DownloadFile is for downloading a single file from DingTalk
type DownloadFile struct {
	MediaID  string
	FileName string
	Reader   io.Reader
}

// http request
func httpRequest(path string, params url.Values, requestData interface{}) (*[]byte, error) {
	client := &http.Client{}

	var request *http.Request
	url := "https://" + conf.BaseURL + "/" + path + "?" + params.Encode()
	if requestData != nil {
		request, _ = http.NewRequest("GET", url, nil)
		goto DOIT
	}

	switch requestData.(type) {
	case UploadFile:
		var b bytes.Buffer
		w := multipart.NewWriter(&b)

		uploadFile := requestData.(UploadFile)
		if uploadFile.Reader == nil {
			return nil, errors.New("upload file is empty")
		}
		fw, err := w.CreateFormFile(uploadFile.FieldName, uploadFile.FileName)
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(fw, uploadFile.Reader); err != nil {
			return nil, err
		}
		if err = w.Close(); err != nil {
			return nil, err
		}
		request, _ = http.NewRequest("POST", url, &b)
		request.Header.Set("Content-Type", w.FormDataContentType())
	default:
		d, _ := json.Marshal(requestData)
		request, _ = http.NewRequest("POST", url, bytes.NewReader(d))
		request.Header.Set("Content-Type", "application/json")
	}

DOIT:
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Server error: " + resp.Status)
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &content, nil
}
