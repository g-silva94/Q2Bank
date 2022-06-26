package utils

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	allNumRe  = regexp.MustCompile("[^0-9]")
	allSame14 = regexp.MustCompile("0{14}|1{14}|2{14}|3{14}|4{14}|5{14}|6{14}|7{14}|8{14}|9{14}")
	allSame11 = regexp.MustCompile("0{11}|1{11}|2{11}|3{11}|4{11}|5{11}|6{11}|7{11}|8{11}|9{11}")
)

func RespondWithError(w http.ResponseWriter, code, errorCode int, message string) {
	RespondWithJSON(w, code, map[string]interface{}{"error": message, "code": errorCode})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithXML(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := xml.Marshal(payload)
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(code)
	w.Write(response)
}

func Encrypt(pass string) string {
	return encrypt(pass, randomSalt())
}

func Matches(clear string, encrypted string) bool {
	parts := strings.Split(encrypted, "$")
	if len(parts) != 3 {
		return false
	}
	return encrypted == encrypt(clear, parts[2])
}

func randomSalt() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var salt string
	for i := 0; i <= 5; i++ {
		n := r.Intn(100)
		salt = salt + strconv.Itoa(n)
	}
	return salt
}

func encrypt(pass string, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt+pass)
	password := fmt.Sprintf("sha1$%x$%s", h.Sum(nil), salt)
	return password
}

func IsValidEmail(email string) bool {
	reMail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(email) > 254 || !reMail.MatchString(email) {
		return false
	}
	return true
}

func IsValidCNPJ(cnpj string) bool {
	n := allNumRe.ReplaceAllString(cnpj, "")
	if n == "" {
		return false
	}
	if len(n) != 14 {
		return false
	}
	if allSame14.Match([]byte(n)) {
		return false
	}
	size := len(n) - 2
	numbers := n[0:size]
	digits := n[size:]
	var sum int
	pos := size - 7
	for i := size; i >= 1; i-- {
		num, _ := strconv.Atoi(string(numbers[size-i]))
		sum += num * pos
		pos = pos - 1
		if pos < 2 {
			pos = 9
		}
	}
	var result int
	if sum%11 < 2 {
		result = 0
	} else {
		result = 11 - sum%11
	}
	x, _ := strconv.Atoi(string(digits[0]))
	if result != x {
		return false
	}
	size = size + 1
	numbers = n[0:size]
	sum = 0
	pos = size - 7
	for i := size; i >= 1; i-- {
		num, _ := strconv.Atoi(string(numbers[size-i]))
		sum += num * pos
		pos = pos - 1
		if pos < 2 {
			pos = 9
		}
	}
	if sum%11 < 2 {
		result = 0
	} else {
		result = 11 - sum%11
	}
	num, _ := strconv.Atoi(string(digits[1]))
	if result != num {
		return false
	}
	return true
}

func IsValidCPF(cpf string) bool {
	cpf = allNumRe.ReplaceAllString(cpf, "")
	if cpf == "" {
		return false
	}
	if len(cpf) != 11 {
		return false
	}
	if allSame11.Match([]byte(cpf)) {
		return false
	}
	var sum int
	var res int
	for i := 1; i <= 9; i++ {
		num, _ := strconv.Atoi(cpf[i-1 : i])
		sum = sum + num*(11-i)
	}
	res = (sum * 10) % 11
	if (res == 10) || (res == 11) {
		res = 0
	}
	num, _ := strconv.Atoi(cpf[9:10])
	if res != num {
		return false
	}
	sum = 0
	for i := 1; i <= 10; i++ {
		num, _ := strconv.Atoi(cpf[i-1 : i])
		sum = sum + num*(12-i)
	}
	res = (sum * 10) % 11
	if (res == 10) || (res == 11) {
		res = 0
	}
	num, _ = strconv.Atoi(cpf[10:11])
	if res != num {
		return false
	}
	return true
}

func OnlyNumbers(s string) string {
	return allNumRe.ReplaceAllString(s, "")
}

func NormalizeCPFCNPJ(cpfcnpj string) (string, bool) {
	cpfcnpj = OnlyNumbers(cpfcnpj)
	if IsValidCNPJ(cpfcnpj) || IsValidCPF(cpfcnpj) {
		return cpfcnpj, true
	}

	if len(cpfcnpj) >= 11 {
		cpfcnpj = cpfcnpj[3:]
		if IsValidCPF(cpfcnpj) {
			return cpfcnpj, true
		}
	}

	return cpfcnpj, false
}

func InArray(s string, list []string) bool {
	for _, el := range list {
		if el == s {
			return true
		}
	}
	return false
}

func GetStringInBetween(str string, start string, end string) string {
	s := strings.Index(str, start)
	if s == -1 {
		return ""
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return ""
	}
	e += s
	if s > e {
		return ""
	}
	if len(end) == 0 {
		e = len(str)
	}
	return str[s:e]
}

func GetTimeNow() time.Time {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return time.Now()
	}
	now := time.Now().In(loc)
	return now
}

func LimpaString(s string) string {
	var (
		b           = bytes.NewBufferString("")
		acentuacoes = map[rune]string{'À': "A", 'Á': "A", 'Â': "A", 'Ã': "A", 'Ä': "A", 'Å': "AA", 'Æ': "AE", 'Ç': "C", 'È': "E", 'É': "E", 'Ê': "E", 'Ë': "E", 'Ì': "I", 'Í': "I", 'Î': "I", 'Ï': "I", 'Ð': "D", 'Ł': "L", 'Ñ': "N", 'Ò': "O", 'Ó': "O", 'Ô': "O", 'Õ': "O", 'Ö': "OE", 'Ø': "OE", 'Œ': "OE", 'Ù': "U", 'Ú': "U", 'Ü': "UE", 'Û': "U", 'Ý': "Y", 'Þ': "TH", 'ẞ': "SS", 'à': "a", 'á': "a", 'â': "a", 'ã': "a", 'ä': "ae", 'å': "aa", 'æ': "ae", 'ç': "c", 'è': "e", 'é': "e", 'ê': "e", 'ë': "e", 'ì': "i", 'í': "i", 'î': "i", 'ï': "i", 'ð': "d", 'ł': "l", 'ñ': "n", 'ń': "n", 'ò': "o", 'ó': "o", 'ô': "o", 'õ': "o", 'ō': "o", 'ö': "oe", 'ø': "oe", 'œ': "oe", 'ś': "s", 'ù': "u", 'ú': "u", 'û': "u", 'ū': "u", 'ü': "ue", 'ý': "y", 'ÿ': "y", 'ż': "z", 'þ': "th", 'ß': "ss"}
	)
	for _, c := range s {
		if val, ok := acentuacoes[c]; ok {
			b.WriteString(val)
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

func RemoveDuplicidadeLista(list []string) []string {
	check := make(map[string]int64, 0)

	if len(list) == 0 {
		return list
	}
	for i := range list {
		check[list[i]] = 1
	}
	list = []string{}
	for apelido := range check {
		list = append(list, apelido)
	}
	return list
}
