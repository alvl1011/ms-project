# Modelbased Software Development: OO in Java versus Go

authors: Vladimir Alyoshin, Nico Fuchs, Yannic Schmudlach

## Overview
This project contains examples of nominal subtyping and virtual methods in Java and structural subtyping in Go. 
Two examples:
- ### Animals
- ### Shapes

## Basics

Nominal subtyping and structural subtyping are two different approaches to type systems in programming languages. Let's discuss each along with examples in Go (which uses structural subtyping) and Java (which uses nominal subtyping).

1. **Nominal Subtyping:**
   - **Definition:** In nominal subtyping, types are compatible if and only if their names match. It's based on the name of the type.
   - **Example (Java):**
     ```java
     interface Animal {
         void makeSound();
     }

     class Dog implements Animal {
         @Override
         public void makeSound() {
             System.out.println("Woof!");
         }
     }

     class Cat implements Animal {
         @Override
         public void makeSound() {
             System.out.println("Meow!");
         }
     }

     void animalSound(Animal animal) {
         animal.makeSound();
     }

     // Usage
     Animal myPet = new Dog();
     animalSound(myPet); // Outputs: Woof!
     ```

2. **Structural Subtyping:**
   - **Definition:** In structural subtyping, types are compatible if their structures match, i.e., if they have the same set of methods or fields.
   - **Example (Go):**
     ```go
     package main

     import "fmt"

     type Animal interface {
         makeSound()
     }

     type Dog struct{}

     func (d Dog) makeSound() {
         fmt.Println("Woof!")
     }

     type Cat struct{}

     func (c Cat) makeSound() {
         fmt.Println("Meow!")
     }

     func animalSound(animal Animal) {
         animal.makeSound()
     }

     func main() {
         var myPet Animal
         myPet = Dog{}
         animalSound(myPet) // Outputs: Woof!

         myPet = Cat{}
         animalSound(myPet) // Outputs: Meow!
     }
     ```

In the Go example, the `Dog` and `Cat` types implicitly satisfy the `Animal` interface because they both have a `makeSound` method. This is an example of structural subtyping. In Java, the `Dog` and `Cat` classes explicitly implement the `Animal` interface, which is an example of nominal subtyping. 

## Example: Animals

- ***Java***:
    In this example, there is a base class `Animal` with a virtual method `scream()`. The two subclasses `Dog` and `Cat` inherit from the base class and overwrite the method. When creating objects with the types of the base class, you can then call the virtual method and the method of the respective subclass will be executed. This is an example of polymorphic behavior in Java, which is made possible by nominal type checking and virtual methods.

    [Animals in Java](https://github.com/alvl1011/ms-project/tree/develop/java/animals)

- ***Go***:
    In this Go example, structural subtyping is represented by the `Animal` interface. The `Dog` and `Cat` structures both implement the interface by providing the MakeLoud method. When creating objects with the types of the interface, you can then call the method via the interface and the corresponding implementation of the method will be executed for the respective structure. This shows the structural subtyping in Go.

    [Animals in Go](https://github.com/alvl1011/ms-project/tree/develop/go/animals)

## Example: Shapes

A larger example could be a simple system for managing shapes. We create an abstract class Shape, from which concrete classes such as `Circle`, `Rectangle` and `Triangle` inherit. Each class can have specific properties and implement the `area` method. We also use polymorphism and a list of shapes in the main class to manage different shapes and calculate their areas.

- ***Java***:
    In this example, there is a hierarchy of shapes that inherit from the abstract class `Shape`. Each concrete shape implements the `area()` method in its own way. The main class creates different shapes, adds them to a list and then calculates the area of each shape by using polymorphism and virtual methods.

    [Shapes in Java](https://github.com/alvl1011/ms-project/tree/develop/java/shapes)

- ***Go***:
    In Go there are no classes like in object-oriented languages, but you can use interfaces to achieve polymorphism.
    In this example, interfaces are used to achieve polymorphism. The `Circle`, `Rectangle` and `Triangle` structures implement the `Shape` interface by providing the `area()` method. The main function creates different shapes, adds them to a list and then iterates over the list to calculate the area of each shape.

    [Shapes in Go](https://github.com/alvl1011/ms-project/tree/develop/go/shapes)

## Summary

In summary, the key difference is that **nominal subtyping** relies on the explicit naming or declaration of types, while structural subtyping is based on the structure of types. Languages like Java typically use nominal typing, while languages like TypeScript or languages with duck typing often use structural typing. Each approach has its advantages and trade-offs, and the choice between them often depends on the design goals and philosophy of a particular programming language.
