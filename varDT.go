package main

import 
	"fmt"


//const a string = "this is constant"
func main() {
	
	//fmt.Println("hello world")
	/* Dynamic type */

	/*y := 4.5666
	fmt.Println(y)
	fmt.Printf("y is type of %T",y)*/

	/* Mixed data type */

	/*var a,b,c,d =3,3.454,"qwerty table",45
	fmt.Println("\n",a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Printf("\n a is type of %T",a)
	fmt.Printf("\n b is type of %T",b)
	fmt.Printf("\n c is type of %T",c)
	fmt.Printf("\n d is type of %T",d)*/
	
	
	// constant

	/*fmt.Println(a)
	const n  = 500000000
	fmt.Println(n,"\n")
	const m = 3e20 / n
	fmt.Println("value of m =",int64(m))
	fmt.Println(int64(m))
	fmt.Println(math.Sin(n))*/

	// program to convert your weight on earth to mars


	var( 
		a float32 
	 	b int
	)
	fmt.Print("Enter your weight in kg:")
	fmt.Scanln(&a)
	fmt.Print("Enter your age:")
	fmt.Scanln(&b)
	fmt.Print("my weight on mars is :")
	fmt.Print(a*0.3783)
	fmt.Print(" lbs,i would be ")
	fmt.Print(b* 365/687)
	fmt.Print(" years old")



}