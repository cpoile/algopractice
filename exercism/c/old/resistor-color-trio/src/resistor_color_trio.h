#ifndef RESISTOR_COLOR_TRIO_H
#define RESISTOR_COLOR_TRIO_H

#include <inttypes.h>

typedef enum {
    BLACK,    // 0
    BROWN,    // 1
    RED,      // 2
    ORANGE,   // 3
    YELLOW,   // 4
    GREEN,    // 5
    BLUE,     // 6
    VIOLET,   // 7
    GREY,     // 8
    WHITE,    // 9
    NUM_BANDS,
} resistor_band_t;  // == signed int

enum {
    OHMS,
    KILOOHMS,
};

typedef struct {
    uint32_t value;
    int unit;
} resistor_value_t;

resistor_value_t color_code(resistor_band_t[]);

#endif
