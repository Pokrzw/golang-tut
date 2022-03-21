# golang-tut
1-szy program w golangu

Tworzenie modulu (terminal):
>go mod init nazwa-projektu

W go wszystko jest zorganiozwane w packages, pierwszy statement w go musi byc package.
Main function to bd entrypoint z ktorego go zacznie kompilowac nasz program

Executing go app:
>go run file_name.go

```go
package henlo

import "fmt"

func henlo() {
	fmt.Println("Henlo owo")
}
```

zmienna: var   
stała: const   
deklarowanie zmiennej, do ktorej potem przypiszemy wartosc: 
>var nazwa typ_zmiennej (np var name string)    

drukowanie typu zmiennej:     
>Printf("%T", nazwa_zmiennej)

W GO, deklarowanie typu zmiennej przy przypisywaniu wartosci ma wiekszy sens, gdy mamy na mysli inny typ zmiennej, niż taki, który zostałby wykryty przez GO
```go
var name string = "name"
const numberOfSth uint = 50 //pozwala tylko na nieujemne wartosci
```

inny sposob na deklarowanie zmiennych (nie da sie w ten sposob zadeklarowac const i typu):
```go
zmienna := 50
```

### Pointers
>pozwalaja na uzyskanie oryginalnej wartosci zmiennej zamiast kopii 

```go
//& -> symbol pointera

fmt.Println(zmienna) //drukuje wartosc
fmt.Println(&zmienna) //drukuje adres w pamieci
```


### User input
```go
fmt.Scan(&nazwaZmiennej)
//za pomoca pointera mozna przypisac wartosc 
//zmiennej, poniewaz teraz ma adres zmiennej w pamieci
```

## arrays
>Arrays in go have fixed size, trzeba deklarowac ich rozmiar przy deklaracji tablicy za pomoca []
>tablice maja tez okreslony typ danych ktore przyjmuja
```go
var tablica [50]string //zakladamy ze w tej tablicy bd tylko 50 elementow o typie string, pusta tablica
var tablica = [50]string

var tablica2 [50]string{'a','b','c'} //tablica z poczatkowymi elementami
```
>dodajemy elementy do tablicy poprzez podanie indexu tablicy:
```go
tablica[0] = 'a'
tablica[5] = 'e'
```
>wyswietlanie tablic:
```go
	fmt.Printf("The whole array: %v\n", tablica) 
	fmt.Printf("0 element of the array: %v\n", tablica[0])
	fmt.Printf("Type of the array: %T\n", tablica) 
	fmt.Printf("Array length: %v\n", len(tablica)) 
```
## Slices
>jak arrays ale jak nie wiemy jaka bedzie koncowa ilosc elementow
>array type but with dynamic size
```go
var slice1 = []string{} //definiujemy jak array, ale bez zdefiniowanego rozmiaru
var slice1 [] string
slice1 := []string{}

```
>dodawanie elementow do slices:
```go
slice1 = append(slice1, 'element')
```
>wyswietlanie slices:
```go
	fmt.Printf("The whole slice: %v\n", slice) 
	fmt.Printf("0 element of the slice: %v\n", slice[0])
	fmt.Printf("Type of the slice: %T\n", slice) 
	fmt.Printf("Slice length: %v\n", len(tablica)) 
```

## Loops in go
>there is ONLY for loop in go
Konstrukcja for loop:
```go
for{ //nieskonczona petla for
	instrukcje
}

for i := 0; i < count; i++ {
	instrukcje
}
```
>loopowanie przez elementy tablicy:
```go
slice2 := []string{}
for index, booking := range bookings  {
	instrukcje			
}

//jezeli nie bedziemy uzywac w petli indexu, mozemy zastapic go _:
for _, booking := range bookings{
	...
}
```
Na kazda iteracje dostajemy 2 wartosci: 
1. Index
1. Aktuany element

Po := podajemy liste, ktorej wartosci bedziemy iterowac

Range pozwala iterowac przez slice

>Range iterates over elements for different data structures (so not only arrays and slices)
_ pozwala na ignorowanie zmiennych, ktore nie beda uzywane
## Importowanie wielu paczek:
```go
import(
	"a"
	"b"
)
```

## if statemens
```go
if condition {

}
```

W go w conditionalach mozna miec dowolna ilosc else if, ale tylko jedno if i tylko jedno else

pare if z rzedu: moga sie wykonac wszystkie
if + else if z rzedu: wykona sie tylko jedno
## For with conditions:
przyklad:
```go
for zmienna > 0 {
	...
}
```
## Walidacja inputow

```go
strings.Contains(email, "@")
```
## Switche
```go

zmienna := "costam"

switch zmienna{
	case "1":
		//instrukcje
	case "2", "3":
		//instrukcje
	default:
		//Instrukcje
}
```

## Funkcje

```go
func add(a int,b int) int{
	return a + b
}
```
funkcje w go moga zwracac wiele wartosci
```go
func multipleRes(a int, b int) (int, int, bool){
	dodawanie := a + b
	mnozenie := a * b
	porownanie := a > b
	return dodawanie, mnozenie, porownanie
}
```
## packages
>collection of go files
Tworzenie nowych packages
```go
package main

func nazwa_fx{
	...
}
```
jezeli mamy w innych plikach jakies fx ktore sa wykonywane w main, to uruchamiamy program za pomoca polecenia:
```bash
go run main.go helper.go
```
albo folder zawierajacy te pliki
```bash
go run .
```
Jezeli tworzymy rozne packages (np nie wszystkie pod main np), to najlepiej podzielic je tak, zeby byly w innych folderach

Jesli chcemy uzych funkcji z innego package, to trzeba je zaimportowac

```go
package main

import "nazwa_modulu/wlasna_paczka"

...

wlasna_paczka.Funkcja()
```

w pliku go.mod, jest linijka module nazwa_modulu, ktora decyduje o import path.
Zeby zaimportowac zmienna/funkcje, trzeba jej nazwe napisac z duzej litery

## Maps
>all keys have the same data type (cannot mix data types!)
Tak jak dicts w pythonie
```go
var userData = make(map[string]string)
userData["firstName"] = zmeinna1
userData["lastName"] = zmeinna2
```
w sytuacjach, gdy chcemy zapisac zmienna innego typu, bedziemy musieli przekonwertowac ja na typ, ktory jest przyjmowany przez map (np z package strconv)

making an empty slice of maps:
```go
var sliceOfMaps = make([]map[string]string, initialSize)
```

## Structs
Basically objekt
```go

type User struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var userData = User {
	firstName: "lorem",
	lastName: "ipsum",
	email: "lorem@ipsum.com",
	numberOfTickets: 10
}

///getting values from struct
userData.firstName

```
## concurency
pozwala sie na wykonywanie reszty kodu w miedzyczasie, gdy wykonuje sie kod, 
ktory zajmuje troche czasu, typu laczenie sie z api

>Creating thread in go
zeby stworzyc jakis thread, przy nazwie fx wystarczy dopisac go
```go
func main{
	...
	go funkcja_async_caly_ten{

	}
}
```
zeby uniknac sytuacji, gdzie program konczy sie zanim skonczy sie ktorys z watkow pobocznych, trzeba dodac sync:
```go
var wg = sync.WaitGroup{}

func main{
	//przed zadeklarowaniem nowego thread
	wg.Add(ilosc_deklarowanych_watkow)
	go jakisWatek()
	go.innywatek()

	//to musi byc na koncu glownego thread -> czekanie na zakonczenie sie reszty watkow
	wg.Wait()
}

func jakisWatek(){
	....
	//to jest dodawane na koncu watku
	wg.Done()
}
```

