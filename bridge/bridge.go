package bridge

import (
	"errors"
	"fmt"
	"io"
)

type PrinterApi interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct {

}

func(p *PrinterImpl1)PrintMessage(msg string)error{
	fmt.Printf("%s\n",msg)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl2)PrintMessage(msg string) error{
	if p.Writer==nil{
		return errors.New("you need to pass an io.writer to printerimpl2")
	}
	_,_= fmt.Fprintf(p.Writer,"%s",msg)
	return nil
}

type PrinterAbstraction interface {
	Print() error

}

type NormalPrinter struct {
	Msg string
	Printer PrinterApi
}

func(p *NormalPrinter) Print() error{
	_=p.Printer.PrintMessage(p.Msg)
	return nil
}

type PacktPrinter struct {
	Msg string
	Printer PrinterApi
}

func(p *PacktPrinter)Print()error{
	_=p.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s",p.Msg))
	return nil
}