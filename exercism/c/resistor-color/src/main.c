#include <stdio.h>
#include "resistor_color.h"

int main2()
{
    int *bands = colors();

    for (int i = 0; i < NUM_BANDS; i++) {
        printf("band[%d] = %d\n", i, bands[i]);
    }

    return 0;
}
