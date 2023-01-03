#include <iostream>

auto main()->int 
{
    for(int i=5;i>0;std::cout<<std::endl,i--)
        for(int j=1;j<=i;std::cout<<j<<" ",++j){}
}