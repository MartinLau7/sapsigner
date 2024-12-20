#include "abort.h"
#include "malloc_good_size.h"

#include "malloc.h"

void *malloc(size_t size) {
    static unsigned char *volatile heap_head = NULL;
    static unsigned char *volatile heap_tail = NULL;
    static volatile size_t heap_size = 0;

    size_t aligned_size_t_size = malloc_good_size(sizeof(size_t));
    size_t aligned_size = malloc_good_size(size);

    if (heap_tail + aligned_size_t_size + aligned_size > heap_head + heap_size) {
        abort();

        return NULL;
    }

    *((size_t *) heap_tail) = size;

    void *ptr = heap_tail + aligned_size_t_size;

    heap_tail = (unsigned char *) ptr + aligned_size;

    return ptr;
}
