#include <iostream>
#include <functional>
auto main() -> int
{
    std::function<std::string(int)> f;
    f = [&f](int n)
    {
        return (n == 1) ? "1" : f(n - 1) + " " + std::to_string(n);
    };
    auto fun = [&f]<typename... Ts>(Ts... n)
    {
        return ((f(n) + "\n") + ...);
    };
    std::cout << fun(5, 4, 3, 2, 1);
}
// g++ .\epic.cpp -o epic.exe -std=c++17