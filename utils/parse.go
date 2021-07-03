package utils

import (
	"fmt"
	"strings"
)

var qu = `ZWNobyAtZSAicXVlcnlccmV4aXRcciJ8amF2YSAtY3AgZ3JleS5qYXIgTWFpbkFwcA==`
var tr = `ZWNobyAtZSAidHJhbnNmZXJcciVzXHIlZFxybm9cciJ8amF2YSAtY3AgZ3JleS5qYXIgTWFpbkFwcA==`

// Query 基础查询
func Query() ([]string, error) {
	decode, err := Decode(qu)
	if err != nil {
		return nil, err
	}

	command, err := Command(decode)
	if err != nil {
		return nil, err
	}

	return strings.Split(command, "\n"), err
}

// Transfer 转账
func Transfer(to string, amount int) (string, error) {
	decode, err := Decode(tr)
	if err != nil {
		return "", err
	}

	command, err := Command(fmt.Sprintf(decode, to, amount))
	if err != nil {
		return "", err
	}

	return command, err
}
