package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sao-datastore-cli/ds-cli/config"
	"strings"
)

const AddFileEndPoint = "https://api.sao.network/saods/api/v1/file/addFile"
const GetFileEndPoint = "https://api.sao.network/saods/api/v1/file/"
const ListFilesEndPoint = "https://api.sao.network/sao-data-store/api/file/listFiles"

func AddFile(c *cli.Context) error {
	loadConfig(c)
	localPath := c.String("localPath")
	if localPath == "" {
		fmt.Println("no localPath set")
		return nil
	}

	if cfg.AppId == "" {
		fmt.Println("no appId set")
		return nil
	}

	if cfg.ApiKey == "" {
		fmt.Println("no apiKey set")
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

	basicAuth := base64.StdEncoding.EncodeToString([]byte(cfg.AppId + ":" + cfg.ApiKey))
	req.Header.Add("Authorization", "Basic " + basicAuth)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode == 502 || res.StatusCode == 404 {
		fmt.Println("Unable to complete due to service error, please try it later")
		return nil
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	result := body
	if c.Bool("pretty") {
		result, err = formatJSON(body)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	fmt.Println(string(result))

	return nil
}

func GetFile(c *cli.Context) error {
	loadConfig(c)
	fileId := c.String("fileId")
	hash := c.String("hash")
	if fileId != "" && hash != "" {
		fmt.Println("please don't input both parameters: fileId and hash")
		return nil
	}

	if cfg.AppId == "" {
		fmt.Println("no appId set")
		return nil
	}

	if cfg.ApiKey == "" {
		fmt.Println("no apiKey set")
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
	basicAuth := base64.StdEncoding.EncodeToString([]byte(cfg.AppId + ":" + cfg.ApiKey))
	req.Header.Add("Authorization", "Basic " + basicAuth)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode == 502 || res.StatusCode == 404 {
		fmt.Println("Unable to complete due to service error, please try it later")
		return nil
	}

	getDispos := res.Header.Get("Content-Disposition")

	fileName := "SAO_FILE"
	if getDispos != "" {
		_, params, err := mime.ParseMediaType(getDispos)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fileName = params["filename"]
	}

	localPath := c.String("localPath")
	if localPath != "" {
		localPath = strings.TrimSuffix(localPath, "/")
		fileName = filepath.FromSlash(localPath + "/" + fileName)
	}
	fmt.Println("The file stored in " + fileName)

	out, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer out.Close()
	_, err = io.Copy(out, res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

func listFiles(c *cli.Context) error {
	loadConfig(c)
	page := c.String("page")

	if page == "" {
		page = "1"
	}

	size := c.String("size")

	if size == "" {
		size = "100"
	}

	url := ListFilesEndPoint + "?size=" + size + "&page=" + page
	method := "GET"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
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
	basicAuth := base64.StdEncoding.EncodeToString([]byte(cfg.AppId + ":" + cfg.ApiKey))
	req.Header.Add("Authorization", "Basic "+basicAuth)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode == 502 || res.StatusCode == 404 {
		fmt.Println("Unable to complete due to service error, please try it later")
		return nil
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	result := body
	if c.Bool("pretty") {
		result, err = formatJSON(body)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	fmt.Println(string(result))
	return nil
}

func loadConfig(c *cli.Context) {
	if c.String("appId") != "" || c.String("apiKey") != "" {
		cfg = config.Config{
			AppId:  c.String("appId"),
			ApiKey: c.String("apiKey"),
		}
	} else {
		cfg, _ = config.GetConfig()
	}
}

func setConfigFile(c *cli.Context) error {
	if c.String("appId") == "" || c.String("apiKey") == "" {
		fmt.Println("appId or apiKey is not properly set")
		return nil
	}
	cfg = config.Config{
		AppId:  c.String("appId"),
		ApiKey: c.String("apiKey"),
	}
	err := config.SetConfig(cfg)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

func formatJSON(data []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "    ")
	if err == nil {
		return out.Bytes(), err
	}
	return data, nil
}