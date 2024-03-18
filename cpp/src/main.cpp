#include <iostream>

namespace advent
{
    static int add(int a, int b) { return a + b; }

    class Printer
    {
    public:
        Printer(const int id) : m_id(id) {}
        void print(const std::string &message)
        {
            std::cout << message
                      << "\n"
                      << "Print id: "
                      << m_id
                      << std::endl;
        }

    private:
        int m_id;
    };
}

int main(void)
{
    auto p = advent::Printer{advent::add(1, 9)};
    p.print("Hello, World!");
    return 0;
}
