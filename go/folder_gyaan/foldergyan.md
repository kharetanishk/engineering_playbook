Rule 1: One directory = one package
fundamentals/
├── main.go
├── chut.go
└── lund.go

All files must start with:

package main

or all must start with:

package fundamentals

You cannot mix package names in the same directory.

Rule 2: Different directory = different package
fundamentals/
├── main.go      (package main)
│
└── lund/
    └── lund.go  (package lund)

Now main and lund are separate packages.

Rule 3: Capital letter = Exported (Public)
package lund

func Hello() {}   // Public
func hello() {}   // Private

From another package:

lund.Hello() // ✅

lund.hello() // ❌
Rule 4: You import packages, not functions

Many beginners think:

import Hello

But Go imports packages:

import "engineering_playbook/fundamentals/lund"

Then access exported members:

lund.Hello()

Just like:

fmt.Println()
strings.Split()
http.ListenAndServe()
fmt = package
Println = exported function
Rule 5: Import path = Module Name + Folder Path

If go.mod contains:

module engineering_playbook

and you have:

fundamentals/
└── lund/

then:

import "engineering_playbook/fundamentals/lund"