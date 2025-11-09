## Topics:

- Object-Oriented Design (OOD)
    - UML Basics
        ○ Introduction to Unified Modeling Language (UML)
        ○ Types of UML Diagrams
    - Important Diagrams for LLD
        ○ Use Case Diagram (borderline HLD, but often used in both)
        ○ Class Diagram
        ○ Sequence Diagram
        ○ Activity Diagram
        ○ State Diagram (optional but helpful)
        ○ Component Diagram (optional, bridges HLD & LLD)
- OOPs Concepts:
    - Encapsulation
    - Inheritance
    - Polymorphism
    - Abstraction
- SOLID Principles
    - Single Responsibility
    - Open/Closed
    - Liskov Substitution
    - Interface Segregation
    - Dependency Inversion
- Design Patterns
    - Creational (Factory, Singleton)
    - Structural (Adapter, Composite)
    - Behavioral (Observer, Strategy, Command)
- Code-Level Design
    - Interfaces, Modules, Function-level decisions
        • Error handling, logging, data structures
    

## OOPS:

OOPS is based on 4 main principles:

    Encapsulation
    Abstraction
    Polymorphism
    Inheritance
        
    Encapsulation:

        → Hiding implementation details of one object from another object.
        → It is a process which encourages decoupling.
        → Encapsulation:
    The Circle and Rectangle structs encapsulate their fields (radius, length, breadth) and provide methods to operate on them.
        
    Abstraction:
    
        → hiding the complex implementation details of an object and exposing only the relevant information to the user.
        → data are visible only to semantically related functions, so as to prevent misuse
        → it can be achieved using interfaces
        → Abstraction:
    The Shape interface abstracts the concept of a shape that can calculate its area through the Area method.
        
    Polymorphism:

        → perform a single action in different ways
        → objects of different types to be treated as if they were the same type
        → Polymorphism makes it easier to write reusable code that can work with different types of objects, without needing to know the specific details of each object.
        → in other words polymorphism allows you to define one interface and have multiple implementations
        → Polymorphism:
    The Area method is implemented by different structs (Circle, Rectangle), and they can all be treated as Shape types due to the interface.
        
    Inheritance:

        → creating a new class that is a modified version of an existing class.
        → The new class, called the child class or subclass, inherits the properties and methods of the parent class or superclass.
        → Can be implemented using embeded struct
        → Inheritance (Composition):
    The Square struct embeds the Rectangle struct, allowing it to inherit its fields and methods. This demonstrates composition in Go, which is used in place of traditional inheritance.
    
Encapsulation: hiding the implementation details of obj from outside world

```go
type Person struct{
 Name string
 Age int
}

func NewPerson(Name string,Age int)*Person{
    return &Person{name,age}
}

func (p *Person)getName()string{
    return p.name
}
func (p *Person)setName()string{
    
}
func (p *Person)getAge()int{
    return p.age
}
func (p *Person)setAge()int{}
```


Polymorphism : objects of different types to be treated as if they are of the same type

```go
type Shape interface{
    Area() float
}

type Square struct{
    length float
}

type Circle struct{
    radius float
}


func (s Square)Area()float{
    return square formula //a2
}
func (s Circle)Area()float{
    return circle formula //3.14*r2
}


func main(){

    var s Shape
    s = Square(5.0)
    
    fmt.Print("Square",s.Area())
    
    s = Circle(3.0)
    fmt.Print("Circle",s.Area())
}
```

Inheritence : inheriting Parent struct properties and methods.

```go
type Animal struct{}

type Dog struct{}


func (d Dog)Speak(){
"Dog barks"
}
func (a Animal)Speak(){
"Animal speaks"
}



func main(){

    a := Animal("Tom")
    a.Speak()
    
    
    d := Dog{
        Animal{"Charlie"},
        "Bulldog"
    }
    
    d.Speak()

}
```


