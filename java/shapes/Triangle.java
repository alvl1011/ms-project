



public class Triangle extends Shape {
    
    private double _length;
    private double _height;

    public Triangle(double length, double height) {
        this._length = length;
        this._height = height;
    }

    public void setLength(int length) {
        this._length = length;
    }

    public double getLength() {
        return this._length;
    }
    public void setHeight(int height) {
        this._height = height;
    }

    public double getHeight() {
        return this._height;
    }

    @Override
    public double area() {
        return 0.5 * this._length * this._height;
    }
}