#include <iostream>

auto main() -> int {
    for (int i = 5; i >0; i--)
    {
        for(int j=1;j<=i;++j)
        {
            std::cout <<j<<" ";
        }
        std::cout <<std::endl;
    }
    
}