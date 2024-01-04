# AI Engineering Residency – Complete Summary

## Vision

The aim is **not** to complete 25 portfolio projects.

The aim is to develop the engineering judgement expected of senior software engineers by understanding:

- Why architectures are designed the way they are.
- Why particular technologies are chosen.
- What trade-offs exist.
- How systems evolve over time.
- How to communicate engineering decisions effectively.

The programme is designed as an **Engineering Residency**, rather than a coding bootcamp.

AI is used as an engineering partner, not as a replacement for understanding.

---

# Objectives

By the end of the residency, every participant should be able to:

- Design software from ambiguous requirements.
- Produce architecture before implementation.
- Defend engineering decisions.
- Understand trade-offs.
- Build production-ready software.
- Review code and architecture.
- Collaborate effectively.
- Explain every part of a system, even if they didn't implement it.

> **Guiding Principle:**  
> *If you can't explain it, you don't understand it.*

---

# Programme Structure

The residency consists of **25 progressively more challenging projects**.

Projects begin with simple engineering concepts and become increasingly realistic.

Each project introduces one or two new engineering concepts.

The focus is **depth rather than quantity**.

---

# Project Modes

## Projects 1–3 — Individual Projects

Everyone completes the project independently.

The objective is to establish strong engineering fundamentals before collaborating.

Each participant receives the same project brief and is responsible for the entire solution.

Deliverables include:

- Requirements
- Architecture
- Folder structure
- Testing strategy
- ADRs
- Retrospective
- Lessons learned

After completion everyone presents and defends their decisions.

---

## Projects 4–25 — Collaborative Projects

Everyone works from a shared Git repository.

Each engineer:

- Creates their own feature branch.
- Raises Pull Requests.
- Reviews other Pull Requests.
- Participates in architecture discussions.
- Owns a specific feature or component.

Leadership rotates every project.

---

## Architecture Competitions

Every few projects the team independently designs competing solutions before agreeing on a final architecture.

The purpose is to develop engineering judgement rather than converging on the first acceptable solution.

---

# Engineering Workflow

Every collaborative project follows the same workflow.

## Phase 1 — Project Brief

Provided:

- Problem statement
- Language(s)
- Technology stack
- Functional requirements
- Non-functional requirements
- Constraints

No architecture is provided.

---

## Phase 2 — Individual Design

Every engineer independently produces:

- Requirements analysis
- Architecture
- Component diagrams
- Data model
- API design
- Testing strategy
- ADRs
- Trade-off analysis

No coding.

---

## Phase 3 — Architecture Review

Each engineer presents their proposed solution.

Discussion includes:

- Simplicity
- Maintainability
- Scalability
- Risks
- Trade-offs
- Alternative approaches

The objective is thoughtful engineering discussion—not finding a single "correct" answer.

---

## Phase 4 — Team Architecture

The team agrees on one design.

Document:

- Chosen solution
- Alternatives considered
- Reasons for the decision
- Trade-offs

---

## Phase 5 — Implementation

The Lead Engineer allocates work.

Each engineer owns a component.

Examples:

- API
- Persistence
- Authentication
- Infrastructure
- Testing
- Frontend
- Background jobs

Everyone develops on their own branch.

Changes are merged through Pull Requests.

---

## Phase 6 — Code Review

Every Pull Request is reviewed.

Reviews focus on:

- Readability
- Simplicity
- Security
- Testing
- Maintainability
- Architecture

The goal is learning.

---

## Phase 7 — Production Review

Ask:

> Would we deploy this?

Review:

- Logging
- Monitoring
- Configuration
- Security
- Error handling
- Recovery
- Deployment
- Performance

---

## Phase 8 — Retrospective

Discuss:

- What worked?
- What assumptions failed?
- What surprised us?
- What would we redesign?
- What did we learn?

Document all lessons.

---

# Rotating Leadership

Every project has a different Lead Engineer.

Responsibilities include:

- Planning
- Facilitating meetings
- Coordinating implementation
- Allocating work
- Keeping discussions focused
- Driving retrospectives
- Ensuring documentation is completed

