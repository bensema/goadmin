package captcha

// 验证码字符类型
type StrType int

const (
	NUM   StrType = iota // 数字
	LOWER                // 小写字母
	UPPER                // 大写字母
	ALL                  // 全部
)

// DisturbLevel 干扰级别
type DisturbLevel int

const (
	NORMAL DisturbLevel = 4
	MEDIUM DisturbLevel = 8
	HIGH   DisturbLevel = 16
)

var fontKinds = [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}
var letters = []byte("34578acdefghjkmnpqstwxyABCDEFGHJKMNPQRSVWXY")

type Config struct {
}

type Captcha struct {
	ImgCaptcha *imgCaptcha
	GifCaptcha *GifCaptcha
}

func New(c *Config) (cap *Captcha) {
	cap = &Captcha{
		GifCaptcha: NewGifCaptcha(c),
		ImgCaptcha: NewImgCaptcha(c),
	}

	return
}
