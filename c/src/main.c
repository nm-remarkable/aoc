#include <stdio.h>

int add(int a, int b)
{
    return a + b;
}

int main(void)
{
    int x = 5;
    int y = 16;
    int z = add(x, y);
    printf("%d + %d = %d\n", x, y, z);
    return 0;
}
