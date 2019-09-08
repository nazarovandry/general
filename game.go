package main

import (
	"flag"
	"image"
	"image/color"
	_ "image/png"
    "math"
	"time"
    "os"
    "fmt"
    "log"
    "strconv"
    "strings"
    "io/ioutil"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
    appleBonusForm int = iota
    livesBonusForm
    energyBonusForm
)

const (
    energyBlock int = -1
)

const (
    emptyBlock int = iota
    wall0Block
    wall1Block
    wall2Block
    wall3Block
    wall4Block
    wall5Block
    startBlock
    finishBlock
    appleBlock
    livesBlock
    mon0Block
    mon1Block
)

const (
    dDir int = iota
    centerDir
)

const (
    menuCase int = iota
    pauseCase
    gameCase
    screamCase
    editorCase
)

const (
    emptySq int = iota
    wallSq
    iSq
    energySq
    exitSq
    bonusSq
    woodSq
    mon1Sq
    mon2Sq
)

const (
    maxLevSide  = 150
)

var (
	fullscreen      = false
	showMap         = true
	width           = 320
	height          = 200
	scale           = 3.0
    wallsAmount     = 10
    maxLight        = 4.0
    forceLight      = 1.0
    mode            = menuCase
    energy          = 100
    energyUnit      = 0.1
    yesLight        = true
    levNow          = 0
    bonusCount      = 0
    livesStart      = 3
    die             = false
    appleEffect     = 1.0
    livesEffect     = 1.0
    mySpeed         = 3.0
    myWallDist      = 0.4

	as actionSquare

	pos, posStart, dir, plane pixel.Vec
    floor map[int](*image.RGBA)
    ceiling *image.RGBA
    livesPic *image.RGBA
    lives int

    monWhoKilled int

    alpha map[string](*image.RGBA)
    mon = []Monster{}
    bon = []Bonus{}
    bonusInfo = map[int](*BonusForm){}
    retBlock = map[int](*image.RGBA){}
)

func setup() {
	dir = pixel.V(-1.0, 0.0)
	plane = pixel.V(0.0, 0.66)
    lives = livesStart
}

type Level [maxLevSide][maxLevSide]int

var world = [](*Level){}

type BonusForm struct {
    graphs  int
    graph   int
    bonusH  int
    bonusW  int
    pic     [](*image.RGBA)
    show    bool
}

type Monster struct {
    X       float64
    Y       float64
    XStart  float64
    YStart  float64
    side    float64
    height  float64
    form    int
    pic     *image.RGBA
    live    float64
    alive   bool
}

type Bonus struct {
    X       float64
    Y       float64
    taken   bool
    form    int
}

type Form struct {
    speed       float64
    rand        bool
    through     bool
    liveMax     float64
    screamer    *image.RGBA
}

var forms = []Form{
    Form{speed: 0.03, rand: false, liveMax: 5.0},
    Form{speed: 0.02, rand: true, through: false, liveMax: 10.0},
}

var wallTexture = make([](*image.RGBA), wallsAmount)

func getLevel(path string) (*Level) {
    levPic := getPic(path)
    bounds := levPic.Bounds()
    lev := Level{}
    for i := 0; i < bounds.Max.X; i++ {
        for j := 2; j < bounds.Max.Y; j++ {
            colorTo := levPic.At(i, j)
            lev[i][j-2] = 0
            switch colorTo {
                case levPic.At(1, 0):
                    lev[i][j-2] = wall0Block
                case levPic.At(2, 0):
	                pos = pixel.V(float64(i)+0.5, float64(j)+0.5-2)
	                posStart = pixel.V(float64(i)+0.5, float64(j)+0.5-2)
                case levPic.At(3, 0):
                    bon = append(bon, Bonus{X: float64(i)+0.5, Y: float64(j)+0.5-2,
                        taken: false, form: energyBonusForm})
                    lev[i][j-2] = energyBlock
                case levPic.At(4, 0):
                    bon = append(bon, Bonus{X: float64(i)+0.5, Y: float64(j)+0.5-2,
                        taken: false, form: livesBonusForm})
                    lev[i][j-2] = 0
                case levPic.At(5, 0):
                    bon = append(bon, Bonus{X: float64(i)+0.5, Y: float64(j)+0.5-2,
                        taken: false, form: appleBonusForm})
                    lev[i][j-2] = 0
                case levPic.At(6, 0):
                    lev[i][j-2] = wall2Block
                case levPic.At(7, 0):
                    lev[i][j-2] = wall3Block
                case levPic.At(8, 0):
                    lev[i][j-2] = wall4Block
                case levPic.At(9, 0):
                    lev[i][j-2] = wall5Block
                case levPic.At(0, 1):
                    mon = append(mon, Monster{XStart: float64(i)+0.5, YStart: float64(j)+0.5-2,
                        side: 0.4, height: 0.8, form: 0})
                case levPic.At(1, 1):
                    mon = append(mon, Monster{XStart: float64(i)+0.5, YStart: float64(j)+0.5-2,
                        side: 0.4, height: 0.8, form: 1})
            }
        }
    }
    return &lev
}

