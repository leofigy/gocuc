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
