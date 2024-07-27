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
        
        console.log(fromX + "," + fromY + " -> " + x + "," + y);
        
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

async function movePiece(fromX, fromY, toX, toY) {
    try {
        const response = await fetch('/move', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ fromX, fromY, toX, toY }),
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Failed to make move');
        }

        boardState = await response.json();
        renderBoard();
    } catch (error) {
        console.error('Error making move:', error.message);
        // alert(error.message); // Display error to the user
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