func getColor(x, y int) color.RGBA {
	switch world[levNow][x][y] {
	case 0:
		return color.RGBA{43, 30, 24, 255}
	case 1:
		return color.RGBA{100, 89, 73, 255}
	case 2:
		return color.RGBA{110, 23, 0, 255}
	case 3:
		return color.RGBA{45, 103, 171, 255}
	case 4:
		return color.RGBA{123, 84, 33, 255}
	case 5:
		return color.RGBA{158, 148, 130, 255}
	case 6:
		return color.RGBA{203, 161, 47, 255}
	case 7:
		return color.RGBA{255, 107, 0, 255}
	case 9:
		return color.RGBA{0, 0, 0, 0}
	default:
		return color.RGBA{255, 194, 32, 255}
	}
}

func setLight(lightDist float64, col *color.Color) {
    d := lightDist / maxLight
    k := 1 / forceLight
    if d >= 1 {
        *col = color.RGBA{0, 0, 0, 0}
    } else {
        d = 1 - d
        root := math.Pow(d, k)
        r, g, b, a := (*col).RGBA()
        r = uint32(float64(r) / 255 * root)
        g = uint32(float64(g) / 255 * root)
        b = uint32(float64(b) / 255 * root)
        a = uint32(float64(a) / 255 * root)
        g = uint32(float64(g) / livesEffect)
        b = uint32(float64(b) / livesEffect)
        r = uint32(float64(r) / appleEffect)
        b = uint32(float64(b) / appleEffect)
        *col = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
    }
}

