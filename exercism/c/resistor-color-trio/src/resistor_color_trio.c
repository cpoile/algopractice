#include "resistor_color_trio.h"
#include <math.h>

resistor_value_t color_code(resistor_band_t colors[])
{
    resistor_value_t ret;
    ret.value = 10 * colors[0] + colors[1];
    ret.value *= pow(10, colors[2]);
    ret.unit = OHMS;

    if (ret.value >= 1000 && ret.value % 1000 == 0) {
        ret.value = ret.value / 1000;
        ret.unit = KILOOHMS;
    }

    return ret;
}
