# 🏗️ SOLID Principles in Go

The **SOLID** principles are a set of five design guidelines that make software **maintainable, scalable, and flexible**.  

---

## 📌 SOLID stands for:

1. **S** – Single Responsibility Principle (SRP)  
2. **O** – Open/Closed Principle (OCP)  
3. **L** – Liskov Substitution Principle (LSP)  
4. **I** – Interface Segregation Principle (ISP)  
5. **D** – Dependency Inversion Principle (DIP)  

---

## 1️⃣ Single Responsibility Principle (SRP)

> A class (or struct in Go) should have **only one reason to change**.  

**❌ Bad Example:**  
One struct handling multiple responsibilities:  

```go
type Employee struct {
    Name    string
    Salary  float64
    Address string
}
```


**✅ Good Example:**  
Split into smaller structs with clear responsibilities:

```go
type EmployeeInfo struct {
    Name   string
    Salary float64
}

type EmployeeAddress struct {
    Address string
}

func (e EmployeeInfo) GetSalary() float64 {
    return e.Salary
}

func (e EmployeeAddress) GetAddress() string {
    return e.Address
}
```



## 2️⃣ Open-Closed Principle (OCP)

Software entities should be open for extension but closed for modification.

**Example: Payment system supporting multiple payment methods.**

```go
package main

import "fmt"

type PaymentMethod interface {
    Pay()
}

type Payment struct{}

func (p Payment) Process(pm PaymentMethod) {
    pm.Pay()
}

type CreditCard struct {
    amount float64
}

func (cc CreditCard) Pay() {
    fmt.Printf("Paid %.2f using Credit Card\n", cc.amount)
}

type PayPal struct {
    amount float64
}

func (pp PayPal) Pay() {
    fmt.Printf("Paid %.2f using PayPal\n", pp.amount)
}

func main() {
    p := Payment{}

    cc := CreditCard{12.23}
    p.Process(cc)

    pp := PayPal{22.33}
    p.Process(pp)
}
```

✅ Adding new payment methods (like UPI, NetBanking, Crypto) requires no changes to the Payment struct.

## 3️⃣ Liskov Substitution Principle (LSP)

Subclasses should be substitutable for their base classes without breaking the system.

✅ Good Example:
All payment methods implement PaymentMethod correctly.

```go
type PaymentMethod interface {
    Pay(amount float64)
}

type CreditCard struct{}

type PayPal struct{}


func (c CreditCard) Pay(amount float64) {
    fmt.Println("Paid with Credit Card:", amount)
}

func (p PayPal) Pay(amount float64) {
    fmt.Println("Paid with PayPal:", amount)
}
```


❌ Bad Example (Violation):
GiftCard panics if balance is insufficient → breaks expected behavior.
```go
type GiftCard struct {
    Balance float64
}

func (g GiftCard) Pay(amount float64) {
    if amount > g.Balance {
        panic("Insufficient balance on gift card!") // 🚨 Breaks LSP
    }
    fmt.Println("Paid with Gift Card:", amount)
}

```
## 4️⃣ Interface Segregation Principle (ISP)

Methods should not be forced to depend on interfaces they don’t use.

❌ Bad Example: One interface forces unnecessary methods.

```go
type Worker interface {
    Work()
    Eat()
}
```

✅ Good Example: Split into smaller interfaces.

```go
type Worker interface {
    Work()
}

type Eater interface {
    Eat()
}

type Human struct{}

func (h Human) Work() { fmt.Println("Working...") }
func (h Human) Eat()  { fmt.Println("Eating...") }

type Robot struct{}

func (r Robot) Work() { fmt.Println("Working...") }
// Robot doesn’t need Eat()
```

## 5️⃣ Dependency Inversion Principle (DIP)

High-level modules should not depend on low-level modules. Both should depend on abstractions.

❌ Bad Example: Switch depends directly on LightBulb.
```go
type LightBulb struct{}

func (l LightBulb) On()  { fmt.Println("LightBulb On") }
func (l LightBulb) Off() { fmt.Println("LightBulb Off") }

type Switch struct {
    bulb LightBulb
}
```


✅ Good Example: Use abstraction (SwitchableDevice).

```go
type SwitchableDevice interface {
    On()
    Off()
}

type LightBulb struct{}

func (l LightBulb) On()  { fmt.Println("LightBulb On") }
func (l LightBulb) Off() { fmt.Println("LightBulb Off") }

type Fan struct{}

func (f Fan) On()  { fmt.Println("Fan On") }
func (f Fan) Off() { fmt.Println("Fan Off") }

type Switch struct {
    device SwitchableDevice
}

func (s Switch) Toggle() {
    s.device.On()
    s.device.Off()
}
```

✅ Summary

    SRP → One responsibility per class/struct.

    OCP → Extend behavior without modifying existing code.

    LSP → Subclasses must be usable as their base class.

    ISP → Prefer smaller, specific interfaces.

    DIP → Depend on abstractions, not concrete implementations.



## Simple Understanding:

    1. Single Responsibility Principle: 
        a class/struct should have one reason to change
    
        Bad Example: One struct handling multiple responsibilities
        Good Example: Split into smaller structs with clear responsibilities

    2. Open Closed Principle:
        Entities should be open for extension but close for modification

        Example: Payment system supporting multiple payment methods

    3. Liskov Substitution Principle:
        Subclasses should be substitutable for their base classes

        Bad Example (Violation): GiftCard panics if balance is insufficient → breaks expected behavior.
        Good Example: All payment methods implement PaymentMethod correctly.

    4. Interface Segregation Principle:
        Methods should not be forced to depend on interfaces they don't use

        Bad Example: One interface forces unnecessary methods.
        Good Example: Split into smaller interfaces.

    5. Dependency Inversion Principle:
        High-level modules(struct) should not depend on low-level modules.

        Bad Example: Switch depends directly on LightBulb
        Good Example: Use abstraction (SwitchableDevice).