func frameGame() *image.RGBA {
	im := image.NewRGBA(image.Rect(0, 0, width, height))
    space := float64(width) / 2

    for x := 0; x < width; x++ {
		cameraX := 2 * float64(x) / float64(width) - 1
        look := (pixel.V(0, 1)).Scaled(space/(float64(width)/2)).Add(pixel.V(cameraX, 0)).Rotated(plane.Angle())
        iterX := 0
        iterY := 0
        signX := -1
        signY := -1
        dX := pos.X - float64(int(pos.X))
        if look.X > 0 {
            dX = 1.0 - dX
            signX = 1
        }
        dY := pos.Y - float64(int(pos.Y))
        if look.Y > 0 {
            dY = 1.0 - dY
            signY = 1
        }
        empty := true
        for empty {
            distX := dX + float64(iterX)
            distY := dY + float64(iterY)
            parX := math.Abs(look.X) * distY
            parY := math.Abs(look.Y) * distX
            inVer := false
            fromLeftSide := 0.0
            if parX > parY {
                iterX += 1
                inVer = true
                if math.Abs(look.X) > 0 {
                    fromLeftSide = distY - math.Abs(look.Scaled(distX/math.Abs(look.X)).Y)
                }
                if look.X > 0 && look.Y < 0 || look.X < 0 && look.Y > 0 {
                    fromLeftSide = 1.0 - fromLeftSide
                }
            } else {
                iterY += 1
                if math.Abs(look.Y) > 0 {
                    fromLeftSide = distX - math.Abs(look.Scaled(distY/math.Abs(look.Y)).X)
                }
                if look.X > 0 && look.Y > 0 || look.X < 0 && look.Y < 0 {
                    fromLeftSide = 1.0 - fromLeftSide
                }
            }
            whereX := int(pos.X) + iterX*signX
            whereY := int(pos.Y) + iterY*signY
            ind := world[levNow][whereX][whereY]
            if ind > 0 {
                empty = false
                finalDist := 0.0
                if inVer && math.Abs(look.X) > 0 {
                    finalDist = look.Len() * distX / math.Abs(look.X)
                } else {
                    finalDist = look.Len() * distY / math.Abs(look.Y)
                }
                for y := 0; y < height; y++ {
		            cameraY := 2 * float64(y) / float64(height) - 1
                    up := (finalDist*(float64(height)/2)/space) * cameraY
                    var colorTo color.Color
                    lightDist := math.Sqrt(finalDist * finalDist + up * up)
                    if up > 0.5 {
                        floorVec := pos.Add(pixel.V(finalDist * 0.5 / up, 0).Rotated(look.Angle()))
                        floorType := world[levNow][int(floorVec.X)][int(floorVec.Y)]
                        bounds := floor[floorType].Bounds()
                        floorVec.X = floorVec.X - float64(int(floorVec.X))
                        floorVec.Y = floorVec.Y - float64(int(floorVec.Y))
                        colorTo = floor[floorType].At(int(float64(bounds.Max.X)*floorVec.X),
                            int(float64(bounds.Max.Y)*floorVec.Y))
                        lightDist *= (0.5 / up)
                    } else if up < -0.5 {
                        bounds := ceiling.Bounds()
                        floorVec := pos.Add(pixel.V(finalDist * 0.5 / (-up), 0).Rotated(look.Angle()))
                        floorVec.X = floorVec.X - float64(int(floorVec.X))
                        floorVec.Y = floorVec.Y - float64(int(floorVec.Y))
                        colorTo = ceiling.At(int(float64(bounds.Max.X)*floorVec.X),
                            int(float64(bounds.Max.Y)*floorVec.Y))
                        lightDist *= (0.5 / (-up))
                    } else {
                        bounds := wallTexture[ind].Bounds()
                        colorTo = wallTexture[ind].At(int(float64(bounds.Max.X)*fromLeftSide),
                            int(float64(bounds.Max.Y)*(up+0.5)))
                    }
                    setLight(lightDist, &colorTo)
                    im.Set(x, y, colorTo)
                }
                for y := 0; y < height; y++ {
                    minDistMon := maxLight + 1.0
                    drown := false
                    for _, m := range mon {
                        draw := false
                        monDist := pixel.V(m.X, m.Y).Sub(pos)
                        if !m.alive || monDist.Len() > maxLight {
                            continue
                        }
                        if monDist.Len() < minDistMon || drown == false {
                            minDistMon = monDist.Len()
                            draw = true
                        }
                        try := look.Rotated(-monDist.Angle())
                        if try.X > 0 {
                            try = try.Scaled(monDist.Len() / try.X)
                        } else {
                            continue
                        }
                        if math.Abs(try.Y) > m.side || try.Len() >= finalDist {
                            continue
                        }
                        bounds := m.pic.Bounds()
                        cameraY := 2 * float64(y) / float64(height) - 1
                        up := (try.X*(float64(height)/2)/space) * cameraY
                        lightDist := math.Sqrt(monDist.Len() * monDist.Len() + up * up)
                        if up > m.height - 0.5 || up < -0.5 {
                            continue
                        }
                        colorTo := m.pic.At(int(float64(bounds.Max.X)*(1-(try.Y/m.side/2+0.5))),
                            int(float64(bounds.Max.Y)*(1-(up+0.5)/m.height)))
                        if draw && colorTo != m.pic.At(0, 0) {
                            setLight(lightDist, &colorTo)
                            im.Set(x, height - y, colorTo)
                            drown = true
                        }
                    }
                    for _, m := range bon {
                        draw := false
                        monDist := pixel.V(m.X, m.Y).Sub(pos)
                        if monDist.Len() > maxLight || m.taken || !bonusInfo[m.form].show {
                            continue
                        }
                        if monDist.Len() < minDistMon || drown == false {
                            minDistMon = monDist.Len()
                            draw = true
                        }
                        try := look.Rotated(-monDist.Angle())
                        if try.X > 0 {
                            try = try.Scaled(monDist.Len() / try.X)
                        } else {
                            continue
                        }
                        if math.Abs(try.Y) > 0.2 || try.Len() >= finalDist {
                            continue
                        }
                        f := m.form
                        g := bonusInfo[f].graph
                        bounds := bonusInfo[f].pic[g].Bounds()
                        cameraY := 2 * float64(y) / float64(height) - 1
                        up := (try.X*(float64(height)/2)/space) * cameraY
                        lightDist := math.Sqrt(monDist.Len() * monDist.Len() + up * up)
                        if up > 0 || up < -0.4 {
                            continue
                        }
                        colorTo := bonusInfo[f].pic[g].At(int(float64(bounds.Max.X)*(1-(try.Y/0.2/2+0.5))),
                            int(float64(bounds.Max.Y)*(1-(up+0.4)/0.4)))
                        if draw && colorTo != bonusInfo[f].pic[g].At(0, 0) {
                            setLight(lightDist, &colorTo)
                            im.Set(x, height - y, colorTo)
                            drown = true
                        }
                    }
                }
            }
            if float64(iterX) > 100 || float64(iterY) > 100 {
                empty = false
            }
        }
    }
    drawPic(im, 2, 2, bonusInfo[appleBonusForm].pic[0], 0.5)
    writeText(im, strconv.Itoa(bonusCount), 30, 2, dDir, color.RGBA{255, 0, 0, 255}, 3)
    drawPic(im, 70, 2, bonusInfo[livesBonusForm].pic[0], 0.3)
    writeText(im, strconv.Itoa(lives), 98, 2, dDir, color.RGBA{255, 0, 0, 255}, 3)
    drawPic(im, 138, 2, bonusInfo[energyBonusForm].pic[0], 0.5)
    writeText(im, strconv.Itoa(energy)+"%", 166, 2, dDir, color.RGBA{255, 0, 0, 255}, 3)
    return im
}

func minimap() *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, 24, 26))

	//center := world[levNow][int(pos.X)][int(pos.X)]

    for x, row := range world[levNow] {
		for y, _ := range row {
			c := getColor(x, y)
			if c.A == 255 {
				c.A = 96
			}
			m.Set(x, y, c)
		}
	}

	m.Set(int(pos.X), int(pos.Y), color.RGBA{255, 0, 0, 255})
	for _, i := range bon {
        if !i.taken {
            switch i.form {
                case appleBonusForm:
                    m.Set(int(i.X), int(i.Y), color.RGBA{0, 255, 0, 255})
                case livesBonusForm:
                    m.Set(int(i.X), int(i.Y), color.RGBA{0, 255, 255, 255})
                case energyBonusForm:
                    m.Set(int(i.X), int(i.Y), color.RGBA{255, 255, 0, 255})
            }
        }
    }
    for _, i := range mon {
        if i.alive {
            m.Set(int(i.X), int(i.Y), color.RGBA{0, 0, 255, 255})
        }
    }


	if as.active {
		m.Set(as.X, as.Y, color.RGBA{255, 255, 255, 255})
	} else {
		m.Set(as.X, as.Y, color.RGBA{64, 64, 64, 255})
	}

	return m
}

