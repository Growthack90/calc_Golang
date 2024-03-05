package main

import (
    "fmt"
   "calcolatrice/operazioni/addizione"
   "calcolatrice/operazioni/sottrazione"
   "calcolatrice/operazioni/moltiplicazione"
   "calcolatrice/operazioni/divisione"

)

func main() {

    Mymath_Operation()
    
    sum := addizione.Aggiungi(5, 7)
    fmt.Println("La somma è:", sum)
    
    sub := sottrazione.Sottrai(5, 7)
    fmt.Println("La sottrazione è:", sub)
    
    mul := moltiplicazione.Moltiplica(5, 7)
    fmt.Println("La moltiplicazione è:", mul)
    
    div := divisione.Dividi(5, 7)
    fmt.Println("La divisione è:", div)

}