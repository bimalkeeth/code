package decorator

import (
	"strings"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza:=&PizzaDecorator{}
	pizzaResult,_:=pizza.AddIngredient()
	expectedText:="pizza with following ingredients:"
	if !strings.Contains(pizzaResult,expectedText){
		t.Errorf("when calling the add ingredient of the pizza decorator it must return the text %s expected text,not '%s'",pizzaResult,expectedText )
	}
}


func TestOnion_AddIngredient(t *testing.T) {
	onion:=&Onion{}
	onionResult,err:=onion.AddIngredient()

	if err!=nil{
		t.Errorf("when calling the add ingredient of the onion decorator without an ingredientAdd on its ingredient field must return error,not a string with %s ",onionResult )
	}

	onion=&Onion{&PizzaDecorator{}}
	onionResult,err=onion.AddIngredient()
	if err!=nil{
		t.Error(err)
	}
	if !strings.Contains(onionResult,"onion"){
		t.Errorf("when calling the add ingredient of the onion decorator it must return the text with the word 'onion',not '%s' ",onionResult )
	}
}