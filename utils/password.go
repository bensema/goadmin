package utils

import (
	"fmt"
	"github.com/bensema/library/ecode"
	"regexp"
)

const (
	levelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)

func Check(minLength, maxLength, minLevel int, pwd string) error {
	if len(pwd) < minLength {
		return fmt.Errorf("密码最短长度为 %d 字符", minLength)
	}
	if len(pwd) > maxLength {
		return fmt.Errorf("密码最长长度为 %d 字符", maxLength)
	}

	var level int = levelD
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}

	if level < minLevel {
		return ecode.PasswordTooLeak
	}
	return nil

}
