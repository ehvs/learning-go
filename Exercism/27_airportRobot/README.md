# About **Interfaces**

An interface is like a painter's color palette, where each color represents a **method**.
Different artists (structs) must use all the colors available in the palette (implement all the methods) to create their art, but each artist can use those colors differently.
This ensures that any artist (struct) that uses the palette (interface) will have all the required tools (methods) available.

In short, an **interface** is a **contract** that guarantees consistency.
Any struct that implements the interface must implement all of its methods, ensuring that the struct can be used wherever that interface is expected