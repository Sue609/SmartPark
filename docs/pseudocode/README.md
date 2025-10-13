# 🅿️ **SmartPark**
## 🧠 Pseudocode Overview
This folder contains all the **pseudocode blueprints** for **SmartPark** - the foundation of the project's logic before any Go code is written.
Each pseudocode file represents a single function or component, keeping the system modular, clear, and easy to test conceptually.

---

## 📂 Folder Structure
```
/pseudocode
│
├── hash.md        → Defines how number plates are converted into hash indexes.
├── insert.md           → Handles adding (parking) a car into the lot.
├── lookup.md           → Checks if a car already exists in the parking lot.
├── remove.md           → Removes a car from the parking lot.
├── printLot.md             → Prints the current parking lot state for debugging.
└── main.md        → Represents the user-facing program flow and menu logic.
```
---

## ⚙️ Flow Summary

1. **Start the program** -> The main pseudocode initializes a parking lot - array.
2. **User selects an option** -> Park, find, remove, or view the lot.
3. **Functions interact** -> 
- **hash** computes the bucket.
- **insert** adds a plate.
- **lookup** checks for existence.
- **remove** deletes te number plate.
- **printLot** shows current state
4. **Loop continues** until user exists.

---

## 🎯 Purpose

The pseudocode stage allows us to **design logic first**, before diving into the syntax or adge-case handling.
This mindset **"think in systems before coding"** makes debugging and scalability far easier later on.
