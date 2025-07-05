//
//  CFBundleGetMainBundle.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-07-04.
//

#ifndef _CFBUNDLEGETMAINBUNDLE_H
#define _CFBUNDLEGETMAINBUNDLE_H

#include <CoreFoundation/CoreFoundation.h>

#ifdef CFBundleGetMainBundle
#undef CFBundleGetMainBundle
#endif	/* CFBundleGetMainBundle */

__attribute__ ((noinline)) CFBundleRef CFBundleGetMainBundle();

#endif	/* _CFBUNDLEGETMAINBUNDLE_H */
