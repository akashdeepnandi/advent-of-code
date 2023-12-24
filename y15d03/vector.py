import math

class Vector:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __repr__(self):
        return f"Vector({self.x}, {self.y})"

    def __add__(self, other):
        return Vector(self.x + other.x, self.y + other.y)

    def __sub__(self, other):
        return Vector(self.x - other.x, self.y - other.y)

    def __mul__(self, scalar):
        return Vector(self.x * scalar, self.y * scalar)

    def dot(self, other):
        return self.x * other.x + self.y * other.y

    def mag(self):
        return int(math.sqrt(self.x**2 + self.y**2))

    def norm(self):
        mag = self.mag()
        if mag == 0:
            return Vector(0, 0)
        return Vector(self.x / mag, self.y / mag)
