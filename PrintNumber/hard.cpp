#include <iostream>

auto main() -> int
{
    for (std::string s{"1 2 3 4 5 "}; s.size() > 0;
         s = s.substr(0, s.size() - 2))
    {
        std::cout << s << std::endl;
    }
}