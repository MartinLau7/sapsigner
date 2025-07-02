#include "memcmp.h"

int memcmp(const void *s1, const void *s2, size_t n) {
    while (n--) {
        int diff = *((const unsigned char *) s1) - *((const unsigned char *) s2);
        if (diff) {
            return diff;
        }

        s1 = (const unsigned char *) s1 + 1;
        s2 = (const unsigned char *) s2 + 1;
    }

    return 0;
}
