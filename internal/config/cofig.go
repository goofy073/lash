package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"strconv"
)

type config struct {
	telegramToken          string
	price1                 int
	price3                 int
	price6                 int
	price12                int
	starsPrice1            int
	starsPrice3            int
	starsPrice6            int
	starsPrice12           int
	privateChannelId       string
	privateChannelUsername string
	databaseURL            string
	cryptoPayURL           string
	cryptoPayToken         string
	botURL                 string
	yookasaURL             string
	yookasaShopId          string
	yookasaSecretKey       string
	yookasaEmail           string
	feedbackURL            string
	channelURL             string
	serverStatusURL        string
	supportURL             string
	tosURL                 string
	isYookasaEnabled       bool
	isCryptoEnabled        bool
	isTelegramStarsEnabled bool
	adminTelegramId        int64
	trialDays              int
	referralDays           int
	miniApp                string
}

var conf config

func GetReferralDays() int {
	return conf.referralDays
}

func GetMiniAppURL() string {
	return conf.miniApp
}

func PrivateChannelId() string {
	return conf.privateChannelId
}

func PrivateChannelUsername() string {
	return conf.privateChannelUsername
}

func TrialDays() int {
	return conf.trialDays
}
func FeedbackURL() string {
	return conf.feedbackURL
}

func ChannelURL() string {
	return conf.channelURL
}

func ServerStatusURL() string {
	return conf.serverStatusURL
}

func SupportURL() string {
	return conf.supportURL
}

func TosURL() string {
	return conf.tosURL
}

func YookasaEmail() string {
	return conf.yookasaEmail
}
func Price1() int {
	return conf.price1
}
func Price3() int {
	return conf.price3
}
func Price6() int {
	return conf.price6
}
func Price12() int {
	return conf.price12
}

func Price(month int) int {
	switch month {
	case 1:
		return conf.price1
	case 3:
		return conf.price3
	case 6:
		return conf.price6
	case 12:
		return conf.price12
	default:
		return conf.price1
	}
}

func StarsPrice1() int {
	return conf.starsPrice1
}

func StarsPrice3() int {
	return conf.starsPrice3
}

func StarsPrice6() int {
	return conf.starsPrice6
}

func StarsPrice12() int {
	return conf.starsPrice12
}

func StarsPrice(month int) int {
	switch month {
	case 1:
		return conf.starsPrice1
	case 3:
		return conf.starsPrice3
	case 6:
		return conf.starsPrice6
	case 12:
		return conf.starsPrice12
	default:
		return conf.starsPrice1
	}
}
func TelegramToken() string {
	return conf.telegramToken
}
func DadaBaseUrl() string {
	return conf.databaseURL
}
func CryptoPayUrl() string {
	return conf.cryptoPayURL
}
func CryptoPayToken() string {
	return conf.cryptoPayToken
}
func BotURL() string {
	return conf.botURL
}
func SetBotURL(botURL string) {
	conf.botURL = botURL
}
func YookasaUrl() string {
	return conf.yookasaURL
}
func YookasaShopId() string {
	return conf.yookasaShopId
}
func YookasaSecretKey() string {
	return conf.yookasaSecretKey
}
// Removed TrafficLimit function as it's not needed for channel access

func IsCryptoPayEnabled() bool {
	return conf.isCryptoEnabled
}

func IsYookasaEnabled() bool {
	return conf.isYookasaEnabled
}

func IsTelegramStarsEnabled() bool {
	return conf.isTelegramStarsEnabled
}

func GetAdminTelegramId() int64 {
	return conf.adminTelegramId
}

