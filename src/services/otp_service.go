package services

import (
	"fmt"
	"time"

	"github.com/ardiost/golang-clean-web-api/config"
	"github.com/ardiost/golang-clean-web-api/constants"
	"github.com/ardiost/golang-clean-web-api/data/cache"
	"github.com/go-redis/redis"
)

type OtpService struct {
	cfg         *config.Config
	redisClient *redis.Client
}

type OtpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	redis := cache.GetRedis()
	return &OtpService{cfg: cfg, redisClient: redis}
}

func (s *OtpService) SetOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	val := &OtpDto{
		Value: otp,
		Used:  false,
	}
	res, err := cache.Get[OtpDto](s.redisClient, key)
	if err == nil && !res.Used {
		return err
	} else if err == nil && res.Used {
		return err
		// یه سرویس باید برای ارور ها بنویسم که توی ویدئو ۳ فصل ۲۶ هست
	}

	err = cache.Set(s.redisClient, key, val, s.cfg.Otp.ExpireTime*time.Second)
	if err != nil {
		return err
	}
}

func (s *OtpService) ValidateOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	res, err := cache.Get[OtpDto](s.redisClient, key)
	if err != nil {
		return err
	} else if err == nil && !res.Used {
		return err
	} else if err == nil && res.Used && res.Value != otp {
		return err
	} else if err == nil && res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(s.redisClient, key, res, s.cfg.Otp.ExpireTime*time.Second)
		return nil
	}
	return nil
}
