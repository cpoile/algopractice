//
// Created by Chris Poile on 2020-11-30.
//

#include "armstrong_numbers.h"
#include <math.h>

int is_armstrong_number(int candidate)
{
    int sum = 0;
    int orig = candidate;
    int len = log10(candidate) + 1;

    while (candidate != 0) {
        int n = candidate % 10;
        sum += pow(n, len);
        candidate /= 10;
    }

    return sum == orig;
}
