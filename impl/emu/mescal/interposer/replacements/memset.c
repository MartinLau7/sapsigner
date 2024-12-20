#include "memset.h"

void *memset(void *b, int c, size_t len) {
    for (unsigned char *it = b; it != (unsigned char *) b + len; ++it) {
        *it = (unsigned char) c;
    }

    return b;
}
