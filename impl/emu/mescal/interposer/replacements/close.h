//
//  close.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-06-28.
//

#ifndef _CLOSE_H
#define _CLOSE_H

#ifdef close
#undef close
#endif	/* close */

__attribute__ ((noinline)) int close(int fildes);

#endif	/* _CLOSE_H */
