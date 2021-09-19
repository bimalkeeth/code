package decorator

import (
	"errors"
)

type IngredientAdd interface {
	AddIngredient()(string,error)
}

type PizzaDecorator struct {
	Ingredient IngredientAdd
}

func(p *PizzaDecorator)AddIngredient()(string ,error){
	return "pizza with the following ingredient",nil
}

type Meat struct {
	Ingredient IngredientAdd
}

func(p *Meat)AddIngredient()(string ,error){
	return "",errors.New("not implemented yet")
}

type Onion struct {
	Ingredient IngredientAdd
}

func(p *Onion)AddIngredient()(string ,error){
    if p.Ingredient==nil{
		return "",errors.New("an ingredient added is needed in the ingredient field of the onion")
	}

	s,err:=p.Ingredient.AddIngredient()
	if err !=nil{
		return "",err
	}

	return s,nil
}