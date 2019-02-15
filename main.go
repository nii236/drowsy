package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

type InsomniaDefinition struct {
	Type         string    `json:"_type"`
	ExportFormat int       `json:"__export_format"`
	ExportDate   time.Time `json:"__export_date"`
	ExportSource string    `json:"__export_source"`
	Resources    []struct {
		ID          string      `json:"_id"`
		Created     int64       `json:"created"`
		Description string      `json:"description,omitempty"`
		Modified    int64       `json:"modified"`
		Name        string      `json:"name"`
		ParentID    interface{} `json:"parentId"`
		Type        string      `json:"_type"`
		Color       interface{} `json:"color,omitempty"`
		Data        struct {
		} `json:"data,omitempty"`
		IsPrivate   bool    `json:"isPrivate,omitempty"`
		MetaSortKey float64 `json:"metaSortKey,omitempty"`
		Cookies     []struct {
			Creation     time.Time `json:"creation"`
			Domain       string    `json:"domain"`
			HostOnly     bool      `json:"hostOnly"`
			HTTPOnly     bool      `json:"httpOnly"`
			ID           string    `json:"id"`
			Key          string    `json:"key"`
			LastAccessed time.Time `json:"lastAccessed"`
			Path         string    `json:"path"`
			Value        string    `json:"value"`
			Expires      time.Time `json:"expires,omitempty"`
			Secure       bool      `json:"secure,omitempty"`
		} `json:"cookies,omitempty"`
		Environment struct {
		} `json:"environment,omitempty"`
		Authentication struct {
		} `json:"authentication,omitempty"`
		Body struct {
		} `json:"body,omitempty"`
		Headers                         []interface{} `json:"headers,omitempty"`
		Method                          string        `json:"method,omitempty"`
		Parameters                      []interface{} `json:"parameters,omitempty"`
		SettingDisableRenderRequestBody bool          `json:"settingDisableRenderRequestBody,omitempty"`
		SettingEncodeURL                bool          `json:"settingEncodeUrl,omitempty"`
		SettingMaxTimelineDataSize      int           `json:"settingMaxTimelineDataSize,omitempty"`
		SettingRebuildPath              bool          `json:"settingRebuildPath,omitempty"`
		SettingSendCookies              bool          `json:"settingSendCookies,omitempty"`
		SettingStoreCookies             bool          `json:"settingStoreCookies,omitempty"`
		URL                             string        `json:"url,omitempty"`
	} `json:"resources"`
}

var file **os.File

func init() {
	file = kingpin.Flag("file", "File to load").Default("./insomnia.json").File()
}

func main() {
	kingpin.Parse()
	def := &InsomniaDefinition{}
	err := json.NewDecoder(*file).Decode(def)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, res := range def.Resources {
		// if res.Type == "request_group" {
		// 	fmt.Println(res.Name)
		// }
		if res.Type == "request" {
			fmt.Println(res.Method, res.URL)
		}

	}
}
