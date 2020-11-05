package inline

/*
#include <stdio.h>
void sayHello(){
    printf("hello world!");
}
*/
import "C" // 切勿换行再写这个

func SayHelloByC() {
	C.sayHello()
}
