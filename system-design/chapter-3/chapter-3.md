# Chapter 3 — A Framework for System Design Interviews

> Source: *System Design Interview* (Alex Xu), Chapter 3. Notes are my own summaries.

## Contents

1. [What the Interviewer Looks For](#1-what-the-interviewer-looks-for)
2. [The 4-Step Framework](#2-the-4-step-framework)
3. [Step 1 — Understand the Problem](#3-step-1--understand-the-problem)
4. [Step 2 — High-Level Design](#4-step-2--high-level-design)
5. [Step 3 — Deep Dive](#5-step-3--deep-dive)
6. [Step 4 — Wrap Up](#6-step-4--wrap-up)
7. [Dos and Don'ts](#7-dos-and-donts)
8. [Time Management](#8-time-management)
9. [Core Mental Models](#9-core-mental-models)
10. [Interview Cheat Sheet](#10-interview-cheat-sheet)

---

System design questions are **intentionally vague**. Nobody expects a full production
system designed in one hour, and there is no single perfect answer.

What is really being evaluated is **how you think**: how you break down a fuzzy problem,
work with another person, and reach reasonable decisions. Communication and collaboration
count as much as technical knowledge.

---

## 1. What the Interviewer Looks For

- **Problem-solving** — turning a vague prompt into a structured plan
- **Communication** — making your reasoning visible
- **Collaboration** — working *with* the interviewer, like a teammate
- **Handling ambiguity** — staying calm when the problem is underspecified
- **Asking good questions** — clarifying instead of guessing
- **Making assumptions** — explicit, reasonable, stated out loud
- **Making tradeoffs** — every choice costs something; name the cost
- **Working under pressure** — steady progress within limited time
- **Responding to feedback** — adapting to hints instead of defending your first idea
- **Prioritization** — spending time where it matters most

### Red flags

- Over-engineering
- Narrow-mindedness (only one way of thinking)
- Stubbornness
- Jumping to conclusions
- Poor communication
- Spending too much time on low-value details

---

## 2. The 4-Step Framework

1. Understand the problem and establish design scope
2. Propose a high-level design and get buy-in
3. Design deep dive
4. Wrap up

### Mental model

```
Clarify
   ↓
Design
   ↓
Deep Dive
   ↓
Critique
```

---

## 3. Step 1 — Understand the Problem

Do **not** start solving. The urge to answer fast is the most common trap; a quick wrong
answer is worse than a slow right one.

What to do:

- Ask questions until the requirements are clear.
- Identify assumptions and **write them down** so you can refer back to them.
- Pin down scale, scope, and constraints.

### Useful question categories

| Category | Example question |
|---|---|
| Features | What are the must-have features? |
| Users | Who uses it? |
| Volume | How many users? |
| Scale | What traffic do we expect? |
| Growth | How fast will it grow in 3 months / 1 year / 3 years? |
| Platforms | Web, mobile, or both? |
| Constraints | Any tech stack, budget, or team limits? |
| Reuse | Can we build on existing services? |

### Example: how a candidate clarifies (news feed prompt)

- "Is this mobile, web, or both?"
- "What are the most important features?"
- "How many users, and how many are active daily?"
- "Do we expect the number of users to grow quickly?"

Notice: only questions, no design yet. The answers set the scope for everything after.

> **Understand before designing.**
