//
//  pthread_mutex_lock.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-07-03.
//

#ifndef _PTHREAD_MUTEX_LOCK_H
#define _PTHREAD_MUTEX_LOCK_H

#include <pthread.h>

#ifdef pthread_mutex_lock
#undef pthread_mutex_lock
#endif	/* pthread_mutex_lock */

__attribute__ ((noinline)) int pthread_mutex_lock(pthread_mutex_t *mutex);

#endif	/* _PTHREAD_MUTEX_LOCK_H */
