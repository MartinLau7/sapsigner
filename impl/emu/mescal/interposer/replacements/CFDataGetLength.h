//
//  CFDataGetLength.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _CFDATAGETLENGTH_H
#define _CFDATAGETLENGTH_H

#include <CoreFoundation/CoreFoundation.h>

#ifdef CFDataGetLength
#undef CFDataGetLength
#endif	/* CFDataGetLength */

__attribute__ ((noinline)) CFIndex CFDataGetLength(CFDataRef theData);

#endif	/* _CFDATAGETLENGTH_H */
