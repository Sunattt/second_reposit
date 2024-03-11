package Days13

type Customer struct {
	Name    string
	Surname string
	Age     int
	Bank    string
	Balance float64
}

type Alif struct {
	Client Customer
	Qr     SinglQR
	Name   string
}

type Humo struct {
	client Customer
	Name   string
	Qr     SinglQR
}