func getActionSquare() actionSquare {
	pt := image.Pt(int(pos.X)+1, int(pos.Y))

	a := dir.Angle()

	switch {
	case a > 2.8 || a < -2.8:
		pt = image.Pt(int(pos.X)-1, int(pos.Y))
	case a > -2.8 && a < -2.2:
		pt = image.Pt(int(pos.X)-1, int(pos.Y)-1)
	case a > -2.2 && a < -1.4:
		pt = image.Pt(int(pos.X), int(pos.Y)-1)
	case a > -1.4 && a < -0.7:
		pt = image.Pt(int(pos.X)+1, int(pos.Y)-1)
	case a > 0.4 && a < 1.0:
		pt = image.Pt(int(pos.X)+1, int(pos.Y)+1)
	case a > 1.0 && a < 1.7:
		pt = image.Pt(int(pos.X), int(pos.Y)+1)
	case a > 1.7:
		pt = image.Pt(int(pos.X)-1, int(pos.Y)+1)
	}

	block := -1
	active := pt.X > 0 && pt.X < 23 && pt.Y > 0 && pt.Y < 23

	if active {
		block = world[levNow][pt.X][pt.Y]
	}

	return actionSquare{
		X:      pt.X,
		Y:      pt.Y,
		active: active,
		block:  block,
	}
}

type actionSquare struct {
	X      int
	Y      int
	block  int
	active bool
}

func (as actionSquare) toggle(n int) {
	if as.active {
		if world[levNow][as.X][as.Y] == 0 {
			world[levNow][as.X][as.Y] = n
		} else {
			world[levNow][as.X][as.Y] = 0
		}
	}
}

func (as actionSquare) set(n int) {
	if as.active {
		world[levNow][as.X][as.Y] = n
	}
}

func bonusMove() {
    for i, b := range bon {
        monDir := pos.Sub(pixel.V(b.X, b.Y))
        if monDir.Len() < 0.5 && !b.taken {
            switch b.form {
                case appleBonusForm:
                    bonusCount += 1
                    bon[i].taken = true
                    appleEffect += 1.0
                case livesBonusForm:
                    lives += 1
                    bon[i].taken = true
                    livesEffect += 1.0
                case energyBonusForm:
                    if !yesLight {
                        yesLight = true
                        maxLight = 4.0
                    }
                    if energy < 100 {
                        energy += 1
                    }
            }
        }
    }
}

func monsterMove() {
    for i, m := range mon {
        monDir := pos.Sub(pixel.V(m.X, m.Y))
        if !m.alive {
            continue
        }
        if monDir.Len() < 0.5 {
            monWhoKilled = m.form
            die = true
            lives -= 1
            break
        }
        if forceLight > 1.1 && monDir.Len() <= maxLight && m.live > 0 {
            mon[i].live -= 0.01 * monDir.Len() * forceLight
        }
        if m.live <= 0 {
            mon[i].alive = false
        }
        speed := forms[m.form].speed
        monDir = pixel.V(speed, 0).Rotated(monDir.Angle())
        newX := m.X + monDir.X
        newY := m.Y + monDir.Y
        signX := -1.0
        signY := -1.0
        if monDir.X > 0 {
            signX = 1.0
        }
        if monDir.Y > 0 {
            signY = 1.0
        }
        if forms[m.form].through {
            mon[i].X = newX
            mon[i].Y = newY
            continue
        }
        stopX := false
        seeCorner := true
        if world[levNow][int(newX + m.side*signX)][int(newY)] != 0 {
            newX = m.X
            newY = m.Y + speed*signY
            stopX = true
            seeCorner = false
        }
        if world[levNow][int(newX)][int(newY + m.side*signY)] != 0 {
            newY = m.Y
            newX = m.X + speed*signX
            if stopX {
                newX = m.X
            }
            seeCorner = false
        }
        if !seeCorner {
            mon[i].X = newX
            mon[i].Y = newY
            continue
        }
        diff := pixel.V((newX - float64(int(newX))) * (-1),
            (newY - float64(int(newY))) * (-1))
        if diff.X > 0.5 {
            diff.X -= 1.0
        }
        if diff.X < -0.5 {
            diff.X += 1.0
        }
        if diff.Y > 0.5 {
            diff.Y -= 1.0
        }
        if diff.Y < -0.5 {
            diff.Y += 1.0
        }
        if world[levNow][int(newX + diff.X*1.01)][int(newY + diff.Y*1.01)] != 0 {
            monDir = monDir.Scaled(m.side/monDir.Len())
            newdiff := diff.Scaled(speed/diff.Len())
            if pixel.V(1, 0).Rotated(monDir.Angle()).Rotated(-diff.Angle()).Y > 0 {
                newdiff = newdiff.Rotated(math.Pi/2)
            } else {
                newdiff = newdiff.Rotated(-math.Pi/2)
            }
            newX = m.X + newdiff.X
            newY = m.Y + newdiff.Y
        }
        mon[i].X = newX
        mon[i].Y = newY
    }
}

