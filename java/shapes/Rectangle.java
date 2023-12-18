

public class Rectangle extends Shape {
    
    private double _length;
    private double _width;

    public Rectangle(double length, double width) {
        this._length = length;
        this._width = width;
    }

    public void setLength(int length) {
        this._length = length;
    }

    public double getLength() {
        return this._length;
    }
    public void setWidth(int width) {
        this._width = width;
    }

    public double getWidth() {
        return this._width;
    }

    @Override
    public double area() {
        return this._length * this._width;
    }
}