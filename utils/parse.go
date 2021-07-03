package utils

import (
	"strings"
)

// Query 基础查询
func Query() ([]string, error) {
	command, err := Command(`echo -e "query\rexit\n"|java -cp grey.jar MainApp`)
	if err != nil {
		return nil, err
	}

	return strings.Split(command, "\n"), err
}

// Transfer 转账
func Transfer(to string) ([]string, error) {
	command, err := Command(`echo -e "query\rexit\n"|java -cp grey.jar MainApp`)
	if err != nil {
		return nil, err
	}

	return strings.Split(command, "\n"), err
}
