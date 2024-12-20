#include "abort.h"

#include "__stack_chk_guard.h"

void __stack_chk_guard(void) {
    abort();
}
