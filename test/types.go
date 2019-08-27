package main



/*
#include "stdio.h"

void test(int n) {
  char dummy[1024];

  printf("in c test func iterator %d\n", n);
  if(n <= 0) {
    return;
  }
  dummy[n] = '\a';
  test(n-1);
}
#cgo CFLAGS: -g
*/

import "C"

type T struct {
}

func (t T) getName() string {
	return "mao"
}
func main() {
	C.test(C.int(20))


}