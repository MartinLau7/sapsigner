#include "abort.h"

void abort(void) {
    __asm__ ("ud2");

    __builtin_unreachable();
}
