🅿️ SmartPark

A Go project exploring open addressing, collisions, and elegant system design.


🌱 Overview

SmartPark is a small educational Go project that models a parking lot using an open-addressing hash map.

It started as a data-structure exercise (insert / lookup / delete with tombstones and probing) and is growing into a modular, testable project that can later extend into an API or a UI.

While educational, SmartPark is inspired by real-world parking logic — building systems that are both efficient and intelligent.

This README will help you get started, run the demo, and understand the core ideas behind the implementation.


🧠 Core Idea

SmartPark simulates a parking system where car number plates are keys in a hash table.
Instead of chaining, we use open addressing and probing (linear probing by default) to resolve collisions.
Deletion is handled with a tombstone marker so lookups don’t break probe chains.

🎯 Goals

Learn how open addressing works (insert, lookup, delete)

Visualize collision and probe behavior

Provide a clean, extensible codebase (rehashing, different probing strategies, API, frontend)

🧩 Key Concepts
Concept	Description
Hash function	Converts a plate (string) to an index: hash(plate) % size
Open addressing	All entries live in the table itself (no linked lists)
Probing	When a slot is full, move to another index (linear, quadratic, or double hashing)
Tombstone ("DELETED")	Special marker used during deletion so lookups continue probing
Circular probing	index = (index + 1) % len(table) ensures wrap-around
⚙️ Features

Insert a car (uses probing)

Lookup car by plate

Delete a car (marks slot as "DELETED")

Print/Visualize the parking lot state (EMPTY, DELETED, or plate)

Modular layout for adding:

Alternative probing strategies

Resizing/rehashing

REST API and frontend demo

🚀 Quick Start
Requirements

Go 1.18+

Git

Clone & Setup
git clone https://github.com/<your-username>/SmartPark.git
cd SmartPark
go mod tidy

Run the CLI Demo
go run cmd/main.go


You’ll see an interactive menu:

--- Parking Lot Menu ---
1. Park a car
2. Find a car
3. Remove a car
4. Show parking lot (debug)
5. Exit
Enter choice:

🧱 Implementation Notes

Insert → place item in the first EMPTY or DELETED slot while probing.

Lookup → skip "DELETED" slots and stop only at EMPTY or full loop.

Delete → mark slot as "DELETED", not EMPTY (to preserve probe chains).

Rehashing (future) → rebuild when tombstones accumulate or load factor > 0.7.

Probing strategies →

Linear: simplest, cache-friendly

Quadratic: reduces clustering

Double hashing: uses two hash functions, minimizes collisions

🛠 Roadmap

✅ Insert / Lookup / Delete (linear probing + tombstones)

✅ CLI demo + print state

🔜 Unit tests for insert / lookup / delete / reuse-of-deleted-slot

🔜 Rehashing and dynamic resizing

🔜 Alternative probing strategies (quadratic / double hashing)

🔜 REST API for remote interaction

🔜 Frontend visualization (HTML + JS)

🔜 Performance benchmarks and comparison doc

🤝 Contributing

Fork the repo and create a feature branch:

git checkout -b feat/rehash-table


Write tests for new features.

Open a pull request with a clear description and related issue.

Be kind, keep PRs focused, and include tests 💛

📜 License

This project is licensed under the MIT License — you’re free to use, modify, and distribute it with proper attribution.
See the LICENSE
 file for full details.

💬 Contact / Credits

Created by Susan Kamau — documenting a learning journey in public.
If you’d like to help extend the project (API, frontend, demos), message me on LinkedIn
 or open an issue in the repository.