document.addEventListener('DOMContentLoaded', () => {
    const boardElement = document.getElementById('board');
    let selectedSquare = null;

    function createBoard() {
        for (let i = 0; i < 8; i++) {
            for (let j = 0; j < 8; j++) {
                const square = document.createElement('div');
                square.classList.add('square');
                square.classList.add((i + j) % 2 === 0 ? 'white' : 'black');
                square.dataset.x = i;
                square.dataset.y = j;
                square.addEventListener('click', onSquareClick);
                boardElement.appendChild(square);
            }
        }
    }

    function onSquareClick(event) {
        const square = event.currentTarget;
        if (selectedSquare) {
            const fromX = selectedSquare.dataset.x;
            const fromY = selectedSquare.dataset.y;
            const toX = square.dataset.x;
            const toY = square.dataset.y;

            makeMove(fromX, fromY, toX, toY);
            selectedSquare.classList.remove('selected');
            selectedSquare = null;
        } else {
            selectedSquare = square;
            selectedSquare.classList.add('selected');
        }
    }

    function makeMove(fromX, fromY, toX, toY) {
        fetch('/move', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ fromX, fromY, toX, toY }),
        })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                updateBoard();
            } else {
                alert('Invalid move');
            }
        });
    }

    function updateBoard() {
        fetch('/')
        .then(response => response.json())
        .then(data => {
            const board = data.board;
            for (let i = 0; i < 8; i++) {
                for (let j = 0; j < 8; j++) {
                    const square = boardElement.querySelector(`.square[data-x="${i}"][data-y="${j}"]`);
                    square.innerHTML = '';
                    const piece = board.Squares[i][j];
                    if (piece) {
                        const pieceElement = document.createElement('span');
                        pieceElement.classList.add('piece');
                        pieceElement.textContent = getPieceSymbol(piece);
                        square.appendChild(pieceElement);
                    }
                }
            }
        });
    }

    function getPieceSymbol(piece) {
        const symbols = {
            'white-king': '♔',
            'white-queen': '♕',
            'white-rook': '♖',
            'white-bishop': '♗',
            'white-knight': '♘',
            'white-pawn': '♙',
            'black-king': '♚',
            'black-queen': '♛',
            'black-rook': '♜',
            'black-bishop': '♝',
            'black-knight': '♞',
            'black-pawn': '♟',
        };
        return symbols[`${piece.Color}-${piece.Type}`];
    }

    createBoard();
    updateBoard();
});