Everyone should experience technical leadership.

---

# Engineering Constitution

1. Understand the problem before writing code.
2. Choose the simplest solution that satisfies today's requirements.
3. Document significant decisions using ADRs.
4. Challenge ideas—not people.
5. If you cannot explain it, you do not understand it.
6. Measure before optimising.
7. Every Pull Request should teach something.
8. Documentation is part of the deliverable.
9. AI supports engineering judgement—it never replaces it.
10. Leave every project better than you found it.

---

# AI Philosophy

We use AI extensively.

AI accelerates engineering but never replaces engineering judgement.

## Claude Code

Use for:

- Requirements refinement
- Architecture discussions
- ADR generation
- Design critiques
- Refactoring suggestions
- Trade-off analysis
- Large-scale reasoning

Think of Claude Code as your **Staff Engineer**.

---

## GitHub Copilot CLI

Use for:

- Scaffolding
- Boilerplate
- Tests
- Docker
- CI/CD
- Small implementation tasks

Think of Copilot as your **Senior Developer**.

---

## AI Rule

Nothing generated by AI may be merged unless the responsible engineer can explain:

- Why it works.
- Why it was implemented this way.
- Alternatives considered.
- Trade-offs involved.

---

# AI Engineering Team

Each project uses specialist AI agents.

- Product & Requirements Analyst
- Project Playbook Editor
- Solution Architect
- Architecture Critic
- Trade-off Analyst
- Scaffolding Engineer
- Staff Engineer
- Senior Code Reviewer
- Testing Architect
- Security Reviewer
- Production Readiness Reviewer
- Scale & Resilience Engineer
- ADR Writer
- Technical Writer
- Reflection & Retrospective Coach
- Portfolio Curator
- Principal Engineer

Each agent has clearly defined responsibilities and a repeatable review process.

---

# Residency Starter Kits

Projects begin with language-specific starter kits.

These remove unnecessary setup while avoiding architectural decisions.

Included:

- GitHub Actions
- `.gitignore`
- `.editorconfig`
- Makefile / Taskfile
- Documentation folders
- ADR templates
- Architecture templates
- Retrospective templates

Not included:

- Application architecture
- Business logic
- Router
- Framework
- Database
- Package structure
- Interfaces

These remain engineering decisions for participants.

---

# Starter Kit Progression

## Projects 1–3

Rich starter kit.

Focus on engineering fundamentals.

---

## Projects 4–10

Minimal starter kit.

More ownership.

Less scaffolding.

---

## Projects 11–25

Blank repository.

The team designs everything from scratch.

---

# Project 1

## Theme

Backend Engineering Fundamentals

## Problem

Build a URL Shortener API.

## Suggested Stack

- Go
- `net/http` or Gin
- SQLite or JSON persistence
- Go testing package
- Optional Docker

## Concepts Learned

- Backend architecture
- REST API design
- Routing
- Data modelling
- Package organisation
- Persistence
- Validation
- Error handling
- Testing
- Documentation
- ADRs

The URL shortener is simply the business domain.

The true objective is learning backend engineering.

---

# Project 1 Workflow

Before writing code:

1. Read the brief.
2. Ask questions.
3. Record assumptions.
4. Produce architecture.
5. Write ADRs.
6. Present the design.
7. Receive feedback.
8. Begin implementation.

Engineering comes before coding.

---

# Deliverables

Every project should produce:

- Source code
- README
- Architecture diagrams
- ADRs
- API documentation
- Data model
- Testing strategy
- Production readiness review
- Retrospective
- Lessons learned

The portfolio should showcase engineering thinking rather than simply completed software.

---

# Long-Term Vision

This residency is designed to compress years of engineering experience into a structured learning journey.

Participants won't simply complete 25 projects.

They will develop:

- Engineering judgement
- Architectural thinking
- Communication skills
- Leadership experience
- Production engineering knowledge
- Confidence working with AI as an engineering partner

## End Goal

To think, communicate and operate like experienced software engineers—not simply developers who know another programming language.