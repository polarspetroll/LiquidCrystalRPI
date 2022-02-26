package main

import (
	"log"

	lcd "github.com/polarspetroll/LiquidCrystalRPI"
)

var Pi []byte = []byte{
	0b00000,
	0b11111,
	0b01010,
	0b01010,
	0b01010,
	0b01010,
	0b01010,
	0b10011,
}
var Random []byte = []byte{
	0b01010,
	0b10101,
	0b01010,
	0b10101,
	0b01010,
	0b10101,
	0b01010,
	0b10101,
}
var Cactus []byte = []byte{
	0b00100,
	0b00101,
	0b10101,
	0b10101,
	0b11111,
	0b00100,
	0b00100,
	0b00100,
}
var Heart []byte = []byte{
	0b00000,
	0b01010,
	0b11111,
	0b11111,
	0b01110,
	0b00100,
	0b00000,
	0b00000,
}

var bell []byte = []byte{
	0b00100,
	0b01110,
	0b01110,
	0b01110,
	0b11111,
	0b00000,
	0b00100,
	0b00000,
}

var Skull []byte = []byte{
	0b00000,
	0b01110,
	0b10101,
	0b11011,
	0b01110,
	0b01110,
	0b00000,
	0b00000,
}

var Lock []byte = []byte{
	0b01110,
	0b10001,
	0b10001,
	0b11111,
	0b11011,
	0b11011,
	0b11111,
	0b00000,
}

func main() {
	l, err := lcd.NewLCD(0x27, 16, 2)
	if err != nil {
		log.Fatal(err)
	}

	l.CreateCustomChar(0, Heart)
	l.CreateCustomChar(1, bell)
	l.CreateCustomChar(2, Skull)
	l.CreateCustomChar(3, Lock)
	l.CreateCustomChar(4, Cactus)
	l.CreateCustomChar(5, Random)
	l.CreateCustomChar(6, Pi)

	l.SetCursor(0, 0)
	l.Write(0)

	l.SetCursor(2, 0)
	l.Write(1)

	l.SetCursor(4, 0)
	l.Write(2)

	l.SetCursor(6, 0)
	l.Write(3)

	l.SetCursor(8, 0)
	l.Write(4)

	l.SetCursor(10, 0)
	l.Write(5)

	l.SetCursor(12, 0)
	l.Write(6)

}
