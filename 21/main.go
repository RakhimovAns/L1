package main

import (
	"fmt"
)

// Проблема: у вас есть объект с одним интерфейсом (например, Dog.WoofWoof() или Cat.MeowMeow(bool)), но клиент ожидает другой интерфейс (например, Reaction()).

// Решение: создаём адаптер, который реализует нужный интерфейс и делегирует вызовы к «несовместимому» объекту.

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
