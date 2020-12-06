#include "nucleotide_count.h"
#include <stdlib.h>
#include <stdio.h>

#define MAX_LENGTH 20

char* count(const char* dna_strand)
{
    char* ret = malloc(MAX_LENGTH);
    int a = 0, c = 0, g = 0, t = 0;
    char ch;

    while ((ch = *dna_strand++) != '\0') {
        if (ch == 'A') {
            a++;
        } else if (ch == 'C') {
            c++;
        } else if (ch == 'G') {
            g++;
        } else if (ch == 'T') {
            t++;
        } else {
            ret[0] = '\0';
            return ret;
        }
    }

    sprintf(ret, "A:%d C:%d G:%d T:%d", a, c, g, t);
    return ret;
}
