//
//  abort.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-01.
//

#ifndef _ABORT_H
#define _ABORT_H

#ifdef abort
#undef abort
#endif	/* abort */

__attribute__ ((noinline)) void abort(void);

#endif	/* _ABORT_H */
