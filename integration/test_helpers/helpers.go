package test_helpers

import (
	"bufio"
	"encoding/json"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"

	. "github.com/onsi/gomega/gexec"
)

func CompareExpectedToGeneratedPo(expectedFilePath string, generatedFilePath string) {
	expectedTranslation := ReadPo(expectedFilePath)
	generatedTranslation := ReadPo(generatedFilePath)

	Ω(reflect.DeepEqual(expectedTranslation, generatedTranslation)).Should(BeTrue())
}

func CompareExpectedToGeneratedTraslationJson(expectedFilePath string, generatedFilePath string) {
	expectedTranslation := ReadJson(expectedFilePath)
	generatedTranslation := ReadJson(generatedFilePath)

	Ω(reflect.DeepEqual(expectedTranslation, generatedTranslation)).Should(BeTrue())
}

func CompareExpectedToGeneratedExtendedJson(expectedFilePath string, generatedFilePath string) {
	expectedTranslation := ReadJsonExtended(expectedFilePath)
	generatedTranslation := ReadJsonExtended(generatedFilePath)

	Ω(reflect.DeepEqual(expectedTranslation, generatedTranslation)).Should(BeTrue())
}

func GetFilePath(input_dir string, fileName string) string {
	return filepath.Join(os.Getenv("PWD"), input_dir, fileName)
}

func RemoveAllFiles(args ...string) {
	for _, arg := range args {
		os.Remove(arg)
	}
}

func Runi18n(args ...string) *Session {
	session := RunCommand(gi18nExec, args...)
	return session
}

func RunCommand(cmd string, args ...string) *Session {
	command := exec.Command(cmd, args...)
	session, err := Start(command, GinkgoWriter, GinkgoWriter)
	Ω(err).ShouldNot(HaveOccurred())
	session.Wait()
	return session
}

func ReadPo(fileName string) map[string]string {
	file, _ := os.Open(fileName)
	r := bufio.NewReader(file)

	myMap := make(map[string]string)
	for rawLine, _, err := r.ReadLine(); err != io.EOF; rawLine, _, err = r.ReadLine() {
		if err != nil {
			Fail(fmt.Sprintf("Error: %v", err))
		}

		line := string(rawLine)
		if strings.HasPrefix(line, "msgid") {
			rawLine, _, err = r.ReadLine()
			if err != nil {
				Fail(fmt.Sprintf("Error: %v", err))
			}

			myMap[line] = string(rawLine)
		}
	}

	return myMap
}

func ReadJson(fileName string) map[string]string {
	fileByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		Fail("Cannot open json file:" + fileName)
	}

	var b interface{}

	if err := json.Unmarshal(fileByte, &b); err != nil {
		Fail(fmt.Sprintf("Cannot unmarshal: %v", err))
	}

	myMap := make(map[string]string)

	for _, value := range b.([]interface{}) {
		valueMap := value.(map[string]interface{})
		myMap[valueMap["id"].(string)] = valueMap["translation"].(string)
	}

	return myMap
}

func ReadJsonExtended(fileName string) map[string]map[string]string {
	fileByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		Fail("Cannot open json file:" + fileName)
	}

	var b interface{}

	if err := json.Unmarshal(fileByte, &b); err != nil {
		Fail(fmt.Sprintf("Cannot unmarshal: %v", err))
	}

	myMap := make(map[string]map[string]string)

	for _, value := range b.([]interface{}) {
		valueMap := value.(map[string]interface{})

		dataMap := make(map[string]string)

		for key, val := range valueMap {
			switch val.(type) {
			case string:
				dataMap[key] = val.(string)
			case float64:
				dataMap[key] = fmt.Sprintf("%v", int(val.(float64)))
			default:
				fmt.Println("We did something wrong", key)
			}
		}

		myMap[valueMap["value"].(string)] = dataMap
	}

	return myMap
}