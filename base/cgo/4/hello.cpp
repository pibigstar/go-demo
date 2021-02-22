#include <iostream>

extern "C" {
    #include "hello.h"
}
// 使用c++实现SayHello
int SayHello() {
    std::cout<<"Hello World";
    return 0;
}