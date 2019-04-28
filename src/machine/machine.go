package machine

type PinConfig struct {
	Mode PinMode
}

type Pin int8

func (p Pin) High() {
	p.Set(true)
}

func (p Pin) Low() {
	p.Set(false)
}

type PWM struct {
	Pin uint8
}

type ADC struct {
	Pin uint8
}
