//
//  IOIteratorNext.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _IOITERATORNEXT_H
#define _IOITERATORNEXT_H

#include <IOKit/IOKitLib.h>

#ifdef IOIteratorNext
#undef IOIteratorNext
#endif	/* IOIteratorNext */

__attribute__ ((noinline)) io_object_t IOIteratorNext(io_iterator_t iterator);

#endif	/* _IOITERATORNEXT_H */
