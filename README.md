# The Celebrity Problem

> Striver SDE Sheet 091 · Day 14 Stack & Queue II — every "does A know B?" answer eliminates somebody; n−1 questions leave one candidate to verify. Go, stdlib only, tested.

## Problem

At a party of n people, a celebrity is known by everyone but knows no one. Given the `knows` relation as a matrix, find the celebrity (or report none) using O(n) relation queries — not the n² of checking everyone against everyone.

## Approach

The key asymmetry: one answer always eliminates one person. If A knows B, A is no celebrity; if A doesn't know B, B isn't (someone doesn't know them). So sweep once holding a candidate — ask "does the candidate know i?" and switch to i on yes. After n−1 eliminations exactly one survivor remains, but elimination only proves the *others* aren't celebrities, so a verification pass confirms the survivor is known by all and knows none. (The sheet files this under stacks: push everyone, pop pairs, eliminate one, push the survivor — identical logic, explicit container.)

**Complexity:** O(n) queries, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
