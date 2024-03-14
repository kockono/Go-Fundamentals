- [Instalación](#instalación)
- [Copilar código](#copilar-código)
- [Inicio](#inicio)
- [fmt documentación](#fmt-documentación)
- [Entradas de teclado](#entradas-de-teclado)
- [Conversores string](#conversores-string)
- [Metodos Math](#metodos-math)
- [If](#if)
- [for (while)](#for-while)
- [Write in files](#write-in-files)
- [Read file](#read-file)
- [Handling Errors](#handling-errors)
- [Packages](#packages)
- [Pointers](#pointers)
- [Structs (Objects)](#structs-objects)
  - [Relacionar Structs](#relacionar-structs)
  - [Metodos de Structs](#metodos-de-structs)
- [Types](#types)
- [Interfaces](#interfaces)
- [Switch](#switch)
- [Any (Generics)](#any-generics)
- [Arrays (Lists)](#arrays-lists)
  - [Slices](#slices)


### Instalación
https://go.dev/dl/

### Copilar código
```sh
go run .
```


### Inicio
```go
fmt.Print("Hello, World!")    : Require import "fmt", Imprimir en consola
:=, var                       : Maneras de declaras variables
"%.2f"                        : Permite imprimir 2 decimales
fmt.Printf("Title %v artist: %v", titulo, artista) : En el orden donde van los params
```

### fmt documentación
https://pkg.go.dev/fmt


### Entradas de teclado
```go
fmt.Scan() : Para entradas de teclado pequeñas
reader, err := bufio.NewReader(os.Stdin) // Entrada grande de teclado reuire "bufio"
text, err   := reader.ReadString('\n')   // leer el binary
```

### Conversores string
```go
strconv.ParseFloat(texto, 64)
// String import "strings"
strings.Clone
strings.Join
strings.Repeat
strings.Map
strings.TrimSuffix(text, "textToDelete") : Texto a eliminar
```

### Metodos Math
```go
float64()  : Convierte en flotante float
```

### If
```go
if choice == 1 && choice == 2
if choice == 1 || choice == 2

if choice == 1 {
	
} else if {

} else {

}
// metodo ternario en go Retorna un boleano
ternario :=  choice == 1

```

### for (while)
```go
// Permite que sea un ciclo infinito
for {
 return // Permite salir del for
 break  // Permite salir del for
 continue // Inicia de nuevo al inicio del for pero no se sale
}
// El quivalente a un while
for tuCondicion { }
// For con condicion de contador
for i := 0; i < 200; i++ { }
```

### Write in files
```go
import ("os")
os.WriteFile("name_file.txt", []byte(unTexto), 0644) // 0644 is the permission
```

### Read file
```go
data, err := os.ReadFile("balance.txt")
if err != nil {
	fmt.Println(err)
}
```

### Handling Errors
```go
errors.New("Failded to parsed the value") : Muestra un error
panic()                                   : Cierra la aplicación totalmente

// Retorna nil si sale un error
func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required.")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}
```

### Packages
```go
// apis.go
package main
import("./utils")

func apis() {

}
// main.go
package main

func main() {

}
// utilis.go, si quieres otro package name meterlo en una carpeta y mayuscula
// Equivale a un export la Mayuscula
package utils

func ShowAlert() {

}
```
Nombramos de la misma manera que el package


### Pointers
Un puntero tiene un valor de la memoria
& obtiene la dirección de memoria 0x55324234

Unos permite evitar copiar valores inecesarios

**Ejemplos**
```go
var agePointer *int = &age // Pointer Variable
age pointer := age         // Manera corta
*agePointer                // 32 hace referencia al valor
metodo(&variable)          // Permite acceder al puntero de la variable
(*objeto).propiedad        // Una manera de acceder al puntero de un struct 
```
Podemos modificar la variable en cualquier lugar sin tener que retornarla

### Structs (Objects)
```go
type User struct {
    Username string
    Age      int
    createAt time.Time
}
```
Usar mayúscula para que sea publico o se pueda acceder desde otros archivos
#### Relacionar Structs
```go
type User struct {
    firstName string
    age      int
    createAt time.Time
}

type Admin struct {
	username string
	password string
	User
}

func newAdmin(email, password string) Admin {
	return Admin {
	email: "admin@gmail.com",
	password: 124455,
	User: User {
		firstName: "Admin",
		age: 19,
		createdAt: time.Now()
	}
  }
}
```
Ahora tengo un admin con las propiedades de **User**
#### Metodos de Structs
```go
func (u User) outputUserDetails() {
	fmtPrint(u.Username)
}
func (u *User) clearUserName() {
	u.Username = ""
	u.Age = 0
}
// Crea copias, no usa los datos originales entonces no mutan
appUser.outputUserDetails()
appUser.clearUserName()
```
Una manera de enlazar métodos con la estructura, si requerimos modificar un valor ocupamos el puntero

### Types
```go
package main

type str string
type f64 float64

func main() {
	var name str
}
```
permite crear otros types, similar a remapear tipos con otras palabras, usar **var**

### Interfaces
```go
type saver interface {
	Save() error
}

type outputsData interface {
	saver
	Display()
}

func  getData(value interface{}) {} // Any puede ob tener cualquier cosa
```
Permite crear una interface de los posibles outputs 

### Switch
```go
func printSomething(value interface{}) {
	switch value.(type) {
		case int:
			fmt.Println("Integer:", value)
		case float64
			fmt.Println("Integer:", value)
		case string:
			fmt.Println(value)
	}
}

switch choice {
        case 1:
            fmt.Println("Your balance is", accountBalance)
        case 2:
            fmt.Println("Your balance is", accountBalance)
        default:
            fmt.Println("Goodbye!")
        return
}
```
Permite imprimir cualquier valor

### Any (Generics)
```go
func add(a, b, interface {}) interface {} {

}

// Type placholder
func add[T any](a,b T) T {

}
// Para multiples valores
func add[T int | float64 | string](a, b T) T {
	return a + b
}

```
Permite volver un método totalmente dinámico

### Arrays (Lists)
```go
prices := [4]float64{10,99,1.55,10.97,2.0} // Guardar 4 valores
var products [3]string  // Otra manera de declararlo

// Imprimir valores
prices[3] // Salida: 2.0
```
#### Slices
```go
priceRecortado  = prices[1:3] // Salida: 10.99, 1.55, 10.97
priceRecortado2 = prices[:3]  // Salida: 10.97
```

