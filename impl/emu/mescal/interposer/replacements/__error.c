#include "__error.h"

int *__error(void) {
    static int _errno = 0;
    return &_errno;
}