func run() {
	cfg := pixelgl.WindowConfig{
		Bounds:      pixel.R(0, 0, float64(width)*scale, float64(height)*scale),
		VSync:       true,
		Undecorated: false,
	}

	if fullscreen {
		cfg.Monitor = pixelgl.PrimaryMonitor()
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

    win.SetCursorVisible(false)

    for !win.Closed() {
        switch mode {
            case gameCase:
                if gamePress(win) {
                    return
                }
            case menuCase, pauseCase:
                if menuPress(win) {
                    return
                }
            case screamCase:
                if screamPress(win) {
                    return
                }
            case editorCase:
                if editorPress(win) {
                    return
                }
        }
    }
}

func lenSpec(text string) (int) {
    ret := 0
    for _ = range text {
        ret += 1
    }
    return ret
}

func drawPic(im *image.RGBA, x int, y int, imFrom *image.RGBA, scale float64) {
    bounds := imFrom.Bounds()
    emptyCol := imFrom.At(0, 0)
    for i := 0; i < bounds.Max.X; i++ {
        for j := 0; j < bounds.Max.Y; j++ {
            colorTo := imFrom.At(i, j)
            if colorTo != emptyCol {
                im.Set(x + int(float64(i)*scale), y + int(float64(j)*scale), colorTo)
            }
        }
    }
}

func drawPicIn(im *image.RGBA, x int, y int, x2 int, y2 int, imFrom *image.RGBA) {
    bounds := imFrom.Bounds()
    emptyCol := imFrom.At(0, 0)
    w := x2 - x
    h := y2 - y
    bw := bounds.Max.X
    bh := bounds.Max.Y
    for i := 0; i < w; i++ {
        for j := 0; j < h; j++ {
            colorTo := imFrom.At(i*bw/w, j*bh/h)
            if colorTo != emptyCol {
                im.Set(x + i, y + j, colorTo)
            }
        }
    }
}

func writeText(im *image.RGBA, text string, x int, y int, dir int, col color.Color, scale int) {
	xStart := x
	yStart := y
    if dir == centerDir {
        xStart = x - lenSpec(text)*(6*scale)/2
        yStart = y - (4*scale)
    }
    ind := 0
    for _, s := range text {
        emptyCol := alpha[string(s)].At(4, 7)
        for i := 0; i < (5*scale); i++ {
            for j := 0; j < (8*scale); j++ {
                colorTo := alpha[string(s)].At(i/scale, j/scale)
                if colorTo != emptyCol {
                    im.Set(xStart + ind * ((5 + 1)*scale) + i, yStart + j, col)
                }
            }
        }
        ind += 1
    }
}

func is(a int, b int) (int) {
    if a == b {
        return 1
    }
    return 0
}

func isnt(a int, b int) (int) {
    if a != b {
        return 1
    }
    return 0
}

func editorPress(win *pixelgl.Window) (bool) {
    startMouse := pixel.V(float64(width)/2, float64(height)/2)
    win.SetMousePosition(startMouse)

    getL, err := ioutil.ReadFile("lev/level01.txt")
    if err != nil {
        fmt.Println(err.Error())
        mode = menuCase
        return false
    }
    l := strings.Replace(strings.Replace(string(getL), "\n", "", -1), "\t", "", -1)

    wh := strings.Split(l, "?")
    widthPlan, _ := strconv.Atoi(wh[0])
    heightPlan, _ := strconv.Atoi(wh[1])
    pars := strings.Split(wh[2], "|")
    levelPlan := [][]int{}
    for i := range pars {
        elems := strings.Split(pars[i], " ")
        levelPlan = append(levelPlan, []int{})
        for j := 0; j < heightPlan; j++ {
            elem, _ := strconv.Atoi(elems[j])
            levelPlan[i] = append(levelPlan[i], elem)
        }
    }

    for !win.Closed() && mode == editorCase {
        if win.JustPressed(pixelgl.KeyEscape) {
			mode = menuCase
		}

		win.Clear(color.Black)

	    im := image.NewRGBA(image.Rect(0, 0, width, height))
        //bx := imFrom.Bounds().Max.X
        //by := imFrom.Bounds().Max.Y

        drawPicIn(im, 20, 2, 38, 20, floor[0])
        drawPicIn(im, 40, 2, 58, 20, floor[-1])
        drawPicIn(im, 60, 2, 78, 20, bonusInfo[0].pic[0])
        drawPicIn(im, 80, 2, 98, 20, bonusInfo[1].pic[0])
        drawPicIn(im, 100, 2, 118, 20, mon[0].pic)
        drawPicIn(im, 120, 2, 138, 20, mon[1].pic)
        drawPicIn(im, 140, 2, 158, 20, wallTexture[1])
        drawPicIn(im, 160, 2, 178, 20, wallTexture[2])
        drawPicIn(im, 180, 2, 198, 20, wallTexture[3])

        starti, startj := 0, 25
        sc := 20
        for i := 0; i < widthPlan; i++ {
            for j := 0; j < heightPlan; j++ {
                drawPicIn(im, starti + i*sc, startj + j*sc,
                    starti + (i+1)*sc, startj + (j+1)*sc, retBlock[levelPlan[i][j]])
            }
        }

		p := pixel.PictureDataFromImage(im)

	    c := win.Bounds().Center()
		pixel.NewSprite(p, p.Bounds()).Draw(win, pixel.IM.Moved(c).Scaled(c, scale))

        win.Update()
    }

    return false
}

func screamPress(win *pixelgl.Window) (bool) {
    startMouse := pixel.V(float64(width)/2, float64(height)/2)
    win.SetMousePosition(startMouse)
    timer := 0

    for !win.Closed() && mode == screamCase {
		win.Clear(color.Black)

        timer += 1
        if timer > 50 {
            mode = gameCase
        }

	    im := image.NewRGBA(image.Rect(0, 0, width, height))
        imFrom := forms[monWhoKilled].screamer
        bx := imFrom.Bounds().Max.X
        by := imFrom.Bounds().Max.Y

        for i := 0; i < width; i++ {
            for j := 0; j < height; j++ {
                im.Set(i, j, imFrom.At(i * bx / width, j * by / height))
            }
        }

		p := pixel.PictureDataFromImage(im)

	    c := win.Bounds().Center()
		pixel.NewSprite(p, p.Bounds()).Draw(win, pixel.IM.Moved(c).Scaled(c, scale))

        win.Update()
    }

    return false
}

func frameMenu(pointer int, graph int) *image.RGBA {
	im := image.NewRGBA(image.Rect(0, 0, width, height))
    switch mode {
        case menuCase:
            writeText(im, "ИГРАТЬ", width/2, height/3, centerDir,
                color.RGBA{50*uint8(isnt(pointer, 0)), uint8((150 + 3*graph)*is(pointer, 0)), 0, 255}, 2)
        case pauseCase:
            writeText(im, "ПРОДОЛЖИТЬ", width/2, height/3, centerDir,
                color.RGBA{50*uint8(isnt(pointer, 0)), uint8((150 + 3*graph)*is(pointer, 0)), 0, 255}, 2)
    }
    writeText(im, "ПОМЕНЯТЬ ЧТО-ТО", width/2, height/2, centerDir,
        color.RGBA{50*uint8(isnt(pointer, 1)), uint8((150 + 3*graph)*is(pointer, 1)), 0, 255}, 2)
    writeText(im, "ВЫЙТИ ИЗ ИГРЫ", width/2, height*2/3, centerDir,
        color.RGBA{50*uint8(isnt(pointer, 2)), uint8((150 + 3*graph)*is(pointer, 2)), 0, 255}, 2)
    return im
}

func menuPress(win *pixelgl.Window) (bool) {
    startMouse := pixel.V(float64(width)/2, float64(height)/2)
    win.SetMousePosition(startMouse)

    pointer := 0
    graph := 0
    pressUpTime := 0
    pressDownTime := 0

    for !win.Closed() && (mode == menuCase || mode == pauseCase) {
        if win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ) {
			return true
		}

		win.Clear(color.Black)
        graph = (graph + 1) % 30

        if win.Pressed(pixelgl.KeyUp) || win.Pressed(pixelgl.KeyW) {
			if pressUpTime == 0 {
                pointer = (pointer - 1 + 2) % 3
            }
            pressUpTime = (pressUpTime + 1) % 20
		} else {
            pressUpTime = 0
        }

		if win.Pressed(pixelgl.KeyDown) || win.Pressed(pixelgl.KeyS) {
			if pressDownTime == 0 {
			    pointer = (pointer + 1) % 3
            }
            pressDownTime = (pressDownTime + 1) % 20
		} else {
            pressDownTime = 0
        }

        if win.Pressed(pixelgl.KeyEnter) || win.Pressed(pixelgl.KeySpace) {
			switch pointer {
                case 0:
                    mode = gameCase
                case 1:
                    mode = editorCase
                case 2:
                    return true
            }
		}

		p := pixel.PictureDataFromImage(frameMenu(pointer, graph))

	    c := win.Bounds().Center()
		pixel.NewSprite(p, p.Bounds()).Draw(win, pixel.IM.Moved(c).Scaled(c, scale))

        win.Update()
    }

    return false
}

