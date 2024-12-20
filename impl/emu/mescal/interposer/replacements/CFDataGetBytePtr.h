//
//  CFDataGetBytePtr.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _CFDATAGETBYTEPTR_H
#define _CFDATAGETBYTEPTR_H

#include <CoreFoundation/CoreFoundation.h>

#ifdef CFDataGetBytePtr
#undef CFDataGetBytePtr
#endif	/* CFDataGetBytePtr */

__attribute__ ((noinline)) const UInt8 *CFDataGetBytePtr(CFDataRef theData);

#endif	/* _CFDATAGETBYTEPTR_H */
