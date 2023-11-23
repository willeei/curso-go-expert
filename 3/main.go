package main

const a = "Hello World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Williams"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
}
