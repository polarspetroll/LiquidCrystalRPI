package LiquidCrystalRPI

import (
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

func NewLCD(addr uint16) (lcd LCD, err error) {
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

func (l LCD) Print(txt string, line int) {
	l.write(LCD_LINES[line], 0)
	for _, v := range txt {
		l.write(byte(int(v)), 1)
	}
}

func (l LCD) Clear() {
	l.write(0x01, 0)
}

func (l LCD) BackLightOff() {
	l.Device.Write([]byte{0x00})
}

func (l LCD) BackLightOn() {
	l.Device.Write([]byte{0x08})
}
