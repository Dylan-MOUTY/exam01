package main 

import "fmt" //Ajout de la librairie "fmt"

func Surprise(s string) bool {     //
if len(s) < 1 {         // Si (s) est supèrieur a 1 alors.
	return false 
    }     return surprise_annexe(s[0]) || // Si inferieur a 1 la commande ne fera rien
	Surprise(s[1:]) //Surprise
}  func surprise_annexe(a byte) bool {     return (a <= 'z' && a >= 'a') || (a <= 'Z' && a >= 'A') } //Alors effectue l'alphabet en majuscule ou minuscule a l'envers

//Exercice 1