export default class Vector {
  constructor(x, y) {
    this.x = x
    this.y = y
  }

  add(o) {
    return new Vector(this.x + o.x, this.y + o.y)
  }

  sub(o) {
    return new Vector(this.x - o.x, this.y - o.y)
  }

  dot(o) {
    return new Vector(this.x * o.x, this.y * o.y)
  }

  scale(scalar) {
    return new Vector(this.x + scalar, this.y + scalar)
  }

  mag() {
    return Number.parseInt(Math.sqrt(v.x ** 2 + v.y ** 2))
  }

  norm() {
    const mag = this.mag();
    if (mag == 0) {
      return new Vector(0, 0)
    }
    return new Vector(this.x / mag, this.y / mag)
  }

  toString() {
    return `Vector(${this.x}, ${this.y})`;
  }
}
