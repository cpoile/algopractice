#include "resistor_color_duo.h"

uint16_t color_code(resistor_band_t colors[])
{
    return 10 * colors[0] + colors[1];
}
