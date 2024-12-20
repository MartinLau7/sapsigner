#include "abort.h"

#include "pthread_rwlock_init.h"

int pthread_rwlock_init(pthread_rwlock_t *lock, const pthread_rwlockattr_t *attr) {
    return 0;

    (void) lock;
    (void) attr;
}
