# 🅿️ **SmartPark**
### A Go project exploring **open addressing**, **collisions**, and **elegant system design**.

---

## 🌱 1. Overview

**SmartPark** is a small educational Go project that models a parking lot using an **open-addressing hash map**.

It began as a **data-structure exercise** (insert / lookup / delete with tombstones and probing) and is evolving into a **modular, testable project** that can extend into an **API or UI**.

While educational, SmartPark draws inspiration from **real-world parking logic** - systems that are both efficient and intelligent.

> This README will help you get started, run the demo, and understand the core ideas behind the implementation.

---

## 🧠 2. Core Idea

SmartPark simulates a parking system where **car number plates** act as **keys** in a hash table.

- Instead of chaining, it uses **open addressing** and **probing** (linear probing by default) to resolve collisions.  
- Deletion is handled with a **tombstone marker** so lookups don’t break probe chains.

---

## 🎯 3. Goals

- Learn how **open addressing** works (insert, lookup, delete)  
- Visualize **collisions** and **probe behavior**  
- Provide a **clean, extensible codebase** (rehashing, different probing strategies, API, frontend)

---

## 🧩 4. Key Concepts

| **Concept** | **Description** |
|--------------|-----------------|
| **Hash function** | Converts a plate (string) to an index: `hash(plate) % size` |
| **Open addressing** | All entries live in the table itself (no linked lists) |
| **Probing** | When a slot is full, move to another index (linear, quadratic, or double hashing) |
| **Tombstone ("DELETED")** | Special marker used during deletion so lookups continue probing |
| **Circular probing** | `index = (index + 1) % len(table)` ensures wrap-around |

---

## ⚙️ 5. Features

1. 🚘 **Insert a car** (uses probing)  
2. 🔍 **Lookup a car** by plate  
3. ❌ **Delete a car** (marks slot as `"DELETED"`)  
4. 🧾 **Print/Visualize** the parking lot state (`EMPTY`, `DELETED`, or plate)  
5. 🧱 **Modular layout** for adding:  
   - Alternative probing strategies  
   - Resizing / rehashing  
   - REST API & frontend demo

---

## 🚀 6. Quick Start

### Requirements
- Go 1.18+  
- Git  

### Clone & Setup
```bash
git clone https://github.com/<your-username>/SmartPark.git
cd SmartPark
go mod tidy

```

### Run the CLI Demo
```
go run cmd/main.go

```
### You’ll see an interactive menu:
```
--- Parking Lot Menu ---
1. Park a car
2. Find a car
3. Remove a car
4. Show parking lot (debug)
5. Exit
Enter choice:
```
---
## 🧱 7. Implementation Notes  

- **Insert** → place item in the first `EMPTY` or `DELETED` slot while probing.  
- **Lookup** → skip `"DELETED"` slots and stop only at `EMPTY` or full loop.  
- **Delete** → mark slot as `"DELETED"`, not `EMPTY` (to preserve probe chains).  
- **Rehashing (future)** → rebuild when tombstones accumulate or load factor > 0.7.  
- **Probing strategies** →  
  - **Linear:** simplest, cache-friendly  
  - **Quadratic:** reduces clustering  
  - **Double hashing:** uses two hash functions, minimizes collisions  

---

## 🛠 8. Roadmap  

- ✅ Insert / Lookup / Delete (linear probing + tombstones)  
- ✅ CLI demo + print state  
- 🔜 Unit tests for insert / lookup / delete / reuse-of-deleted-slot  
- 🔜 Rehashing and dynamic resizing  
- 🔜 Alternative probing strategies (quadratic / double hashing)  
- 🔜 REST API for remote interaction  
- 🔜 Frontend visualization (HTML + JS)  
- 🔜 Performance benchmarks and comparison doc  

---

## 🤝 9. Contributing  

1. Fork the repo and create a feature branch:  
   ```bash
   git checkout -b feat/rehash-table
   
2. Write tests for new features.

3. Open a pull request with a clear description and related issue.

Be kind, keep PRs focused, and include tests 💛

---

## 📜 10. License

This project is licensed under the **MIT License** - you’re free to use, modify, and distribute it with proper attribution.
See the **LICENSE** file for full details.

---

## 💬 11. Contact / Credits

Created by Susan Kamau — documenting a learning journey in public.
If you’d like to help extend the project (API, frontend, demos), message me on **LinkedIn** or open an issue in the repository.

---





 