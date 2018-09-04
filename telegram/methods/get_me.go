package methods

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 获取机器人信息
func GetMe(apiaccess, token string) (*BotExt, error) {
	res, err := http.Get(apiaccess + "bot" + token + "/getme")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		reason := fmt.Sprintf("http status code: %d", res.StatusCode)
		return nil, errors.New(reason)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var resonpe GetMeResonpe
	err = json.Unmarshal(body, &resonpe)
	if err != nil {
		return nil, err
	}

	return &BotExt{
		Bot:       *resonpe.Result,
		Token:     token,
		APIAccess: apiaccess,
	}, nil
}
