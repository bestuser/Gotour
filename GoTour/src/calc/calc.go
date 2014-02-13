package main 
import "fmt"
import (
	"os"
	"strconv"
	"simplemath"
)

var Usage = func(str string){
	if(len(str) == 0){
		fmt.Println("Usage:calc command [arguments]...");
		}else{
			fmt.Println(str);
			}
}

func main() {
	args := os.Args;
	fmt.Println(args);
	if(args == nil || len(args) == 0){	
		Usage("");
		return;
		}
	var v1,v2 int;
	var err1,err2 error;	
	switch args[1]{
		case "Add"://加法运算,调用simplemath.Add方法
			if(len(args) < 4){
				Usage("args != 2");
				return
				}
			v1,err1 = strconv.Atoi(args[2]);
			v2,err2 = strconv.Atoi(args[3]);
			if(err1 != nil || err2 != nil){
				Usage("Add [arg1 int],[arg2 int]");
				return;
			}
			var returnV = simplemath.Add(v1,v2);
			fmt.Println(returnV);
		case "Sqrt"://开根号运算,调用simplemath.Sqrt方法
			if(len(args) < 3){
				Usage("arg is not 1");
				return 
				}
			v1,err1 = strconv.Atoi(args[2]);
			if(err1 != nil){
				Usage("Sqrt [arg float64]");
				return;
			}
			var returnV = simplemath.Sqrt(v1);
			fmt.Println(returnV);
		default:
			Usage("no operator");
		
	}
}

