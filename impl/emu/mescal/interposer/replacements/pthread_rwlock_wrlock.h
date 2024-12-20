//
//  pthread_rwlock_wrlock.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _PTHREAD_RWLOCK_WRLOCK_H
#define _PTHREAD_RWLOCK_WRLOCK_H

#include <pthread.h>

#ifdef pthread_rwlock_wrlock
#undef pthread_rwlock_wrlock
#endif	/* pthread_rwlock_wrlock */

__attribute__ ((noinline)) int pthread_rwlock_wrlock(pthread_rwlock_t *rwlock);

#endif	/* _PTHREAD_RWLOCK_WRLOCK_H */
