# LiquidCrystalRPI
A simple LCD controller package for raspberry pi liquid crystal I²C displays.

## Example

```go
import (
  "log"
  "time"

  lcd "github.com/polarspetroll/LiquidCrystalRPI"
)


func main() {
  l, err := lcd.NewLCD(0x27, 16, 2) // specify the I²C device address, width and height
  if err != nil {
    log.Fatal(err)
  }

  /* You can also use the default configuration like so :

  l := lcd.DefaultLCD
  */
  l.Print("Hello World!", 1) // print 'Hello World' at line 1
  l.Print("Second Line", 2) // print 'Second Line' at line 2
  time.Sleep(3 * time.Second)
  l.Clear() // clear the display
  time.Sleep(3 * time.Second)
  l.BackLightOff() // turn off backlight
  time.Sleep(2 * time.Second)
  l.BackLightOn() //turn on backlight
}

```
