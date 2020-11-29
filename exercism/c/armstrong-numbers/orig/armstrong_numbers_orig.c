#include "armstrong_numbers_orig.h"
#include <math.h>

int is_armstrong_number(int candidate)
{
    int len = ceil(log10(1 + candidate));

    int sum = 0;
    int cur = candidate;
    while (cur > 0)
    {
        sum += pow(cur % 10, len);
        cur /= 10;
    }

    return sum == candidate;
}


