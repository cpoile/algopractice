#include <stdint.h>
#include "gigasecond.h"

#define GIGASECOND 1000000000

time_t gigasecond_after(time_t start)
{
    return (uintmax_t)start + GIGASECOND;
}
