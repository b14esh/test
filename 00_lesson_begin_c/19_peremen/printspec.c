#include <stdio.h>

int main()
{
	char letter = 'A';
	double decimal = 7.389;
	int number = 100;
	printf("привет мир\n");
	printf("%c = одиночный символ\n", letter);
	printf("decimal = %g, number = %d\n", decimal, number);
	printf("%05.2g = decimal\n", decimal);
}
