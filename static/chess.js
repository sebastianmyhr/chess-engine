const pieceImages = {
    1: { white: 'pw.svg', black: 'pb.svg' },
    2: { white: 'nw.svg', black: 'nb.svg' },
    3: { white: 'bw.svg', black: 'bb.svg' },
    4: { white: 'rw.svg', black: 'rb.svg' },
    5: { white: 'qw.svg', black: 'qb.svg' },
    6: { white: 'kw.svg', black: 'kb.svg' }
};

const boardElement = document.getElementById('board');
let boardState = [];
let selectedPiece = null;

function createBoard() {
    for (let i = 0; i < 8; i++) {
        for (let j = 0; j < 8; j++) {
            const cell = document.createElement('div');
            cell.classList.add('cell', (i + j) % 2 === 0 ? 'white' : 'black');
            cell.setAttribute('data-x', j);
            cell.setAttribute('data-y', i);
            cell.addEventListener('click', onCellClick);
            boardElement.appendChild(cell);
        }
    }
}

function onCellClick(event) {
    const cell = event.target.closest('.cell');
    if (!cell) return;

    const x = parseInt(cell.getAttribute('data-x'), 10);
    const y = parseInt(cell.getAttribute('data-y'), 10);

    if (isNaN(x) || isNaN(y)) {
        console.error('Invalid cell coordinates:', x, y);
        return;
    }

    if (!selectedPiece) {
        // Select a piece
        if (boardState[y] && boardState[y][x]) {
            selectedPiece = { x, y };
            cell.classList.add('selected');
            
        }
    } else {
        // Move the selected piece
        const fromX = selectedPiece.x;
        const fromY = selectedPiece.y;
        
        if (fromX !== x || fromY !== y) {
            // Only move if the destination is different from the start
            movePiece(fromX, fromY, x, y);
        }

        // Clear selection
        boardElement.querySelector('.selected')?.classList.remove('selected');
        selectedPiece = null;
    }
}

function renderBoard() {
    boardElement.querySelectorAll('.cell').forEach(cell => {
        const x = parseInt(cell.getAttribute('data-x'), 10);
        const y = parseInt(cell.getAttribute('data-y'), 10);
        const piece = boardState[y] && boardState[y][x];
        cell.innerHTML = '';
        if (piece) {
            const img = document.createElement('img');
            const pieceImage = pieceImages[piece.type];
            if (pieceImage) {
                img.src = `/static/pieces/${pieceImage[piece.color === 0 ? 'white' : 'black']}`;
                img.classList.add('piece');
                cell.appendChild(img);
            } else {
                console.error('Invalid piece type:', piece.type);
            }
        }
    });
}

function movePiece(fromX, fromY, toX, toY) {
    // Basic move logic (no validation)
    if (boardState[fromY] && boardState[fromY][fromX]) {
        const piece = boardState[fromY][fromX];
        if (!boardState[toY]) boardState[toY] = [];
        boardState[toY][toX] = piece;
        boardState[fromY][fromX] = null;
        renderBoard();
    }
}

async function fetchBoard() {
    try {
        const response = await fetch('/board');
        if (!response.ok) {
            throw new Error('Failed to fetch board: ' + response.statusText);
        }
        boardState = await response.json();
        renderBoard();
    } catch (error) {
        console.error('Error fetching board:', error);
    }
}

createBoard();
fetchBoard();