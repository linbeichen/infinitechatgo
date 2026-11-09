package service

import (
	"context"
	"crypto/md5"
	"database/sql"
	"fmt"
	"infinitechatgo/common/common"
	utils2 "infinitechatgo/common/utils"
	"infinitechatgo/service/user/api/internal/types/constants"
	"infinitechatgo/service/user/api/internal/types/dto"
	"infinitechatgo/service/user/api/internal/utils"
	"infinitechatgo/service/user/model/user"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type UserService struct {
	userModel user.UserModel
	rds       *redis.Redis
}

func NewUserService(userModel user.UserModel, rds *redis.Redis) *UserService {
	return &UserService{
		userModel: userModel,
		rds:       rds,
	}
}

func (s *UserService) GetUser(account string) (*user.User, error) {
	if constants.PhoneRegex.MatchString(account) {
		user, err := s.userModel.FindOneByPhone(context.Background(), sql.NullString{
			String: account,
			Valid:  true,
		})
		if err != nil {
			return nil, err
		}
		return user, nil
	} else if constants.EmailRegex.MatchString(account) {
		user, err := s.userModel.FindOneByEmail(context.Background(), sql.NullString{
			String: account,
			Valid:  true,
		})
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

func (s *UserService) EncryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(constants.PasswordSalt + password))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (s *UserService) GetTwoToken(userId int64) dto.TwoTokens {
	// 踢掉旧会话（单点登录）
	userIdStr := strconv.FormatInt(userId, 10)

	refreshKey := constants.UserRefreshTokenKeyPrefix + userIdStr
	oldRefreshKey := constants.UserRefreshTokenKeyPrefix + userIdStr
	if oldRefreshKey == "" {
		s.rds.Del(refreshKey)
		common.Infof("用户 %s 被踢下线，旧 Refresh Token: %s\n", userIdStr, oldRefreshKey)
	}

	// 生成新 Token
	accessToken := utils.GenerateAccessToken(userId)
	refreshToken := utils2.NextIdStr()

	s.rds.Setex(constants.AccessTokenKeyPrefix, userIdStr, constants.AccessTokenExpireMinutes*60)
	s.rds.Setex(constants.RefreshTokenKeyPrefix, userIdStr, constants.RefreshTokenExpireMinutes*60)
	s.rds.Setex(constants.UserRefreshTokenKeyPrefix, refreshToken, constants.RefreshTokenExpireMinutes*60)
	// 用户ID 映射 当前 AccessToken
	// s.rds.Setex(constants.UserAccessTokenKeyPrefix + userIdStr, accessToken, constants.AccessTokenExpireMinutes*60)

	common.Infof("用户: %s，accessToken: %s, refreshToken: %s", userIdStr, accessToken, refreshToken)
	return dto.TwoTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
