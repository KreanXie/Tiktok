package redis

import "time"

// Insert given token into blocked_tokens
func BlockToken(token string, expirationSec int) error {
	if _, err := Rdb.Set(Ctx, token, "blocking", 0).Result(); err != nil {
		return err
	}
	if _, err := Rdb.Expire(Ctx, token, time.Duration(expirationSec)*time.Second).Result(); err != nil {
		return err
	}
	return nil
}

// Check if given token exists in blocked_tokens, if so return false
func CheckToken(token string) (bool, error) {
	if status, err := Rdb.Get(Ctx, token).Result(); err == nil && status == "blocking" {
		return false, err
	}
	return true, nil
}
