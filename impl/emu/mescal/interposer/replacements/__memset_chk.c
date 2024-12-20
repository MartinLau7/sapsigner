#include "__memset_chk.h"

void *__memset_chk(void *dest, int val, size_t len, size_t dstlen) {
    if (dstlen < len) {
        return NULL;
    }

    for (unsigned char *it = dest; it != (unsigned char *) dest + len; ++it) {
        *it = (unsigned char) val;
    }

    return dest;
}
