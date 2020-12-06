#include "space_age.h"

const float planetEarthYears[] = {
        0.2408467,
        0.61519726,
        1,
        1.8808158,
        11.862615,
        29.447498,
        84.016846,
        164.79132,
};

float convert_planet_age(planet_t planet, int64_t input)
{
    int64_t sPerEYear = 31557600;
    return input / (planetEarthYears[planet] * sPerEYear);
}
