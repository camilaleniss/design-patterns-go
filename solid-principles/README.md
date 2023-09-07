# SOLID Design Principles

Five design principles introduced by Uncle Bob :)

## 1. Single Responsibility Principle

A type/class or whatever should have a single primary responsibility.  
You can create another type to handle another responsibility.  
Maybe break into packages as well -> separation of concerns  

## 2. Open-Closed Principle

Types should be open for extension, but close for modification.  
It's well illustrated using Specification pattern.  

## 3. Liskov Substitution Principle

Stands that if there is any inheritance or extension of a class  
It should continue working and behaving correctly  
Based on the principles and assumptions it relies on.  
It's not truly applicable in Go, but we made a case that applies  

## 4. Interface Segregation Principle

You shouldn't put too much in an interface  
Makes sense to break into several interfaces  
So we won't end up having to implement and leaving alone methods that are not neccessary  

## 5. Dependency Inversion Principle

Dependency inversion != Dependency injection  
High level modules should not depend on low level modules
Both should depend on abstractions  
