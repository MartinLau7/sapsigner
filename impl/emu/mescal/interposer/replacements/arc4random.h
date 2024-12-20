//
//  arc4random.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _ARC4RANDOM_H
#define _ARC4RANDOM_H

#include <stdint.h>

#ifdef arc4random
#undef arc4random
#endif	/* arc4random */

__attribute__ ((noinline)) uint32_t arc4random(void);

#endif	/* _ARC4RANDOM_H */
