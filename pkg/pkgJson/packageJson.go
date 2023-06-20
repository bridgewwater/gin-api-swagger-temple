package pkgJson

import (
	"encoding/json"
	"fmt"
	"strings"
)

var (
	pkgJsonContent          string
	pkgVersionPrefixGoStyle = "v"
	pkgJson                 *PkgJson
)

func InitPkgJsonContent(content string) {
	pkgJsonContent = content
}

func SetVersionPrefixGoStyle(prefix string) {
	pkgVersionPrefixGoStyle = prefix
}

func GetPackageJsonVersionGoStyle() string {
	jsonVersion := GetPackageJsonVersion()
	if !strings.HasPrefix(jsonVersion, pkgVersionPrefixGoStyle) {
		return fmt.Sprintf("%s%s", pkgVersionPrefixGoStyle, jsonVersion)
	}
	return jsonVersion
}

func GetPackageJsonName() string {
	checkPackageJsonLoad()
	return pkgJson.Name
}

func GetPackageJsonVersion() string {
	checkPackageJsonLoad()
	return pkgJson.Version
}

func checkPackageJsonLoad() {
	if pkgJson == nil {
		if pkgJsonContent == "" {
			panic(fmt.Errorf("pkgJson must use pkgJson.InitPkgJsonContent(content) , then usex"))
		}
		pkgJ := PkgJson{}
		err := json.Unmarshal([]byte(pkgJsonContent), &pkgJ)
		if err != nil {
			panic(fmt.Errorf("pkgJson parse package.json err: %v", err))
		}
		pkgJson = &pkgJ
	}
}

type PkgJson struct {
	Name        string     `json:"name"`
	Version     string     `json:"version"`
	Description string     `json:"description"`
	Repository  Repository `json:"repository"`
	Keywords    []string   `json:"keywords"`
	Author      string     `json:"author"`
	License     string     `json:"license"`
	Homepage    string     `json:"homepage"`
}

type Repository struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}
