package main

import (
	"fmt"
	"math/rand"
	"second_reposit/Days11"
	"second_reposit/Days13"
	"second_reposit/Days8"
)

func main() {

	//выбор дня
	var day int
	_, err := fmt.Scan(&day)
	//проверка на содержания ошибки при печати
	if err != nil {
		fmt.Println("Вы ввели неправильное число !!!")
	}
	switch day {
	case 8, 9:
		//Создаём
		fmt.Println()
		z := Days8.CreateM()
		fmt.Println(z)

		//Читаем
		fmt.Println()
		Days8.ReadM(z)

		//Модифицыруем изменяем
		fmt.Println()
		Days8.UpdateM(z)

		x := Days8.UpdateM(z)

		fmt.Println()
		Days8.ReadM(z)

		fmt.Println()
		Days8.ReadM(x)

		Days8.DeleteM(z)

	case 11:
		// функции для переобразования строки в byte
		var array string = "My name is Sunat"
		Days11.TransStr(array)

		// функция для переобразования byte в string
		slice := []byte{106, 119, 110, 108, 103, 114, 106}
		Days11.TransByte(slice)

		//шифт Цезоря
		text := "cwvmoskoswmkws svnodvspmvp"
		Days11.Caezar(text)

		//замена символов "+" на " "
		str := ". Мы+получили+текст,+в+котором+по+какой-то+ошибке+все+пробелы+заменились+на+плюсы.+Надо+это+исправить"
		Days11.ReplaceSim(str)

		//можно указать определённый симовл который надо удалить 
		// var words string
		// fmt.Scan(&words)
		//Days11.ReplaceSim(words)
	case 13:
		//создаем двух клиентов 
		var user Days13.Customer = Days13.Customer{"Vladimir", "Nikolaevich", 23, Days13.NameBank().Bank, 1000.0}
		var user2 Days13.Customer = Days13.Customer{"Mahmud", "Karimovich", 24, Days13.NameBank().Bank, 1234.0}
		//проверка всего что есть у клиентов 
		fmt.Printf("%+v\n", user) 
		fmt.Printf("%+v\n", user2)

		fmt.Println() 
		fmt.Println()

		//Метод Pay принимает цену за какой-то товар,за который должны оплатить 
		user.Pay(34)
		//Метод TransferTo это перевод день из одного кошелька на другой
		user2.TransferTo(&user, 234)

		fmt.Println(user) 
		fmt.Println(user2)
		//Этот метод для получения денег(когда деньги переводят к нам ) 
		user2.AcceptTo(&user, 10000)
		//проверка баланса за определённую константную комиссию 
		user.CheckBalance()
		user2.CheckBalance()
		fmt.Println() 
		fmt.Println()
		//Что у нас получилось в итоге 
		fmt.Println(user) 
		fmt.Println(user2)

		// Сконструируем какую нибудь ситуацию 
		var cl1 = Days13.Customer{Name: "Arseniy", Surname: "Evklidovich", Age: 25, Bank: Days13.NameBank().Bank,Balance:  123.0 }
		var cl2 = Days13.Customer{Name: "Misha", Surname: "Ivanovich", Age: 22, Bank: Days13.NameBank().Bank,Balance:  70.0 }

		//ресторан
		restBill := rand.Intn(200)/2
		var nums int
		_, err :=  fmt.Scan(&nums)
		if err != nil{
			fmt.Println("Вы ввели неправильное число !!!")
		}
		switch {
		case nums == 1 && restBill <= 70 :
			fmt.Printf("Делим счет попалам. Заплати мне %d сомон\n", restBill)
			cl2.Pay(float64(restBill))
			cl1.Pay(float64(restBill))
			
			cl1.CheckBalance()
			cl2.CheckBalance()

		case nums == 2 && restBill > 70:
			fmt.Printf("%s можешь дать мне в долг %v сомон\n",cl1.Name, restBill-70 )
			fmt.Printf("Да конечно %s, я тебе отправлю на счет.\n",cl2.Name)
			//перевод деньг с cl1 на cl2
			cl1.AcceptTo(&cl2, float64(restBill-70))
			if cl1.Bank == cl2.Bank{
				//елси имя баннков совподает то коммисии нет 
				cl1.CheckBalance()
				cl2.CheckBalance()

				fmt.Printf("%+v\n",cl1 ); fmt.Printf("%+v\n",cl2 )
			}else{
				//взымается коммиссия за перевод
				cl1.Commission()
				cl1.CheckBalance()
				cl2.CheckBalance()

				fmt.Printf("%+v\n",cl1 ); fmt.Printf("%+v\n",cl2 )
			}

			
		case nums == 3 && restBill > 70:
			fmt.Printf("%s можешь дать мне в долг %v\n",cl2.Name, restBill-70 )
			fmt.Printf("Да конечно %s, я тебе отправлю на счет.\n",cl1.Name)
			//перевод деньг с cl1 на cl2
			cl2.TransferTo(&cl1, float64(restBill-70))
			
			if cl1.Bank == cl2.Bank{
				//елси имя баннков совподает то коммисии нет 

				cl1.CheckBalance()
				cl2.CheckBalance()

				fmt.Printf("%+v\n",cl1 )
				fmt.Printf("%+v\n",cl1 )
			}else{
				//взымается коммиссия за перевод
				cl2.Commission()

				cl1.CheckBalance()
				cl2.CheckBalance()

				fmt.Printf("%+v\n",cl1 )
				fmt.Printf("%+v\n",cl1 )
			}
		}

	}
}
