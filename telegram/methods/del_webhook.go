package methods

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func DelWebhook(apiaccess, token string) error {
	a := [...]string{apiaccess, "bot", token, "/", "deleteWebhook"}
	res, err := http.Get(strings.Join(a[:], ""))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		reason := fmt.Sprintf("http status code: %v", res.StatusCode)
		return errors.New(reason)
	}
	return nil
}
