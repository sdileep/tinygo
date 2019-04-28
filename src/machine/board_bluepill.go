// +build stm32,bluepill

package machine

// https://wiki.stm32duino.com/index.php?title=File:Bluepillpinout.gif
const (
	PA0  Pin = portA + 0
	PA1  Pin = portA + 1
	PA2  Pin = portA + 2
	PA3  Pin = portA + 3
	PA4  Pin = portA + 4
	PA5  Pin = portA + 5
	PA6  Pin = portA + 6
	PA7  Pin = portA + 7
	PA8  Pin = portA + 8
	PA9  Pin = portA + 9
	PA10 Pin = portA + 10
	PA11 Pin = portA + 11
	PA12 Pin = portA + 12
	PA13 Pin = portA + 13
	PA14 Pin = portA + 14
	PA15 Pin = portA + 15
	PB0  Pin = portB + 0
	PB1  Pin = portB + 1
	PB2  Pin = portB + 2
	PB3  Pin = portB + 3
	PB4  Pin = portB + 4
	PB5  Pin = portB + 5
	PB6  Pin = portB + 6
	PB7  Pin = portB + 7
	PB8  Pin = portB + 8
	PB9  Pin = portB + 9
	PB10 Pin = portB + 10
	PB11 Pin = portB + 11
	PB12 Pin = portB + 12
	PB13 Pin = portB + 13
	PB14 Pin = portB + 14
	PB15 Pin = portB + 15
	PC13 Pin = portC + 13
	PC14 Pin = portC + 14
	PC15 Pin = portC + 15
)

const (
	LED = PC13
)

// UART pins
const (
	UART_TX_PIN = PA9
	UART_RX_PIN = PA10
)

// SPI pins
const (
	SPI0_SCK_PIN  = PA5
	SPI0_MOSI_PIN = PA7
	SPI0_MISO_PIN = PA6
)

// I2C pins
const (
	SDA_PIN = PB7
	SCL_PIN = PB6
)
