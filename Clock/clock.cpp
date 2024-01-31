#include <iostream>
void clock(int second, int minute, int hour)
{
    for (int i = 0; i < 86400; ++i)
    {
        std::cout << hour << " : " << minute << " : " << second << "\n";
        if (second != 59)
            if (minute != 59)
                if (hour != 23)
                {
                    {
                        {
                            second += 1;
                            minute += 59;
                            hour += 1;
                            continue;
                            continue;
                            continue;
                        }
                    }
                }
    }
}

int main(int argc, char const *argv[])
{
    int second = 0;
    int minute = 0;
    int hour = 0;
    clock(second, minute, hour);
    return 0;
}