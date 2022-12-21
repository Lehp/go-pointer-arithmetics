# ptrArit
Allows Arithmetics with pointers in golang 

In web development with golang a variable is not nullable. To solve this problem many go dev's sadly have to use pointers. But go does not support 
pointer-arthmetic. This lib converts your input to floats, does the calculations and returns the result as the type you want it to be. 

If one of the params is nil the function always returns nil to prevent you from getting results you should not have.

The result is always returned as a pointer of whatever type you described. 



Examples: 

asInt := Add[int](float64(13.25), int(2)) -> 15
substract := Substract[int]("13", "2") -> 11

float := Multiply[float64](ptr(string("13")), float64(2)) -> float64(26)
divide := Divide[int]("13", ptr(string("2") -> 6

Disclaimer: 
Unsafe can only handle intlike types and is thus not fitted for real arithmetic.
Additionally you have to convert all your types to unsafe.Pointer before you can use it.


