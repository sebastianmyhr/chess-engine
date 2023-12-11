#include "board.h"

Board::Board() {
    board_array = new char*[8];
    for (int i = 0; i < 8; ++i) {
        board_array[i] = new char[8];
        for (int j = 0; j < 8; ++j) {
            board_array[i][j] = ' ';
        }
    }
}

Board::~Board() {
    for (int i = 0; i < 8; ++i) {
        delete [] board_array[i];
    }
    delete [] board_array;
}

void Board::print_board() {
    for (int i = 0; i < 8; ++i) {
        for (int j = 0; j < 8; ++j) {
            std::cout << "[" << board_array[i][j] << "]";
        }
        std::cout << std::endl;
    }
}