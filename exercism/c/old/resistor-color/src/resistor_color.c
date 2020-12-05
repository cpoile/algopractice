#include <stddef.h>
#include "resistor_color.h"

static int bands[NUM_BANDS];
static bool bands_init = false;

int color_code(int color)
{
    return color;
}

int *colors()
{
    if (!bands_init) {
        for (size_t i = 0; i < NUM_BANDS; i++) {
            bands[i] = i;
        }
        bands_init = true;
    }

    return bands;
}
