//
//  gettimeofday.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _GETTIMEOFDAY_H
#define _GETTIMEOFDAY_H

#include <sys/time.h>

#ifdef gettimeofday
#undef gettimeofday
#endif	/* gettimeofday */

__attribute__ ((noinline)) int gettimeofday(struct timeval *restrict tp, void *restrict tzp);

#endif	/* _GETTIMEOFDAY_H */