func InitConfig() {
	err := godotenv.Load(".env")

	conf.adminTelegramId, err = strconv.ParseInt(os.Getenv("ADMIN_TELEGRAM_ID"), 10, 64)
	if err != nil {
		panic("ADMIN_TELEGRAM_ID .env variable not set")
	}

	conf.telegramToken = os.Getenv("TELEGRAM_TOKEN")
	if conf.telegramToken == "" {
		panic("TELEGRAM_TOKEN .env variable not set")
	}

	conf.miniApp = os.Getenv("MINI_APP_URL")
	if conf.miniApp == "" {
		conf.miniApp = ""
	}

	conf.trialDays, err = strconv.Atoi(os.Getenv("TRIAL_DAYS"))
	if err != nil {
		panic("TRIAL_DAYS .env variable not set")
	}

	conf.privateChannelId = os.Getenv("PRIVATE_CHANNEL_ID")
	if conf.privateChannelId == "" {
		panic("PRIVATE_CHANNEL_ID .env variable not set")
	}

	conf.privateChannelUsername = os.Getenv("PRIVATE_CHANNEL_USERNAME")
	if conf.privateChannelUsername == "" {
		panic("PRIVATE_CHANNEL_USERNAME .env variable not set")
	}

	strPrice := os.Getenv("PRICE_1")
	if strPrice == "" {
		panic("PRICE_1 .env variable not set")
	}
	price, err := strconv.Atoi(strPrice)
	if err != nil {
		panic(err)
	}
	conf.price1 = price

	starsPrice1Str := os.Getenv("STARS_PRICE_1")
	if starsPrice1Str != "" {
		priceStar1, err := strconv.Atoi(starsPrice1Str)
		if err != nil {
			panic(err)
		}
		conf.starsPrice1 = priceStar1
	} else {
		conf.starsPrice1 = conf.price1
	}

	strPrice3 := os.Getenv("PRICE_3")
	if strPrice3 == "" {
		panic("PRICE_3 .env variable not set")
	}
	price3, err := strconv.Atoi(strPrice3)
	if err != nil {
		panic(err)
	}
	conf.price3 = price3

	starsPrice3Str := os.Getenv("STARS_PRICE_3")
	if starsPrice3Str != "" {
		priceStar3, err := strconv.Atoi(starsPrice3Str)
		if err != nil {
			panic(err)
		}
		conf.starsPrice3 = priceStar3
	} else {
		conf.starsPrice3 = conf.price3
	}

	strPrice6 := os.Getenv("PRICE_6")
	if strPrice6 == "" {
		panic("PRICE_6 .env variable not set")
	}
	price6, err := strconv.Atoi(strPrice6)
	if err != nil {
		panic(err)
	}
	conf.price6 = price6

	starsPrice6Str := os.Getenv("STARS_PRICE_6")
	if starsPrice6Str != "" {
		priceStar6, err := strconv.Atoi(starsPrice6Str)
		if err != nil {
			panic(err)
		}
		conf.starsPrice6 = priceStar6
	} else {
		conf.starsPrice6 = conf.price6
	}

	strPrice12 := os.Getenv("PRICE_12")
	if strPrice12 == "" {
		panic("PRICE_12 .env variable not set")
	}
	price12, err := strconv.Atoi(strPrice12)
	if err != nil {
		panic(err)
	}
	conf.price12 = price12

	starsPrice12Str := os.Getenv("STARS_PRICE_12")
	if starsPrice12Str != "" {
		priceStar12, err := strconv.Atoi(starsPrice12Str)
		if err != nil {
			panic(err)
		}
		conf.starsPrice12 = priceStar12
	} else {
		conf.starsPrice12 = conf.price12
	}

	conf.databaseURL = os.Getenv("DATABASE_URL")
	if conf.databaseURL == "" {
		panic("DADA_BASE_URL .env variable not set")
	}

	conf.isTelegramStarsEnabled = os.Getenv("TELEGRAM_STARS_ENABLED") == "true"

	conf.isCryptoEnabled = os.Getenv("CRYPTO_PAY_ENABLED") == "true"
	if conf.isCryptoEnabled {
		conf.cryptoPayURL = os.Getenv("CRYPTO_PAY_URL")
		if conf.cryptoPayURL == "" {
			panic("CRYPTO_PAY_URL .env variable not set")
		}
		conf.cryptoPayToken = os.Getenv("CRYPTO_PAY_TOKEN")
		if conf.cryptoPayToken == "" {
			panic("CRYPTO_PAY_TOKEN .env variable not set")
		}
	}

	conf.isYookasaEnabled = os.Getenv("YOOKASA_ENABLED") == "true"
	if conf.isYookasaEnabled {
		conf.yookasaURL = os.Getenv("YOOKASA_URL")
		conf.yookasaShopId = os.Getenv("YOOKASA_SHOP_ID")
		conf.yookasaSecretKey = os.Getenv("YOOKASA_SECRET_KEY")

		if conf.yookasaURL == "" || conf.yookasaShopId == "" || conf.yookasaSecretKey == "" {
			panic("YOOKASA_URL, YOOKASA_SHOP_ID, YOOKASA_SECRET_KEY .env variables not set")
		}

		conf.yookasaEmail = os.Getenv("YOOKASA_EMAIL")
		if conf.yookasaEmail == "" {
			panic("YOOKASA_EMAIL .env variable not set")
		}
	}

	conf.referralDays, err = strconv.Atoi(os.Getenv("REFERRAL_DAYS"))
	if err != nil {
		panic("REFERRAL_DAYS .env variable not set")
	}

	conf.serverStatusURL = os.Getenv("SERVER_STATUS_URL")
	conf.supportURL = os.Getenv("SUPPORT_URL")
	conf.feedbackURL = os.Getenv("FEEDBACK_URL")
	conf.channelURL = os.Getenv("CHANNEL_URL")
	conf.tosURL = os.Getenv("TOS_URL")
}
