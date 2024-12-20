//
//  pthread_rwlock_init.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-15.
//

#ifndef _PTHREAD_RWLOCK_INIT_H
#define _PTHREAD_RWLOCK_INIT_H

#include <pthread.h>

#ifdef pthread_rwlock_init
#undef pthread_rwlock_init
#endif	/* pthread_rwlock_init */

__attribute__ ((noinline)) int pthread_rwlock_init(pthread_rwlock_t *lock, const pthread_rwlockattr_t *attr);

#endif	/* _PTHREAD_RWLOCK_INIT_H */
