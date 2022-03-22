package common

import (
	"Antioxidant/server/model"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

/**
yaml工具集
*/

/*
ReadYaml
读取 yaml 文件
*/
func ReadYaml(filename string) (*model.RepoURLs, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var repos model.RepoURLs

	err = yaml.Unmarshal(buf, &repos)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}
	return &repos, nil
}
