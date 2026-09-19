Interfaces aur design

- Interface implicit hoti hai: method hai to type interface satisfy karta hai.
- "Accept interfaces, return structs".
- Interface chhoti rakho (1-2 methods), consumer ke package me define karo.
- Pointer receiver jab state badalni ho ya struct bada ho, warna value receiver.
- Zero value useful banao: `var mu sync.Mutex` bina init ke kaam karta hai.
- Generics tab use karo jab same logic multiple types pe ho (Map, Sum, Stack), warna interface kaafi hai.
- Project layout: `cmd/` (binaries), root packages (logic), internal/ agar private rakhna ho.

Code: structs/, generics/, lru/
