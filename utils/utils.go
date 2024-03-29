package utils

import (
	"library/ecode"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

func GetInt64(str string) int64 {
	d, _ := strconv.ParseInt(str, 10, 64)
	return d
}

func GetInt(str string) int {
	d, _ := strconv.Atoi(str)
	return d
}

func GetInt32(str string) int {
	d, _ := strconv.Atoi(str)
	return d
}

// 检查账户是否符合规范
func CheckNameLegal(name string) (err error) {
	var b bool
	if b, err = regexp.MatchString("^[a-zA-Z]\\w{5,17}$", name); err != nil {
		err = ecode.MemberNameFormatErr
		return
	}
	if b == false {
		err = ecode.MemberNameFormatErr
	}
	return
}

// 检查密码是否符合规范
func CheckPasswordLegal(password string) (err error) {
	err = Check(5, 20, LevelB, password)
	return
}

func CheckIn[T comparable](dis []T, k T) bool {
	for i, _ := range dis {
		if dis[i] == k {
			return true
		}
	}
	return false
}

func RandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func RandInt(n int) string {
	str := "0123456789"
	str19 := "123456789"

	bytes := []byte(str)
	bytes19 := []byte(str19)

	result := []byte{}

	for i := 0; i < n; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result[0] = bytes19[r.Intn(len(bytes19))]
	return string(result)
}

func S2IList(l []string) ([]int, error) {
	_l := make([]int, len(l))

	for i, v := range l {
		b, err := strconv.Atoi(v)
		if err != nil {
			return _l, err
		} else {
			_l[i] = b
		}
	}
	return _l, nil
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
