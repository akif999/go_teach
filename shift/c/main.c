// This is main.c

#include <stdio.h>

void main () {
	unsigned char  a;
	unsigned char  b;
	unsigned short c;

	a = 0x12;
	b = 0x34;
	c = 0x0000;

	c |= (unsigned short)(a << 8);
	c |= (unsigned short)(b << 0);

	printf("c is 0x%04X\n", c);
}
