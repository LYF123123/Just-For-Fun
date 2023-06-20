package factorial
import "testing"
func TestFactorial(t *testing.T){
	num:=10
	t.Logf("Factorial1: %d",Factorial1(num))
	t.Logf("Factorial2: %d",Factorial2(num))
	t.Logf("Factorial3: %d",Factorial3(num))
	t.Logf("Factorial4: %d",Factorial4(num))
}