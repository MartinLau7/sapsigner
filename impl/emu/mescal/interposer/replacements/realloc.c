#include "malloc.h"
#include "malloc_good_size.h"
#include "malloc_size.h"
#include "memcpy.h"

#include "realloc.h"

void *realloc(void *old_ptr, size_t new_size) {
    if (!old_ptr) {
        return malloc(new_size);
    }

    size_t old_size = malloc_size(old_ptr);
    if (new_size <= old_size) {
        if (new_size < old_size) {
            size_t aligned_size_t_size = malloc_good_size(sizeof(size_t));
            *((size_t *) ((unsigned char *) old_ptr - aligned_size_t_size)) = new_size;
        }

        return old_ptr;
    }

    void *new_ptr = malloc(new_size);
    return memcpy(new_ptr, old_ptr, old_size);
}
