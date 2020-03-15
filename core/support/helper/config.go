package helper

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func Config(path string) (map[string]interface{}, error) {
	cnfFile, _ := filepath.Abs("./config/" + path)
	// 检测配置文件是否存在
	content, err := ioutil.ReadFile(cnfFile)
	if err != nil {
		fmt.Println("读取配置文件失败", cnfFile, err)
		return nil, err
	}

	data := make(map[string]interface{}, 0)
	// 读取配置信息&解析
	if err = yaml.Unmarshal(content, data); err != nil {
		return nil, err
	}

	return data, nil
}

func CoreConfig(path string) (map[string]interface{}, error) {
	cnfFile, _ := filepath.Abs("./core/config/" + path)
	fmt.Println("coreConfig是", cnfFile)
	return GetConfig(cnfFile)
}

func GetConfig(path string) (map[string]interface{}, error) {
	cnfFile, _ := filepath.Abs(path)
	// 检测配置文件是否存在
	content, err := ioutil.ReadFile(cnfFile)
	if err != nil {
		fmt.Println("读取配置文件失败", cnfFile, err)
		return nil, err
	}

	data := make(map[string]interface{}, 0)
	// 读取配置信息&解析
	if err = yaml.Unmarshal(content, data); err != nil {
		return nil, err
	}

	return data, nil
}
