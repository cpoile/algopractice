#include "resistor_color_trio.h"
#include <math.h>

resistor_value_t color_code(resistor_band_t bands[]) {
    resistor_value_t ret = {
        .value =  10 * bands[0] + bands[1],
        .unit = OHMS,
    };

    if (bands[2] > 0) {
        ret.value = ret.value * pow(10, bands[2]);
    }

    if (ret.value >= 1000) {
        ret.value = ret.value / 1000;
        ret.unit = KILOOHMS;
    }

    return ret;
}
