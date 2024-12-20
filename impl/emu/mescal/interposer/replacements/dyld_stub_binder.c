#include "abort.h"

#include "dyld_stub_binder.h"

void dyld_stub_binder(void) {
    abort();
}
