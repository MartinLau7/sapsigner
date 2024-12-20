//
//  pthread_self.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _PTHREAD_SELF_H
#define _PTHREAD_SELF_H

#include <pthread.h>

#ifdef pthread_self
#undef pthread_self
#endif	/* pthread_self */

__attribute__ ((noinline)) pthread_t pthread_self(void);

#endif	/* _PTHREAD_SELF_H */
