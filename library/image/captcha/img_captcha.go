package captcha

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"library/image/captcha/fonts"

	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"math"
	"math/rand"
	"time"
)

type imgCaptcha struct {
	frontColors  []color.Color
	bkgColors    []color.Color
	disturbLevel DisturbLevel
	fonts        []*truetype.Font
	size         image.Point
}

func NewImgCaptcha(c *Config) *imgCaptcha {
	ic := &imgCaptcha{
		disturbLevel: MEDIUM,
		size:         image.Point{X: 128, Y: 48},
	}
	ic.frontColors = []color.Color{color.Black, color.RGBA{R: 255, A: 255}, color.RGBA{B: 255, A: 255}, color.RGBA{G: 153, A: 255}}
	ic.bkgColors = []color.Color{color.White}

	fontData, _ := fonts.Asset("comic.ttf")
	ic.AddFontFromBytes(fontData)

	return ic
}

// AddFont 添加一个字体
func (c *imgCaptcha) AddFont(path string) error {
	fontData, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	font, err := freetype.ParseFont(fontData)
	if err != nil {
		return err
	}
	if c.fonts == nil {
		c.fonts = []*truetype.Font{}
	}
	c.fonts = append(c.fonts, font)
	return nil
}

//AddFontFromBytes allows to load font from slice of bytes, for example, load the font packed by https://github.com/jteeuwen/go-bindata
func (c *imgCaptcha) AddFontFromBytes(contents []byte) error {
	font, err := freetype.ParseFont(contents)
	if err != nil {
		return err
	}
	if c.fonts == nil {
		c.fonts = []*truetype.Font{}
	}
	c.fonts = append(c.fonts, font)
	return nil
}

// SetFont 设置字体 可以设置多个
func (c *imgCaptcha) SetFont(paths ...string) error {
	for _, v := range paths {
		if err := c.AddFont(v); err != nil {
			return err
		}
	}
	return nil
}

func (c *imgCaptcha) SetDisturbance(d DisturbLevel) {
	if d > 0 {
		c.disturbLevel = d
	}
}

func (c *imgCaptcha) SetFrontColor(colors ...color.Color) {
	if len(colors) > 0 {
		c.frontColors = c.frontColors[:0]
		for _, v := range colors {
			c.frontColors = append(c.frontColors, v)
		}
	}
}

func (c *imgCaptcha) SetBkgColor(colors ...color.Color) {
	if len(colors) > 0 {
		c.bkgColors = c.bkgColors[:0]
		for _, v := range colors {
			c.bkgColors = append(c.bkgColors, v)
		}
	}
}

func (c *imgCaptcha) SetSize(w, h int) {
	if w < 48 {
		w = 48
	}
	if h < 20 {
		h = 20
	}
	c.size = image.Point{w, h}
}

func (c *imgCaptcha) randFont() *truetype.Font {
	return c.fonts[rand.Intn(len(c.fonts))]
}

// 绘制背景
func (c *imgCaptcha) drawBkg(img *Image) {
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))
	//填充主背景色
	bgColorIndex := ra.Intn(len(c.bkgColors))
	bkg := image.NewUniform(c.bkgColors[bgColorIndex])
	img.FillBkg(bkg)
}

// 绘制噪点
func (c *imgCaptcha) drawNoises(img *Image) {
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 待绘制图片的尺寸
	size := img.Bounds().Size()
	dlen := int(c.disturbLevel)
	// 绘制干扰斑点
	for i := 0; i < dlen; i++ {
		x := ra.Intn(size.X)
		y := ra.Intn(size.Y)
		r := ra.Intn(size.Y/20) + 1
		colorIndex := ra.Intn(len(c.frontColors))
		img.DrawCircle(x, y, r, i%4 != 0, c.frontColors[colorIndex])
	}

	// 绘制干扰线
	for i := 0; i < dlen; i++ {
		x := ra.Intn(size.X)
		y := ra.Intn(size.Y)
		o := int(math.Pow(-1, float64(i)))
		w := ra.Intn(size.Y) * o
		h := ra.Intn(size.Y/10) * o
		colorIndex := ra.Intn(len(c.frontColors))
		img.DrawLine(x, y, x+w, y+h, c.frontColors[colorIndex])
	}

}

// 绘制文字
func (c *imgCaptcha) drawString(img *Image, str string) {

	if c.fonts == nil {
		panic("没有设置任何字体")
	}
	tmp := NewImage(c.size.X, c.size.Y)

	// 文字大小为图片高度的 0.6
	fsize := int(float64(c.size.Y) * 0.6)
	// 用于生成随机角度
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 文字之间的距离
	// 左右各留文字的1/4大小为内部边距
	padding := fsize / 4
	gap := (c.size.X - padding*2) / (len(str))

	// 逐个绘制文字到图片上
	for i, char := range str {
		// 创建单个文字图片
		// 以文字为尺寸创建正方形的图形
		str := NewImage(fsize, fsize)
		// str.FillBkg(image.NewUniform(color.Black))
		// 随机取一个前景色
		colorIndex := r.Intn(len(c.frontColors))

		//随机取一个字体
		font := c.randFont()
		str.DrawString(font, c.frontColors[colorIndex], string(char), float64(fsize))

		// 转换角度后的文字图形
		rs := str.Rotate(float64(r.Intn(40) - 20))
		// 计算文字位置
		s := rs.Bounds().Size()
		left := i*gap + padding
		top := (c.size.Y - s.Y) / 2
		// 绘制到图片上
		draw.Draw(tmp, image.Rect(left, top, left+s.X, top+s.Y), rs, image.ZP, draw.Over)
	}
	if c.size.Y >= 48 {
		// 高度大于48添加波纹 小于48波纹影响用户识别
		tmp.distortTo(float64(fsize)/10, 200.0)
	}

	draw.Draw(img, tmp.Bounds(), tmp, image.ZP, draw.Over)
}

// Create 生成一个验证码图片
func (c *imgCaptcha) Create(num int, t StrType) (*Image, string) {
	if num <= 0 {
		num = 4
	}
	dst := NewImage(c.size.X, c.size.Y)
	//tmp := NewImage(c.size.X, c.size.Y)
	c.drawBkg(dst)
	c.drawNoises(dst)

	str := string(c.randStr(num, int(t)))
	c.drawString(dst, str)
	//c.drawString(tmp, str)

	return dst, str
}

func (c *imgCaptcha) CreateCustom(str string) *Image {
	if len(str) == 0 {
		str = "unkown"
	}
	dst := NewImage(c.size.X, c.size.Y)
	c.drawBkg(dst)
	c.drawNoises(dst)
	c.drawString(dst, str)
	return dst
}

// 生成随机字符串
// size 个数 kind 模式
func (c *imgCaptcha) randStr(size int, kind int) []byte {
	ikind, result := kind, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll {
			ikind = rand.Intn(3)
		}
		scope, base := fontKinds[ikind][0], fontKinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
		// 不易混淆字符模式：重新生成字符
		if kind == 4 {
			result[i] = letters[rand.Intn(len(letters))]
		}
	}
	return result
}