func gamePress(win *pixelgl.Window) (bool) {
	c := win.Bounds().Center()
	last := time.Now()
	mapRot := -1.6683362599999894

    startMouse := pixel.V(float64(width)/2, float64(height)/2)
    win.SetMousePosition(startMouse)

	for !win.Closed() && mode == gameCase {
        if die && lives > 0 {
            mode = screamCase
            die = false
            pos = posStart
            yesLight = true
            energy = 100
            energyUnit = 0.1
            maxLight = 4.0
            forceLight = 1.0
        }

        if win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyP) {
			mode = pauseCase
		}

        if appleEffect > 1.01 {
            appleEffect -= 0.2
        }
        if livesEffect > 1.01 {
            livesEffect -= 0.2
        }

        if die && lives == 0 {
            mode = screamCase
            die = false
            pos = posStart
            yesLight = true
            energy = 100
            energyUnit = 0.1
            maxLight = 4.0
            forceLight = 1.0
            for i, m := range mon {
                mon[i].X = m.XStart
                mon[i].Y = m.YStart
                mon[i].live = forms[mon[i].form].liveMax
                mon[i].alive = true
            }
            bonusCount = 0
            for i := range bon {
                bon[i].taken = false
            }
            lives = livesStart
			mode = menuCase
        }

        if yesLight {
            energyUnit -= maxLight * forceLight * 0.01
            if energyUnit < 0 {
                energyUnit = 1
                if energy > 0 {
                    energy -= 1
                } else {
                    yesLight = false
                }
            }
        } else {
            maxLight = 0
        }

        bonusInfo[appleBonusForm].graph = (bonusInfo[appleBonusForm].graph + 1) %
            bonusInfo[appleBonusForm].graphs
        bonusInfo[livesBonusForm].graph = (bonusInfo[livesBonusForm].graph + 1) %
            bonusInfo[livesBonusForm].graphs
        bonusInfo[energyBonusForm].graph = (bonusInfo[energyBonusForm].graph + 1) %
            bonusInfo[energyBonusForm].graphs

		win.Clear(color.Black)

		dt := time.Since(last).Seconds()
		last = time.Now()

		as = getActionSquare()

        monsterMove()
        bonusMove()

        newMousePosition := win.MousePosition()
        diffMouse := newMousePosition.Sub(startMouse).X
        win.SetMousePosition(startMouse)

		if win.Pressed(pixelgl.KeyUp) || win.Pressed(pixelgl.KeyW) {
			moveForward(mySpeed * dt)
		}

		if win.Pressed(pixelgl.KeyLeft) || win.Pressed(pixelgl.KeyA) {
			moveLeft(mySpeed * dt)
		}

		if win.Pressed(pixelgl.KeyDown) || win.Pressed(pixelgl.KeyS) {
			moveBackwards(mySpeed * dt)
		}

		if win.Pressed(pixelgl.KeyRight) ||  win.Pressed(pixelgl.KeyD) {
			moveRight(mySpeed * dt)
		}

		if win.Pressed(pixelgl.KeyU) || diffMouse > 0.01 {
			turnRight(0.2 * dt * diffMouse)
		}

		if win.Pressed(pixelgl.KeyI) || diffMouse < -0.01 {
			turnLeft(-0.2 * dt * diffMouse)
		}

        if yesLight {
            if win.Pressed(pixelgl.KeyKPSubtract) {
			    if maxLight > 4.01 {
			        maxLight -= 0.1
                }
		    }
            if win.Pressed(pixelgl.KeyKPAdd) {
			    if maxLight < 9.99 {
                    maxLight += 0.1
                }
		    }
            if forceLight > 1.01 {
                forceLight -= 0.1
            }
            if win.Pressed(pixelgl.MouseButtonLeft) {
			    if forceLight < 4.99 {
                    forceLight += 0.3
                }
		    }
        }

		if win.JustPressed(pixelgl.KeyM) {
			showMap = !showMap
		}

		if win.JustPressed(pixelgl.KeySpace) {
			as.toggle(3)
		}

		p := pixel.PictureDataFromImage(frameGame())

		pixel.NewSprite(p, p.Bounds()).Draw(win, pixel.IM.Moved(c).Scaled(c, scale))

		if showMap {
			m := pixel.PictureDataFromImage(minimap())

			mc := m.Bounds().Min.Add(pixel.V(-m.Rect.W(), m.Rect.H()))

			pixel.NewSprite(m, m.Bounds()).
				Draw(win, pixel.IM.
					Moved(mc).
					Rotated(mc, mapRot).
					ScaledXY(pixel.ZV, pixel.V(-scale*2, scale*2)))
		}

		win.Update()
    }
    return false
}

