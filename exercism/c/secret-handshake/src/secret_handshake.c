#include "secret_handshake.h"
#include <stdlib.h>

const int MAX_ARR_LEN = 4;
char* actions[] = {"wink", "double blink", "close your eyes", "jump"};

const char** commands(size_t number)
{
    char** ret = malloc(sizeof(char*) * MAX_ARR_LEN);
    ret[0] = NULL;

    int p = 0;
    for (int i = 0; i < 4; ++i) {
        int idx = (number & 1 << 4) ? 4-1-i : i;
        if (number & 1 << idx) {
            ret[p++] = actions[idx];
        }
    }
    return (const char**) ret;
}
