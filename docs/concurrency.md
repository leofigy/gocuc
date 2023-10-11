### Un paseo por la concurrencia

Go ofrece diferentes formas de trabajar con concurrencia, a diferencia de lenguajes 
tradicionales donde hay ejecutores, pools de hilos, y la sincronizacion se vuelve tediosa.

Go intenta evitar el problema llamado "shared mutable state" (compartiendo el estado mutable),
a traves de comunicarse por medio de mensajes. 

Las principales herramientas que tenemos son:

- Go routines: Una funcion que puede correr de forma independiente, sobre un hilo ligero o simulado.
- Channel    : Una herramienta que nos permite comunicar a las Go routines, por medio de mensajes. 
- Select     : Una primitiva para sincronizar el acceso a los channels. 

```go 
   // go routine example  
   go func() { fmt.Println("Hello go !") }()
   
   // channel 
   data := make(chan string, 10)
   
   // select statement 
   for {
       select {
         case res := <-data:
           fmt.Println(res)
        case <-time.After(3 * time.Second):
           fmt.Println("timeout 2")
        } 
   }
```

Go al ser un lenguaje con recoleccion de basura (GC), tiene un costo extra en runtime, 
pero nos libra de manejar la liberacion de memoria. 
Ahora la pregunta es como se manejan las go routines?

Contexto:
El runtime de go decide como se va a utilizar el CPU, existen tres tipos de modelo que se utilizan

Kernel Level Thread model (modelo de hilo a nivel de kernel   1:1)  
```
Los hilos son administrados por el kernel por lo tanto se necesita hacer un system call, 
para obtener acceso, sin embargo es costoso estar creando y regresando el control de los hilos al scheduler del sistema operativo, la relacion es 1:1 con respecto a la aplicacion. El beneficio a pesar del costo, como el kernel tiene contexto de los threads, tiene mecanismos para evitar los puntos muertos (deadlock)

```
User Level Thread model   (modelo de hilo a nivel de usuario  N:1)
```
Los hilos a nivel de usuario tienen una relacion de N:1, es decir varios hilos simulados del lado de la aplicacion se terminan ejecutando en un solo hilo del kernel, esto permite tener concurrencia, son mas baratos para crear, sin embargo el kernel no tiene contexto de cuales hilos pueden estar bloqueados del lado del usuario (aplicacion).
```
Hybrid Thread model       (modelo de hilo hibrido N:M)
```
El modelo hibrido consiste en usar los dos tipos de threads desde la aplicacion, desde el lado de la aplicacion se crean los hilos de usuario, ademas que se proveen mecanismos de sincronizacion para evitar los puntos muertos. 
```

### GMP model

Como mencionamos go tiene un recolector de basura, pero ademas tiene un planificador (scheduler) para gestionar las go-routines en tiempo de ejecucion (runtime)

```go
package main

import (
   "time"
   "runtime"
)

func main(){
   for i := 0 ; i < 100; i++ {
      go func() {
         time.Sleep(time.Second)
      }()
   } 
   // NumGoroutine returns the number of goroutines that currently exist. 
   fmt.Println(runtime.NumGoroutine())
}
```
El codigo anterior lanza 100 go-routines , despues en la funcion main que corre en su propia rutina, 
imprime el numero total de go-routines existentes

#### Ahora la pregunta, que hace el scheduler cada vez que creamos una go-routina 1-to-1, N-to-1 o M-to-N (user/kernel threads)?

Pues depende, puede usar cualquiera de las 3 estrategias. GPM significa:
- G: goroutine (rutina de go)
- M: Machine   (maquina, threads disponibles en el sistema , kernel threads, el numero maximo es 10,000)
- P: Processor (es una cola de go-routines esperando a ser distribuidas en cada kernel thread)



