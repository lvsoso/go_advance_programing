#include <stdio.h>
#include <string.h>
#include "hello.h"

int main(int argc, char *argv[]){
    char *type = argv[1];
    char buf[] = "test for";
    char input[16] = {0};

    sprintf(input, "%s %s", buf, type);

    GoString str;
    str.p = input;
    str.n = strlen(input);

    Hello(str);
    return 0;
}