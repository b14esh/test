#include <stdio.h>

int main ()
{
	int age, weight;
        char name[50];
        printf("Inter you name: ");
	//scanf("%s", &name); перед символ аперсант не ставится
	scanf("%s", name);
	printf("Inter you age: ");
	scanf("%d", &age);
        printf("Inter you weight: ");
        scanf("%d", &weight);
        printf("You inter: \n");
	printf("You name %s, you ages %d, you weight %d \n", name, age, weight);
	printf("Address var name for memory pc - %p\n", &name);	
}     	
