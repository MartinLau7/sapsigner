#include "IOIteratorNext.h"

io_object_t IOIteratorNext(io_iterator_t iterator) {
    static io_object_t o = 0;
    return --o % 2;

    (void) iterator;
}
