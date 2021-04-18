import random

class Brain:
    def __init__(self):
        self.id = None
        self.board = None
        self.body = None
        self.turn = None
        self.lastTurn = None
        self.allowedMoves = []
        self.possibleMoves = ["up", "down", "left", "right"]

    def biteSelf(self):
        if self.lastmove != None:
            self.allowedMoves.remove[self.lastMove]

    def outOfBoard(self):
        x = self.body["head"]["x"]
        y = self.body["head"]["y"]
        if x + 1 == self.board["width"]:
            self.allowedMoves.remove["right"]
        elif x == 1:
            self.allowedMoves.remove["left"]
        if y + 1 == self.board["height"]:
            self.allowedMoves.remove["up"]
        elif y == 1:
            self.allowedMoves.remove["down"]
    
    def move(self):
        self.allowedMoves = self.possibleMoves
        self.biteSelf()
        self.outOfBoard()
        if len(self.allowedMoves) == 0:
            return random.choice(self.possibleMoves)
        return self.allowedMoves