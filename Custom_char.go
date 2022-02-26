package LiquidCrystalRPI

import "time"

/*
FUNCTIONS IN THIS FILE ARE A CLONE OF Arduino-LiquidCrystal-I2C-library
GITHUBE : https://github.com/fdebrabander/Arduino-LiquidCrystal-I2C-library
*/

func (l LCD) CreateCustomChar(location uint8, b []byte) {
	location &= 0x7
  l.command(0x40 | (location << 3))
  for i := 0; i < 8; i++ {
		l.write2(b[i])
	}
}

func (l LCD) SetCursor(col, row uint8) {
  var row_offsets []uint8 = []uint8{0x00, 0x40, 0x14, 0x54}
  if int(row) > l.Width {
		row = uint8(l.Width-1)
	}
  l.command(0x80 | (col + row_offsets[row]))
}

func (l LCD) Write(b uint8) {
  l.send(b, 1)
}

/*-------------------------------------------------------------------------*/

func (l LCD) command(value uint8) {
  l.send(value, 0)
}

func (l LCD) send(value, mode uint8) {
  high := value&0xf0
  low := (value<<4)&0xf0
  l.write4bits((high)|mode)
	l.write4bits((low)|mode)
}

func (l LCD) write4bits(value uint8) {
  l.expanderwrite(value)
	l.pulse_enable(value)
}

func (l LCD) expanderwrite(data uint8) {
  l.Device.Write([]byte{byte((data) | 0x08)})
}

func (l LCD) pulse_enable(data uint8) {
  l.expanderwrite(data | 4)
  time.Sleep(time.Microsecond)
  l.expanderwrite(uint8(int(data) & ^4))
  time.Sleep(50 * time.Microsecond)
}

func (l LCD) write2(value uint8) {
  l.send(value, 1);
}
