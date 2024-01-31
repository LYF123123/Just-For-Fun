#include <iostream> 

class console{
    public:
    void log(const char* logmsg){
        std::cout << logmsg << std::endl;
    }
}console;

class System{ 
    public:
        struct 
        {
            void println(const char* str){
                std::cout << str << std::endl;
            }
        }out;
        
}System;

int main(){ 
    console.log("Hello, world!");
    System.out.println("Hello, world!");
    return 0;
}