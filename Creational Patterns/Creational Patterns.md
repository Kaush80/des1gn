# Creational Patterns

---

## Builder Pattern

### Purposes:

- Encapsulates an object's construction process along with specifying the various parts of a complex API.
- Enables flexible creation of an object that can have many different representations.
- Increases code readability for complex types.

### Scenarios

- Objects that have complex APIs, multiple constructor options, and several different possible representations.

## Factory Pattern

### Purpose

-Allows for the construction of objects when the types of those objects is not predetermined at runtime.

### Scenarios

-Produces code that is more readable when there are multiple ways of creating a particular object.
-Situations where object creation needs to be flexible and cannot be known beforehand.

## Singleton pattern

### Purpose

- The singleton pattern restricts the instantiation of a class to a single instance and provides global access
- Allows for lazy initialization of the class

### Scenarios

- Situations where you want to ensure there is only one instance of a class: logging, configuration, telemetry, debugging
