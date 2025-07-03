#include "abort.h"
#include "memcpy.h"

#include "__memcpy_chk.h"

char *__memcpy_chk(void *restrict dst, const void *restrict src, size_t len, size_t dstlen) {
    if (dstlen < len) {
        abort();

        return NULL;
    }

    return memcpy(dst, src, len);
}
