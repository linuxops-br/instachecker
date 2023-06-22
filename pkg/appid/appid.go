package appid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type InstagramApiID interface {
	Get() string
}

type api struct {
	AppID string `json:"appId"`
	url   string
}

func New() InstagramApiID {
	return &api{
		url: "https://www.instagram.com/",
	}
}

func (i *api) Get() string {
	body, _ := requestAppID(i.url)

	i.getValue(body)

	return i.AppID

}

func requestAppID(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", responseBody), nil
}

func (i *api) getValue(body string) {
	reg, _ := regexp.Compile(`"appId":"\d*"`)

	json.Unmarshal(
		[]byte(fmt.Sprintf("{%s}", reg.FindString(body))),
		&i,
	)
}