func moveForward(s float64) {
	if world[levNow][int(pos.X+dir.X*(s+myWallDist))][int(pos.Y)] <= 0 {
		pos.X += dir.X * s
	}
	if world[levNow][int(pos.X)][int(pos.Y+dir.Y*(s+myWallDist))] <= 0 {
		pos.Y += dir.Y * s
	}
}

func moveLeft(s float64) {
	if world[levNow][int(pos.X-plane.X*(s+myWallDist))][int(pos.Y)] <= 0 {
		pos.X -= plane.X * s
	}
	if world[levNow][int(pos.X)][int(pos.Y-plane.Y*(s+myWallDist))] <= 0 {
		pos.Y -= plane.Y * s
	}
}

func moveBackwards(s float64) {
	if world[levNow][int(pos.X-dir.X*(s+myWallDist))][int(pos.Y)] <= 0 {
		pos.X -= dir.X * s
	}
	if world[levNow][int(pos.X)][int(pos.Y-dir.Y*(s+myWallDist))] <= 0 {
		pos.Y -= dir.Y * s
	}
}

func moveRight(s float64) {
	if world[levNow][int(pos.X+plane.X*(s+myWallDist))][int(pos.Y)] <= 0 {
		pos.X += plane.X * s
	}
	if world[levNow][int(pos.X)][int(pos.Y+plane.Y*(s+myWallDist))] <= 0 {
		pos.Y += plane.Y * s
	}
}

func turnRight(s float64) {
	oldDirX := dir.X
	dir.X = dir.X*math.Cos(-s) - dir.Y*math.Sin(-s)
	dir.Y = oldDirX*math.Sin(-s) + dir.Y*math.Cos(-s)

	oldPlaneX := plane.X
	plane.X = plane.X*math.Cos(-s) - plane.Y*math.Sin(-s)
	plane.Y = oldPlaneX*math.Sin(-s) + plane.Y*math.Cos(-s)
}

