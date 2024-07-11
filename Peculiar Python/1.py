class Coords:
    def __init__(self, x:int, y:int) -> None:
        self.x = x
        self.y = y
    
    def __lt__(self, other):
        if not isinstance(other, Coords):
            return NotImplemented
        return self.x < other.x and self.y < other.y
        

a = Coords(1,2)
b = Coords(3,4)

print(a<b)
print(b<a)
print(a<2)