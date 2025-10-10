🅿️ 𝐒𝐦𝐚𝐫𝐭𝐏𝐚𝐫𝐤

𝐀 𝐆𝐨-𝐩𝐨𝐰𝐞𝐫𝐞𝐝 𝐞𝐝𝐮𝐜𝐚𝐭𝐢𝐨𝐧𝐚𝐥 𝐩𝐫𝐨𝐣𝐞𝐜𝐭 𝐭𝐡𝐚𝐭 𝐦𝐨𝐝𝐞𝐥𝐬 𝐚 𝐩𝐚𝐫𝐤𝐢𝐧𝐠 𝐥𝐨𝐭 𝐮𝐬𝐢𝐧𝐠 𝐚𝐧 𝐨𝐩𝐞𝐧-𝐚𝐝𝐝𝐫𝐞𝐬𝐬𝐢𝐧𝐠 𝐡𝐚𝐬𝐡 𝐦𝐚𝐩.


🌱 𝗢𝘃𝗲𝗿𝘃𝗶𝗲𝘄
SmartPark simulates a parking system where car number plates act as 𝐤𝐞𝐲𝐬 in a hash table.
Instead of chaining, we use 𝗼𝗽𝗲𝗻 𝗮𝗱𝗱𝗿𝗲𝘀𝘀𝗶𝗻𝗴 and 𝗽𝗿𝗼𝗯𝗶𝗻𝗴 (linear probing by default) to resolve collisions.
Deletion is handles with a 𝘁𝗼𝗺𝗯𝘀𝘁𝗼𝗻𝗲 marker so lookups don't break probe chains.

It started as a simple data structure exercise (insert / lookup / delete with tombstones and probing) and has grown into a modular, testable project that can later be extended via an API or UI.

This README provides everything you need to run the demo and understand the core implementation concepts.


🎯 𝗚𝗼𝗮𝗹𝘀: 
- Learn how open addressing works (insert, lookup, delete)
- Visualize collision and probe behavior
- Provide a clean extendible Go codebase
- Enable future expansion (rehashing, different probing strategies, API, frontend)


🧠 𝗞𝗲𝘆 𝗖𝗼𝗻𝗰𝗲𝗽𝘁𝘀 (𝘀𝗵𝗼𝗿𝘁):
- 𝗛𝗮𝘀𝗵 𝗳𝘂𝗻𝗰𝘁𝗶𝗼𝗻 → converts a plate (string) to an index
Formula: 𝗵𝗮𝘀𝗵(𝗽𝗹𝗮𝘁𝗲) % 𝘀𝗶𝘇𝗲 
- 𝗢𝗽𝗲𝗻 𝗮𝗱𝗱𝗿𝗲𝘀𝘀𝗶𝗻𝗴 → all entries live in the table itself, no linked lists 
- 𝗣𝗿𝗼𝗯𝗶𝗻𝗴 → when a slot is occupied, move to another index (linear, quadratic, double hashing)
- 𝗧𝗼𝗺𝗯𝘀𝘁𝗼𝗻𝗲 (𝗗𝗘𝗟𝗘𝗧𝗘𝗗) → a special marker used during deletion so further lookups keep probing
- 𝗖𝗶𝗿𝗰𝘂𝗹𝗮𝗿 𝗽𝗿𝗼𝗯𝗶𝗻𝗴 - 𝗶𝗻𝗱𝗲𝘅 = (𝗶𝗻𝗱𝗲𝘅 + 𝟭) % 𝗹𝗲𝗻(𝘁𝗮𝗯𝗹𝗲) ensures wrap-around.


⚙️ 𝗙𝗲𝗮𝘁𝘂𝗿𝗲𝘀: 
1. Insert a car (uses probing) 
2. Lookup car by plate 
3. Delete a car (marks slot as DELETED / tombstone) 
4. Print/Visualize the parking lot state (EMPTY, DELETED, or plate) 
5. Modular layout - easily extended with: 
- alternative probing strategies 
- resizing/rehashing 
- REST API & frontend demo


𝗤𝘂𝗶𝗰𝗸 𝘀𝘁𝗮𝗿𝘁
𝗥𝗲𝗾𝘂𝗶𝗿𝗲𝗺𝗲𝗻𝘁𝘀: 
1. Go 1.18+ installed
2. Git


