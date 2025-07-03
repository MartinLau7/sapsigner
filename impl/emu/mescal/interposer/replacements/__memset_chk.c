#include "abort.h"
#include "memset.h"

#include "__memset_chk.h"

void *__memset_chk(void *dest, int val, size_t len, size_t dstlen) {
    if (dstlen < len) {
        abort();

        return NULL;
    }

    return memset(dest, val, len);
}
