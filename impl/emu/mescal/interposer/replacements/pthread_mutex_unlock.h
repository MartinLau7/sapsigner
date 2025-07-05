//
//  pthread_mutex_unlock.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-07-03.
//

#ifndef _PTHREAD_MUTEX_UNLOCK_H
#define _PTHREAD_MUTEX_UNLOCK_H

#include <pthread.h>

#ifdef pthread_mutex_unlock
#undef pthread_mutex_unlock
#endif	/* pthread_mutex_unlock */

__attribute__ ((noinline)) int pthread_mutex_unlock(pthread_mutex_t *mutex);

#endif	/* _PTHREAD_MUTEX_UNLOCK_H */
