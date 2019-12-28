#include <math.h>
#include "darts.h"

uint8_t score(coordinate_t hit)
{
    double hyp = hypot(hit.x, hit.y);
    if (hyp > 10.0)
    {
        return 0;
    } else if (hyp > 5.0)
    {
        return 1;
    } else if (hyp > 1.0)
    {
        return 5;
    } else
    {
        return 10;
    }
}
