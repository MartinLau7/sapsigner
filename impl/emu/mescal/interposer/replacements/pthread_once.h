//
//  pthread_once.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _PTHREAD_ONCE_H
#define _PTHREAD_ONCE_H

#include <pthread.h>

#ifdef pthread_once
#undef pthread_once
#endif	/* pthread_once */

__attribute__ ((noinline)) int pthread_once(pthread_once_t *once_control, void (*init_routine)(void));

#endif	/* _PTHREAD_ONCE_H */
