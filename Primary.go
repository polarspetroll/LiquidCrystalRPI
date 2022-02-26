package LiquidCrystalRPI

import (
	"strings"
	"time"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

func initialize() error {
	_, err := host.Init()
	if err != nil {
		return err
	}
	_, err = driverreg.Init()
	if err != nil {
		return err
	}
	return nil

}

// This function returns an object that represents a LCD.
// The addr is the I²C device address (0x27 is the most common one).
// width and height is the LCD's size.
func NewLCD(addr uint16, width, height int) (lcd LCD, err error) {
	err = initialize()
	if err != nil {
		return LCD{}, err
	}
	bus, err := i2creg.Open("1")
	if err != nil {
		return LCD{}, err
	}
	lcd = LCD{
		Device: &i2c.Dev{Addr: addr, Bus: bus},
		Height: height,
		Width:  width,
	}
	lcd.write(0x33, 0)
	lcd.write(0x32, 0)
	lcd.write(0x06, 0)
	lcd.write(0x0C, 0)
	lcd.write(0x28, 0)
	lcd.write(0x01, 0)
	time.Sleep(SLEEP)
	return lcd, nil
}

// This function prints a string on the LCD.
//txt is the string that you want to print.
//line is the line of the text.
/*
Example:
	func main() {
		...
		lcd.Print("Hello World!", 1) // print 'Hello World' at line 1

		lcd.Print("Second Line", 2) // print 'Second Line' at line 2
	}
*/
func (l LCD) Print(txt string, line int) {
	l.write(LCD_LINES[line], 0)
	for _, v := range txt {
		l.write(byte(int(v)), 1)
	}
}

//Clear, clears the entire screen
func (l LCD) Clear() {
	l.write(0x01, 0)
}

//BackLightOff turns off the LCD's backlight
func (l LCD) BackLightOff() {
	l.Device.Write([]byte{0x00})
}

//BackLightOn turns on the LCD's backlight
func (l LCD) BackLightOn() {
	l.Device.Write([]byte{0x08})
}

// ScrollText iterates over the txt and prints it on the LCD
// txt : Text to be printed
// row : line number
// delay : scroll speed(lower = faster)

/*
Example:
	import (
		"time"

		lcd "github.com/polarspetroll/LiquidCrystalRPI"
	)

	func main() {
		l := lcd.DefaultLCD
		defer l.Close()
		delay := 300 * time.Millisecond
		for i := 0; i < 2; i++ {
			l.ScrollText("Long Text : github.com/polarspetroll/LiquidCrystalRPI", 1, delay)
		}

		l.Clear()

		for i = 0; i < 2; i++ {
			l.ScrollText("Short Text!", 1, delay)
		}

	}

*/
func (l LCD) ScrollText(txt string, line int, delay time.Duration) {
	if len(txt) < l.Width {
		for range txt {
			txt += " "
		}
		l.short_text_scroll(txt, line, delay)
	} else if len(txt) > l.Width {
		l.long_text_scroll(txt, line, delay)
	}
}

func (l LCD) long_text_scroll(txt string, line int, delay time.Duration) {
	var tmp string
	for range txt {
		tmp = string(txt[0])
		txt = strings.Replace(txt, string(txt[0]), "", 1)
		txt += tmp
		l.Print(txt[0:l.Width-1], line)
		time.Sleep(delay)
	}
}

func (l LCD) short_text_scroll(txt string, line int, delay time.Duration) {
	var tmp string
	for range txt {
		tmp = string(txt[0])
		txt = strings.Replace(txt, string(txt[0]), "", 1)
		txt += tmp
		l.Print(txt, line)
		time.Sleep(delay)
	}
}


func (l LCD) write_byte(b byte) {
	l.Device.Write([]byte{b})
	l.Device.Write([]byte{b | 0b00000100})
	time.Sleep(SLEEP)
	l.Device.Write([]byte{b & ^byte(0b00000100)})
	time.Sleep(SLEEP)
}

func (l LCD) write(b, n byte) {
	l.write_byte(n | (b & 0xF0) | 0x08)
	l.write_byte(n | ((b << 4) & 0xF0) | 0x08)
}
