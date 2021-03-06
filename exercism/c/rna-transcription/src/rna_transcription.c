#include "rna_transcription.h"
#include <stddef.h>
#include <string.h>
#include <stdlib.h>

char* to_rna(const char* dna)
{
    if (dna == NULL) {
        return NULL;
    }

    int len = strlen(dna);
    char* rna = malloc(len + 1);

    for (int i = 0; i < len; i++) {
        switch (dna[i]) {
            case 'G':
                rna[i] = 'C';
                break;
            case 'C':
                rna[i] = 'G';
                break;
            case 'T':
                rna[i] = 'A';
                break;
            case 'A':
                rna[i] = 'U';
                break;
            default:
                free(rna);
                return NULL;
        }
    }

    rna[strlen(dna)] = '\0';
    return rna;
}