𝐂𝐥𝐨𝐧𝐞 & 𝐬𝐞𝐭𝐮𝐩:
git clone https://github.com/<your-username>/SmartPark.git
cd SmartPark
go mod tidy


𝗥𝘂𝗻 𝘁𝗵𝗲 𝗖𝗟𝗜 𝗱𝗲𝗺𝗼:
go run cmd/main.go


𝗬𝗼𝘂’𝗹𝗹 𝘀𝗲𝗲 𝗮𝗻 𝗶𝗻𝘁𝗲𝗿𝗮𝗰𝘁𝗶𝘃𝗲 𝗺𝗲𝗻𝘂:
--- Parking Lot Menu ---
1. Park a car
2. Find a car
3. Remove a car
4. Show parking lot (debug)
5. Exit
Enter choice:


𝐈𝐦𝐩𝐥𝐞𝐦𝐞𝐧𝐭𝐚𝐭𝐢𝐨𝐧 𝐧𝐨𝐭𝐞𝐬 (𝐭𝐢𝐩𝐬 𝐟𝐨𝐫 𝐜𝐨𝐧𝐭𝐫𝐢𝐛𝐮𝐭𝐨𝐫𝐬)
- 𝗜𝗻𝘀𝗲𝗿𝘁 should place the item in the first EMPTY or DELETED slot found while probing.
- 𝐋𝐨𝐨𝐤𝐮𝐩 must skip DELETED slots (i.e., continue probing) and stop only at EMPTY or when it loops back to the start.
- 𝐃𝐞𝐥𝐞𝐭𝐞 marks the slot as DELETED, not EMPTY. If you set it to EMPTY, you break probe chains.
- 𝐑𝐞𝐡𝐚𝐬𝐡 (future task): When tombstones accumulate (or load factor > 0.6–0.7), rebuild the table into a new array, reinserting non-deleted entries to remove tombstones and improve performance.
- 𝐏𝐫𝐨𝐛𝐢𝐧𝐠 𝐬𝐭𝐫𝐚𝐭𝐞𝐠𝐢𝐞𝐬: linear probing is easiest and cache-friendly, quadratic reduces primary clustering, double hashing minimizes clustering more but needs two hash functions.


𝐑𝐨𝐚𝐝𝐦𝐚𝐩 / 𝐍𝐞𝐱𝐭 𝐬𝐭𝐞𝐩𝐬

- ✅ Insert/lookup/delete (linear probing + tombstone) — done
- ✅ CLI demo + print state — done
- ➕ Unit tests covering insert/lookup/delete/reuse-of-deleted-slot
- ➕ Rehashing and dynamic resizing
- ➕ Alternative probing startegies - quadratic / double hashing
- ➕ Add a small REST API to exercise the hash table programmatically
- ➕ Tiny frontend visualization HTML + JS to show slot states and probe path
- ➕ Performance benchmarks and comparison doc

𝐂𝐨𝐧𝐭𝐫𝐢𝐛𝐮𝐭𝐢𝐧𝐠

1. For the repo and create a feature branch: 
𝗴𝗶𝘁 𝗰𝗵𝗲𝗰𝗸𝗼𝘂𝘁 -𝗯 𝗳𝗲𝗮𝘁/𝗿𝗲𝗵𝗮𝘀𝗵𝗲𝗱-𝘁𝗮𝗯𝗹𝗲
2. Write tests for the new feature.
3. Open a pull request with a clear description and link to related issue.
Be kind, leave tests, and keep changes small and focused.


🧰 Tech Stack

Language: Go 1.18+
Interface: Command Line (CLI)
Architecture: Modular design (cmd/, pkg/)
Future Extensions: REST API, Frontend Visualization


𝐋𝐢𝐜𝐞𝐧𝐬𝐞
This project is licensed under the MIT License — you’re free to use, modify, and distribute it with proper attribution.
See the LICENSE file for full details.

𝐂𝐨𝐧𝐭𝐚𝐜𝐭 / 𝐂𝐫𝐞𝐝𝐢𝐭𝐬
Created by 𝐒𝐮𝐬𝐚𝐧 𝐊𝐚𝐦𝐚𝐮 - documenting a learning journey in public.
If you want help extending the project (API, frontend, demos), message me on Linkedin or open an issue in the repository.
