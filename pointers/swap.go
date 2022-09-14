// A program that can swap to integer

/* 

#include<stdio.h>
void swap(char *str1, char *str2)
{
  char *temp = str1;
  str1 = str2;
  str2 = temp;
}  
 
int main()
{
  char *str1 = "geeks";
  char *str2 = "forgeeks";
  swap(str1, str2);
  printf("str1 is %s, str2 is %s", str1, str2);
  getchar();
  return 0;
}

*/
package main

import "fmt"

func swap(x, y *int) {
    //var temp *int = x
    *x = *x
    *y = *y
}

func main() {
    x := 90
    y := 80
    fmt.Println("x:", x)
    fmt.Println("y",y)
    swap(&x, &y)
    fmt.Println("x:",x,"y:",y)
}
