#include "pthread_once.h"

int pthread_once(pthread_once_t *once_control, void (*init_routine)(void)) {
    if (once_control->__sig) {
        once_control->__sig = 0;

        init_routine();
    }

    return 0;
}
