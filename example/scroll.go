package main

import (
	"time"

	lcd "github.com/polarspetroll/LiquidCrystalRPI"
)

func main() {
	l := lcd.DefaultLCD
	delay := 300 * time.Millisecond
	for i := 0; i < 2; i++ {
		l.ScrollText("Long Text : github.com/polarspetroll/LiquidCrystalRPI", 1, delay)
	}
	l.Clear()
	for i := 0; i < 2; i++ {
		l.ScrollText("Short Text!", 2, delay)
	}

}
