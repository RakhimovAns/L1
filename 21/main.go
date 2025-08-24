package main

import (
	"fmt"
)

// Проблема: имеется несовместимый интерфейс существующего класса (OldPaymentGateway), который не соответствует ожидаемому клиентскому интерфейсу PaymentProcessor.
// Клиент ожидает метод ProcessPayment(amount float64), но старый шлюз предоставляет только PayAmount(amount float64).

// Решение: создаём адаптер (PaymentAdapter), который реализует целевой интерфейс PaymentProcessor и преобразует вызовы ProcessPayment в вызовы PayAmount старого платежного шлюза.

// Плюсы:
// Позволяет использовать несовместимые интерфейсы без изменения исходного кода.
// Упрощает интеграцию сторонних библиотек.
// Изолирует клиентский код от конкретных реализаций.

// Минусы:
// Увеличивает количество классов/структур в проекте.
// Может замедлять выполнение при большом количестве адаптеров (не критично для большинства случаев).

type OldPaymentGateway struct{}

func (o *OldPaymentGateway) PayAmount(amount float64) {
	fmt.Printf("Оплачено через старый шлюз: %.2f₽\n", amount)
}

type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

type PaymentAdapter struct {
	oldGateway *OldPaymentGateway
}

func (p *PaymentAdapter) ProcessPayment(amount float64) {
	p.oldGateway.PayAmount(amount)
}

func main() {
	var processor PaymentProcessor

	oldGateway := &OldPaymentGateway{}
	processor = &PaymentAdapter{oldGateway}

	processor.ProcessPayment(1500.50)
}
