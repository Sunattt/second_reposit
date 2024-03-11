package Days13

import (
"fmt"
"math/rand"
)

type SinglQR interface {
	Pay(amount float64) bool
	CheckBalance() float64
	TransferTo(other SinglQR, amount float64) bool
	AcceptTo(other SinglQR, amount float64) bool
	Commission() float64
}

func (w *Customer ) Commission() float64 {
	const com = 0.5 
	w.Balance -= com 
	return w.Balance 
}

func NameBank() *Customer {
	num := rand.Intn(3)
	var name string
	if num == 0 {
	 name = "Alif"
	} else if num == 1 {
	 name = "Humo"
	} else if num == 2 {
	 name = "Eschata"
	}
	return &Customer{
	 Bank: name,
	}
}

func (w *Customer) Pay(amount float64) bool {
	if w.Balance >= amount {
		w.Balance -= amount
		return true
	} else {
		return false
	}
}

func (w *Customer) CheckBalance() float64 {
	const cumitions = 0.10
	w.Balance -= cumitions
	fmt.Printf("У %v вот с только денег на счету %v \n",w.Name, w.Balance)
	return w.Balance
}

func (w *Customer) TransferTo(other *Customer, amount float64) bool {
	if w.Balance >= amount {
		w.Balance -= amount
		other.Balance += amount
		return true
	} else {
		return false
	}
}

func (w *Customer) AcceptTo(other *Customer, amount float64) bool {
	if w.Balance < w.Balance+amount {
		w.Balance += amount
		other.Balance -= amount
		return true
	} else {
		return false
	}
}
