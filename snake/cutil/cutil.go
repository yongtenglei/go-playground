package cutil

/*
#include<stdio.h>
#include"getch.h"

//
void gotoxy(int x,int y)   //Fantasy
{
  printf("%c[%d;%df", 0x1B, y, x);
}

int direction()
{
	return getch();
}

void clrscr(void)
{
  stdio.system("clear");
}

*/
import "C"

func GotoPosition(x int, y int) {
	C.gotoxy(C.int(x), C.int(y))
}

func Direction() int {
	return int(C.direction())
}

func Clrscr() {
	C.clrscr()
}
