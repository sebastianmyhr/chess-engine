#include "board.h"

int main(int arcg, char** argv) {

    Board *b1 = new Board();
    b1->print_board();
    delete b1;

    return 0;
}