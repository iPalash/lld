
from enum import Enum
from random import randint
class Player(Enum):
    ONE = 1
    TWO = 2

class Board:


    def __init__(self,n,m,k) -> None:
        self.n = n
        self.m = m
        self.k = k
        self.board = [[0 for _ in range(m)] for _ in range(n)]
        self.horizontal = [[0 for _ in range(m)] for _ in range(n)]
        self.vertical = [[0 for _ in range(m)] for _ in range(n)]
        self.diagonal = [[0 for _ in range(m)] for _ in range(n)]

        self.top = [0 for _ in range(m)]


    def neighbors(self, i,j, direction):
        curr = self.board[i][j]
        if direction == 'h':
            modifiers = [[0,-1],[0,+1]]
        elif direction == 'v':
            modifiers = [[-1,0],[+1,0]]
        elif direction == 'd1':
            modifiers = [[-1,-1],[+1,+1]]
        elif direction == 'd2':
            modifiers = [[-1,+1],[+1,-1]]
        for [x,y] in modifiers:
            a,b=i+x,j+y
            if self.valid(a,b) and self.board[a][b]==curr:
                yield [a,b]
    def valid(self,a,b):
        return 0<=a<self.n and 0<=b<self.m
    
    def allUpto(self, i,j, direction):
        curr = self.board[i][j]
        if direction == 'l':
            [a,b] = [0,-1]
        elif direction == 'r':
            [a,b] = [0,+1]
        elif direction == 'u':
            [a,b] = [1,0]
        elif direction == 'd':
            [a,b] = [-1,0]
        elif direction == 'lu':
            [a,b] = [1,-1]
        elif direction == 'ld':
            [a,b] = [-1,-1]
        elif direction == 'ru':
            [a,b] = [1,1]
        elif direction == 'rd':
            [a,b] = [-1,1]
        x,y = i+a,j+b
        # print(direction, "of",i,j,(x,y))
        while self.valid(x,y) and self.board[x][y]==curr:
            yield [x,y]
            [x,y] = [x+a,y+b]
        



    def _update(self, row, col):
        curr = self.board[row][col]
        print("marked", row,col)
        # horizontal
        sm = 0
        for [a,b] in self.neighbors(row,col,'h'):
            sm+=self.horizontal[a][b]
        # print('updating with val',sm+1)
        self.horizontal[row][col]=sm+1
        # Update horizontal to left and right
        for [a,b] in self.allUpto(row,col,'l'):
            # print('updating with val',sm+1, "at", (a,b))    
            self.horizontal[a][b]=sm+1
        
        for [a,b] in self.allUpto(row,col,'r'):
            self.horizontal[a][b]=sm+1
        #vertical
        sm = 0
        for [a,b] in self.neighbors(row,col,'v'):
            if self.board[a][b]==curr:
                sm+=self.vertical[a][b]
        self.vertical[row][col]=sm+1
        # Update vertical
        for [a,b] in self.allUpto(row,col,'u'):
            self.vertical[a][b]=sm+1
        
        for [a,b] in self.allUpto(row,col,'d'):
            self.vertical[a][b]=sm+1

        #diagonal /
        sm = 0
        for [a,b] in self.neighbors(row,col,'d1'):
            if self.board[a][b]==curr:
                sm+=self.diagonal[a][b]
        self.diagonal[row][col]=sm+1 # 
        # Update d1
        for [a,b] in self.allUpto(row,col,'ru'):
            self.diagonal[a][b]=sm+1
        
        for [a,b] in self.allUpto(row,col,'ld'):
            self.diagonal[a][b]=sm+1


        #Diagonal \
        sm = 0
        for [a,b] in self.neighbors(row,col,'d2'):
            if self.board[a][b]==curr:
                sm+=self.diagonal[a][b]
        self.diagonal[row][col]=max(self.diagonal[row][col],sm+1)
        # Update vertical
        for [a,b] in self.allUpto(row,col,'lu'):
            self.diagonal[a][b]=sm+1
        
        for [a,b] in self.allUpto(row,col,'rd'):
            self.diagonal[a][b]=sm+1

        

        return max(self.horizontal[row][col],self.vertical[row][col],self.diagonal[row][col])==self.k
        


    def drop(self, col, p:Player) -> bool:
        row = self.top[col]
        if row==self.n:
            
            raise Exception("Invalid Move:Col Full")
        self.top[col]+=1
        self.board[row][col]=p.value
        return self._update(row,col)
    def print(self):
        print("Board")
        for b in self.board[::-1]:
            print(b)


    def score(p: Player) -> int:
        return 0

class Game:
    def __init__(self) -> None:
        self.board = Board(7,7,4)
        self.turn = 0
        self.players = [Player.ONE,Player.TWO]
        
    def getActivePlayer(self):
        return self.players[self.turn%2]
    
    def start(self):
        while 1:
            
            active = self.getActivePlayer()
            col = randint(0,6)
            print(self.turn, active, col)
            won = self.board.drop(col,active)

            self.board.print()
            if won:
                print("H:V:D")
                for h in self.board.horizontal[::-1]:
                    print(h)
                print()
                for v in self.board.vertical[::-1]:
                    print(v)
                print()
                for d in self.board.diagonal[::-1]:
                    print(d)
                print()
                print(active,"wins")
                break
            self.turn+=1
            print()
            # input()
            

if __name__=='__main__':
    Game().start()


