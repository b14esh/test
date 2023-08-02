/* execve01.c */
#include <unistd.h>
#include <stdio.h>

int main (void)
{
        printf ("pid=%d\n", getpid ());
        execve ("/bin/cat", NULL, NULL);

        return 0;
}
