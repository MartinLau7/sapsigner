#include "__memcpy_chk.h"

#include "memcpy.h"

char *__memcpy_chk(void *restrict dst, const void *restrict src, size_t len, size_t dstlen) {
    if (dstlen < len) {
        return NULL;
    }

    return memcpy(dst, src, len);
}
