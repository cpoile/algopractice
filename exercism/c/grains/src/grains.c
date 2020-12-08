#include "grains.h"

uint64_t square(uint8_t index)
{
    if (index < 1 || index > 64) {
        return 0;
    }

    return (uint64_t) 1 << (index - 1);
}

uint64_t total(void)
{
    uint64_t t = 0;
    for (int i = 1; i <= 64; i++) {
        t += square(i);
    }

    return t;
}
