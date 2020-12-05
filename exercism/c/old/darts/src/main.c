#include <stdio.h>
#include "darts.h"

int main2()
{
    coordinate_t attempt = {0.5, 0.5};
    int sc = score(attempt);
    printf("score: %d\n", sc);
    return 0;
}
