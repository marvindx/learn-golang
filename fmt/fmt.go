package main

import (
	"fmt"
	"os"
)

func main() {
	// func Print(a ...interface{}) (n int, err error) {}
	// n = len(a)
	fmt.Print("fmt.Print(): not format, no line break")

	// func Println(a ...interface{}) (n int, err error) {}
	// n = len(a + '\n')
	fmt.Println("fmt.Println(): line break")

	// func Printf(format string, a ...interface{}) (n int, err error) {}
	fmt.Printf("%s: format string, no line break \n", "fmt.Printf()")

	// func Errorf(format string, a ...interface{}) error {}
	err := fmt.Errorf("%s: format, return error", "fmt.Errorf") // return errors.New(s string)
	fmt.Printf("err: %v\n", err)

	file, _ := os.OpenFile("log.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)

	// func Fprint(w io.Writer, a ...interface{}) (n int, err error) {}
	_, _ = fmt.Fprint(file, "fmt.Fprint(): write to file, no line break\n")

	// func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {}
	_, _ = fmt.Fprintln(file, "fmt.Fprintln(): write to file, line break")

	// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {}
	_, _ = fmt.Fprintf(file, "%s: write to file, format, no line break\n", "fmt.Fprintf()")

	// func Sprint(a ...interface{}) string {}
	s := fmt.Sprint("fmt.Sprint(): ", "connet string, ", "no line break, ", "return new string")
	fmt.Println(s)

	// func Sprintln(a ...interface{}) string {}
	s = fmt.Sprintln("fmt.Sprintln(): ", "connet string, ", "line break, ", "return new string")
	fmt.Println(s)

	// func Sprintf(format string, a ...interface{}) string {}
	s = fmt.Sprintf("%s: %s, %s, %s, %s\n", "fmt.Sprintf()", "connet string", "format", "no line break", "return new string")
	fmt.Println(s)
}
