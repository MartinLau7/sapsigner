//
//  kIOMasterPortDefault.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-08.
//

#ifndef _KIOMASTERPORTDEFAULT_H
#define _KIOMASTERPORTDEFAULT_H

#include <mach/port.h>

#ifdef kIOMasterPortDefault
#undef kIOMasterPortDefault
#endif	/* kIOMasterPortDefault */

const mach_port_t kIOMasterPortDefault;

#endif	/* _KIOMASTERPORTDEFAULT_H */
