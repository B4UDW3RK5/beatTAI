#include <stdio.h>
#include "leapsecs.h"
#include "tai.h"
#include "caltime.h"

// get libtai from:
// https://cr.yp.to/libtai.html
// cc -o beatai beatai.c libtai.a

struct tai sec;
struct caltime ct;

int main(int argc,char *argv[]) {
  if (leapsecs_init() == -1) {
    fprintf(stderr,"fatal: unable to init leapsecs\n");
    return 111;
  }
  tai_now(&sec);
  caltime_utc(&ct,&sec,(int *) 0,(int *) 0);
  unsigned int milibeats=(((((ct.hour) * 3600) + (ct.minute * 60) + ct.second)) * 10000) / 864;
  unsigned int wholebeats=milibeats / 1000;
  unsigned int partialbeats=milibeats % 1000;
  printf("@%u.%02u\n",wholebeats,partialbeats/10);
  return 0;
}