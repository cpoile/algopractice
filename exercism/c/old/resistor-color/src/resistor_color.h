#ifndef RESISTOR_COLOR_H
#define RESISTOR_COLOR_H

#include <inttypes.h>
#include <stdbool.h>

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
    NUM_BANDS,
} resistor_band_t;  // == signed int

int color_code(int color);

int *colors();

#endif
