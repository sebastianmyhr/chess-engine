#ifndef BOARD_H
#define BOARD_H

#include <iostream>

class Board {
    public:
        Board();
        ~Board();
        void print_board();

    private:
        char** board_array;
};
#endif