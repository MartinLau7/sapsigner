#include "__memset_chk.h"

#include "memset.h"

void *__memset_chk(void *dest, int val, size_t len, size_t dstlen) {
    if (dstlen < len) {
        return NULL;
    }

    return memset(dest, val, len);
}
