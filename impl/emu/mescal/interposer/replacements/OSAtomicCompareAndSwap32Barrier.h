//
//  OSAtomicCompareAndSwap32Barrier.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-07-03.
//

#ifndef _OSATOMICCOMPAREANDSWAP32BARRIER_H
#define _OSATOMICCOMPAREANDSWAP32BARRIER_H

#include <stdbool.h>
#include <stdint.h>

#ifdef OSAtomicCompareAndSwap32Barrier
#undef OSAtomicCompareAndSwap32Barrier
#endif	/* OSAtomicCompareAndSwap32Barrier */

__attribute__ ((noinline)) bool OSAtomicCompareAndSwap32Barrier(int32_t oldValue, int32_t newValue, volatile int32_t *theValue);

#endif	/* _OSATOMICCOMPAREANDSWAP32BARRIER_H */
