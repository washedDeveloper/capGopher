package capGopher

import (
	"bytes"
	"fmt"
	"net/http"
)

const (
	createTaskURL = "https://api.capmonster.cloud/createTask"
)

type Solver struct {
	apiKey string
}

func (s Solver) RecaptchaV2(siteKey, url string) {
	payload := fmt.Sprintf(`{"clientKey":"%v", {"type": "NoCaptchaProxyless","websiteURL": "%v", "websiteKey": "%v"}}`, s.apiKey, url, siteKey)
	req, err := http.NewRequest(http.MethodPost, createTaskURL, bytes.NewBuffer(bytes(payload)))
}

func main() {
	s := Solver{apiKey: "123"}
	s.RecaptchaV2("6LccSjEUAAAAANCPhaM2c-WiRxCZ5CzsjR_vd8uX", "https://geo.captcha-delivery.com/captcha/?initialCid=AHrlqAAAAAMAkv70eyPjxu0AR0MUlg==&cid=1P-D1yh04r4UAsGESzpI23QxQto8TExFKVZP5MuzDip~.hXrdUFP0PkMXDcBHivAwG9CGg_6MOMub1Sphqq4mVKg2O7jaca5j4xzl_m6NH&referer=http%3A%2F%2Fwww.footlocker.com%2Fapi%2Fusers%2Fcarts%2Fcurrent%2Fentries&hash=A55FBF4311ED6F1BF9911EB71931D5&t=fe&s=17434")
}
