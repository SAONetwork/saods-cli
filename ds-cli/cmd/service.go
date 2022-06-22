package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const AddFileEndPoint = "https://api.sao.network/saods/api/v1/file/addFile"
const GetFileEndPoint = "https://api.sao.network/saods/api/v1/file/"
const ListFilesEndPoint = "https://api.sao.network/sao-data-store/api/file/listFiles"

func AddFile(c *cli.Context) error {
	config = Config{
		appId:     c.String("appId"),
		apiKey:    c.String("apiKey"),
	}
	localPath := c.String("localPath")
	if localPath == "" {
		fmt.Println("no localPath passed")
		return nil
	}

	if config.appId == "" {
		fmt.Println("no appId passed")
		return nil
	}

	if config.apiKey == "" {
		fmt.Println("no apiKey passed")
		return nil
	}

	url := AddFileEndPoint
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(localPath)
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("file", filepath.Base(localPath))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return nil
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	basicAuth := base64.StdEncoding.EncodeToString([]byte(config.appId + ":" + config.apiKey))
	fmt.Println("Basic Authorization Credential:", basicAuth)
	req.Header.Add("Authorization", "Basic " + basicAuth)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(body))

	return nil
}

func GetFile(c *cli.Context) error {
	config = Config{
		appId:     c.String("appId"),
		apiKey:    c.String("apiKey"),
	}
	fileId := c.String("fileId")
	hash := c.String("hash")
	if fileId != "" && hash != "" {
		fmt.Println("please don't input both parameters: fileId and hash")
		return nil
	}

	if fileId == "" && hash == "" {
		fmt.Println("please input either one parameters: fileId and hash")
		return nil
	}

	if config.appId == "" {
		fmt.Println("no appId passed")
		return nil
	}

	if config.apiKey == "" {
		fmt.Println("no apiKey passed")
		return nil
	}

	url := GetFileEndPoint
	if fileId != "" {
		url = url + "by-id/" + fileId
	} else if hash != "" {
		url = url + "by-hash/" + hash
	} else {
		return nil
	}
	method := "GET"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}


	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	basicAuth := base64.StdEncoding.EncodeToString([]byte(config.appId + ":" + config.apiKey))
	fmt.Println("Basic Authorization Credential:", basicAuth)
	req.Header.Add("Authorization", "Basic " + basicAuth)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(body))
	return nil
}

func listFiles(c *cli.Context) error {
	config = Config{
		appId:     c.String("appId"),
		apiKey:    c.String("apiKey"),
	}
	page := c.String("page")

	if page == "" {
		page = "1"
	}


	url := ListFilesEndPoint + "?size=100&page=" + page
	method := "GET"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}


	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	basicAuth := base64.StdEncoding.EncodeToString([]byte(config.appId + ":" + config.apiKey))
	fmt.Println("Basic Authorization Credential:", basicAuth)
	req.Header.Add("Authorization", "Basic " + basicAuth)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(body))
	return nil
}