#include <iostream>
#include <ctime>
#include <vector>
#include <memory>
namespace 🔵=std;
using 🔢=int;
using 💀=void;
using 🕖=time_t;
using 👌=bool;
#define 👂 auto
#define 🎌 enum
#define 👎 false
#define 👍 true
#define 💪 🔵::make_shared
#define 🍸 virtual
#define 🖥 🔵::cout
#define 🔫 🔵::endl
template<class 🔮>
using 📚 = 🔵::vector<🔮>;
template<class 🔮>
using 👇 =  🔵::shared_ptr<🔮>;

🎌 🐒{🐵,🙈,🙉,🙊};
🔢 🎲(){return 🔵::rand();}
👌 😎(){return 👎;}

struct 🍴{🍸 💀 👀()=0;};
struct 🍊:🍴{🍸 💀 👀(){🖥<<"🍊"<<🔫;};};
struct 🍉:🍴{🍸 💀 👀(){🖥<<"🍉"<<🔫;};};
struct 🍒:🍴{🍸 💀 👀(){🖥<<"🍒"<<🔫;};};
struct 🍓:🍴{🍸 💀 👀(){🖥<<"🍓"<<🔫;};};
struct 🍍:🍴{🍸 💀 👀(){🖥<<"🍍"<<🔫;};};
struct 🍎:🍴{🍸 💀 👀(){🖥<<"🍎"<<🔫;};};

🔢 main() 
{
    if(😎()==👎)
        🖥<<"💩"<<🔫;
    
    📚<👇<🍴>> 🍛={💪<🍊>(),💪<🍉>(),💪<🍒>(),💪<🍓>(),💪<🍍>(),💪<🍎>()};

    for (👂 &🍏: 🍛)
        🍏->👀();
    
    return 🎲();
}