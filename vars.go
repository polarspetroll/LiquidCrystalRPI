package LiquidCrystalRPI

import (
	"periph.io/x/conn/v3/i2c"
	"time"
)

type LCD struct {
	Device *i2c.Dev
	Height int
	Width  int
}

var DefaultLCD, _ = NewLCD(0x27, 16, 2)

const SLEEP = 50 * time.Microsecond

var LCD_LINES = map[int]byte{
	1: 0x80,
	2: 0xC0,
	3: 0x94,
	4: 0xD4,
}
