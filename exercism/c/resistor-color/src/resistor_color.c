#include "resistor_color.h"

int color_code(resistor_band_t band) {
    return band;
}

const resistor_band_t allColors[] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};

const resistor_band_t* colors() {
    return allColors;
}
