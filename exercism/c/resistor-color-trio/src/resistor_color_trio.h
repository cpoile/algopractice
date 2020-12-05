#ifndef RESISTOR_COLOR_TRIO_H
#define RESISTOR_COLOR_TRIO_H

#include <stdint.h>

typedef enum {
    BLACK,
    BROWN,
    RED,
    ORANGE,
    YELLOW,
    GREEN,
    BLUE,
    VIOLET,
    GREY,
    WHITE,
} resistor_band_t;

typedef enum {
    OHMS,
    KILOOHMS,
} ohm_t;

typedef struct {
    int value;
    ohm_t unit;
} resistor_value_t;

resistor_value_t color_code(resistor_band_t[]);

#endif
