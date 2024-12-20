//
//  pthread_rwlock_unlock.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _PTHREAD_RWLOCK_UNLOCK_H
#define _PTHREAD_RWLOCK_UNLOCK_H

#include <pthread.h>

#ifdef pthread_rwlock_unlock
#undef pthread_rwlock_unlock
#endif	/* pthread_rwlock_unlock */

__attribute__ ((noinline)) int pthread_rwlock_unlock(pthread_rwlock_t *rwlock);

#endif	/* _PTHREAD_RWLOCK_UNLOCK_H */
