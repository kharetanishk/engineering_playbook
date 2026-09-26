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

---

## 4. Step 2 — High-Level Design

Produce an initial blueprint and check it with the interviewer.

- Keep it **high level** — the big picture, not the details.
- Think out loud so the interviewer can follow and steer.
- Draw a simple diagram when it helps.
- Ask for feedback; treat the interviewer as a teammate, not an examiner.
- Get agreement on the approach **before** going deeper.
- Do a quick back-of-the-envelope estimate if the scale affects the approach (see Chapter 2).

How much API or data-model detail to include depends on the scope of the interview.
When unsure, ask.

> **Draw the big picture before diving into details.**

---

## 5. Step 3 — Deep Dive

By now you and the interviewer have agreed on the overall shape. Use the remaining time
on the parts that matter most.

- Identify the most important components.
- Prioritize using the requirements and the interviewer's feedback.
- Go into important details, look for bottlenecks, and discuss tradeoffs.
- Spend time on areas that show engineering judgment.
- Skip unnecessary detail.

The interviewer may steer you to a specific area — follow their lead. Otherwise, pick
what is riskiest or most relevant to the requirements.

> **Not every component deserves equal depth.**
>
> **Deep dive where it matters.**

---

## 6. Step 4 — Wrap Up

Use the last few minutes to step back and evaluate your own work.

- Summarize the design briefly.
- Point out bottlenecks and possible improvements.
- Discuss failure and error cases at a high level.
- Mention monitoring and operational concerns at a high level.
- Say how the design could handle the next order of magnitude of growth.
- Note what you would improve with more time.
- Never claim the design is perfect.

> **Critique your own design.**

---

## 7. Dos and Don'ts

### DO

- Ask clarifying questions.
- Understand requirements.
- State assumptions.
- Think out loud.
- Communicate continuously.
- Ask for feedback.
- Suggest alternatives when useful.
- Prioritize critical areas.
- Discuss tradeoffs.
- Manage time.

### DON'T

- Jump into a solution immediately.
- Assume requirements.
- Over-engineer.
- Dive too deeply too early.
- Think silently for long periods.
- Get stuck on one component.
- Claim the design is perfect.
- Give up when stuck.
- Ignore interviewer feedback.

---

## 8. Time Management

Rough split for a 45-minute interview:

| Step | Time |
|---|---:|
| Understand problem | 3–10 min |
| High-level design | 10–15 min |
| Deep dive | 10–25 min |
| Wrap up | 3–5 min |

> These are rough guidelines, not strict rules.

Actual allocation depends on:

- Problem scope
- The interviewer
- Candidate level
- Direction of the discussion

---

## 9. Core Mental Models

1. **Clarify → Design → Deep Dive → Critique**
2. **Requirements before architecture.**
3. **Big picture before details.**
4. **Communicate your thinking instead of thinking silently.**
5. **Optimize for useful discussion, not design perfection.**

---

## 10. Interview Cheat Sheet

| Question | Answer |
|---|---|
| Goal of the interview? | Show how you solve an ambiguous problem and collaborate, not deliver a perfect design. |
| Before designing? | Ask questions, clarify scope and scale, and write down your assumptions. |
| During high-level design? | Sketch the big picture, think aloud, and get the interviewer's agreement before going deeper. |
| How to pick a deep-dive area? | Follow the interviewer's hints, else choose what is most critical to the requirements. |
| What to discuss at the end? | Bottlenecks, failure cases, monitoring, next-scale growth, and what you'd improve. |
| Common red flags? | Over-engineering, stubbornness, jumping to conclusions, poor communication, low-value detail. |
| If you get stuck? | Say so, explain your thinking, and ask for a hint — don't go silent or give up. |
| Why does communication matter? | The interviewer can only evaluate reasoning they can hear, and it shows you'd be a good teammate. |

---

# Chapter 3 — 2-Minute Revision

```text
1. Clarify requirements.
2. State assumptions.
3. Create high-level design.
4. Get interviewer buy-in.
5. Deep dive into critical areas.
6. Discuss tradeoffs.
7. Manage time.
8. Wrap up with bottlenecks + improvements.
9. Communicate throughout.
```