func turnLeft(s float64) {
	oldDirX := dir.X
	dir.X = dir.X*math.Cos(s) - dir.Y*math.Sin(s)
	dir.Y = oldDirX*math.Sin(s) + dir.Y*math.Cos(s)

	oldPlaneX := plane.X
	plane.X = plane.X*math.Cos(s) - plane.Y*math.Sin(s)
	plane.Y = oldPlaneX*math.Sin(s) + plane.Y*math.Cos(s)
}

func getPic(path string) (*image.RGBA) {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println(err.Error())
        return image.NewRGBA(image.Rect(0, 0, 1, 1))
    }
    defer file.Close()
    pic, _, err := image.Decode(file)
    if err != nil {
        fmt.Println(err.Error())
        return image.NewRGBA(image.Rect(0, 0, 1, 1))
    }
    bounds := pic.Bounds()
	to := image.NewRGBA(image.Rect(0, 0, bounds.Max.X, bounds.Max.Y))
    for i := 0; i < bounds.Max.X; i++ {
        for j := 0; j < bounds.Max.Y; j++ {
            to.Set(i, j, pic.At(i, j))
        }
    }
    return to
}

func getBonus(form int, spritesAmount int, path string, show bool) {
    bonusInfo[form] = &BonusForm{}
    a := getPic(path)
    bounds := a.Bounds()
    bh := bounds.Max.Y
    bw := bounds.Max.X / spritesAmount
    bonusInfo[form].bonusH = bh
    bonusInfo[form].bonusW = bw
    bonusInfo[form].graphs = spritesAmount
    bonusInfo[form].graph = 0
    bonusInfo[form].show = show
    bonusInfo[form].pic = [](*image.RGBA){}
    for k := 0; k < spritesAmount; k++ {
        bonusInfo[form].pic = append(bonusInfo[form].pic,
            image.NewRGBA(image.Rect(0, 0, bw, bh)))
        for i := 0; i < bw; i++ {
            for j := 0; j < bh; j++ {
                bonusInfo[form].pic[k].Set(i, j, a.At(k*bw + i, j))
            }
        }
    }
}

func getAlpha(w int, h int) (map[string](*image.RGBA)) {
    res := map[string](*image.RGBA){}
    a := getPic("tex/alpha.png")
    allStr := []string{
        "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ",
        "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
        ".,\"-()[]!?_ :;/\\$*#%+|=",
        "0123456789",
    }
    for iInd, iGird := range allStr {
        jInd := 0
        for _, jGird := range iGird {
            res[string(jGird)] = image.NewRGBA(image.Rect(0, 0, w, h))
            for i := 0; i < w; i++ {
                for j := 0; j < h; j++ {
                    res[string(jGird)].Set(i, j, a.At(jInd * w + i, iInd * h + j))
                }
            }
            jInd += 1
        }
    }
    return res
}

func main() {
	flag.BoolVar(&fullscreen, "f", fullscreen, "fullscreen")
	flag.IntVar(&width, "w", width, "width")
	flag.IntVar(&height, "h", height, "height")
	flag.Float64Var(&scale, "s", scale, "scale")
	flag.Parse()

	setup()

    world = append(world, getLevel("lev/level01.png"))

    for i, m := range mon {
        mon[i].pic = getPic("tex/mon" + strconv.Itoa(m.form) + ".png")
        forms[mon[i].form].screamer = getPic("tex/screamer" + strconv.Itoa(m.form) + ".png")
        mon[i].X = m.XStart
        mon[i].Y = m.YStart
        mon[i].alive = true
        mon[i].live = forms[mon[i].form].liveMax
    }

    for i := 0; i < wallsAmount; i++ {
        wallTexture[i] = getPic("tex/wall" + strconv.Itoa(i) + ".png")
    }

    livesPic = getPic("tex/lives.png")
    floor = map[int](*image.RGBA){}
    floor[0] = getPic("tex/floor0.png")
    floor[-1] = getPic("tex/floor1.png")
    ceiling = getPic("tex/ceiling.png")

    alpha = getAlpha(5, 8)
    getBonus(appleBonusForm, 34, "tex/bonus.png", true)
    getBonus(energyBonusForm, 1, "tex/energy.png", false)
    getBonus(livesBonusForm, 1, "tex/lives.png", true)

    retBlock[startBlock] = getPic("tex/start.png")
    retBlock[wall0Block] = wallTexture[0]
    retBlock[wall1Block] = wallTexture[1]
    retBlock[wall2Block] = wallTexture[2]
    retBlock[wall3Block] = wallTexture[3]
    retBlock[wall4Block] = wallTexture[4]
    retBlock[wall5Block] = wallTexture[5]
    retBlock[mon0Block] = mon[0].pic
    retBlock[mon1Block] = mon[1].pic
    retBlock[appleBlock] = bonusInfo[0].pic[0]
    retBlock[livesBlock] = bonusInfo[1].pic[0]
    retBlock[emptyBlock] = floor[0]
    retBlock[energyBlock] = floor[-1]

    log.Println("Start!")
    pixelgl.Run(run)
}
