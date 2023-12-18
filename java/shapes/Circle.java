

public class Circle extends Shape {

    private double _radius;

    public Circle(double radius) {
        this._radius = radius;
    }

    public void setRadius(int radius) {
        this._radius = radius;
    }

    public double getRadius() {
        return this._radius;
    }

    @Override
    public double area() {
        return Math.PI * Math.pow(this._radius, 2);
    }
}