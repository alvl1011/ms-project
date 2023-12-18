# Modelbased Software Development: OO in Java versus Go

## Overview
This project contains examples of nominal subtyping and virtual methods in Java and structural subtyping in Go.

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