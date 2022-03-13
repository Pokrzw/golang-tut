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
deklarowanie zmiennej, do ktorej potem przypiszemy wartosc: var nazwa typ_zmiennej (np var name string)
drukowanie typu zmiennej: Printf("%T", nazwa_zmiennej)

https://www.youtube.com/watch?v=yyUHQIec83I
41:19