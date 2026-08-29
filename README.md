# Sushupti

<p align="center">
  <img src="assets/demo.gif" width="900" alt="Sushupti terminal dashboard">
</p>

<p align="center">
  <b>A terminal dashboard for your Hackatime coding activity.</b>
</p>

<p align="center">
  <i>Because I hate browsers</i>
</p>

---

## What is Sushupti?

Sushupti is a **TUI (Terminal User Interface)** for Hackatime.

It pulls your coding activity from Hackatime and turns it into a dashboard that you can open directly in your terminal/cmd

Instead of opening a browser just to check your coding statistics, Sushupti has fllowing features
:

* total coding time
* active days
* daily average
* best coding day
* daily activity graphs
* project statistics
* project leaderboard
* project time distribution
* current date and time
* live status
* animated UI elements

And yes, there is also a tiny animated cat.

Dont ask me why i even added that cat...idk myself just felt cool lol

---

## Why did I make this?

The original idea was much smaller.

I wanted to look at my Hackatime activity and see how much I coded each day in a graph.

So I made a tool that prints graph in terminal

Felt useless made a tool that exports graph as PNG

Felt more useless so I made what you see now

Im a Linux user. If I'm already sitting in a terminal writing code, opening a browser just to check my coding statistics feels like an unnecessary context switch to me lwk

And somehow the small graph program became Sushupti.

---

# Features

### Daily Coding Statistics

Sushupti calculates useful statistics from your coding activity:

* **Total coding time**
* **Active days**
* **Daily average**
* **Best coding day**
* **Daily activity**
* **Project totals**

(All are in utils/stats) (TALKING OF MATH FUNCS NOT THE UI)

### Project Analytics

It also has:

* project leaderboard
* project time distribution
* most-used project
* percentage of total coding time spent on projects

The project graph also uses the same project colours shown in the leaderboard, so you can tell what you're looking at (atp i believe that i can make a bug into feature).

### Live Dashboard (this is latest thing was supposed to be added early)

It includes:

* live clock
* current date
* live status indicator
* animated graphs
* animated terminal elements
* keyboard-driven navigation

[NOTE: FOR SOME REASON ON TERMUX THE CLOCK SHOWS TIME OF THIS AFRICAN PLACE DAKAR IDK THE REASON THID IS ANDROID ONLY BUG]

### The Cat

There is a cat.

It started because there was an empty space in the interface.

Then I spent way too much time animating it.

It now moves across the sidebar.

(CAT FRAMES ARE MADE BY AI SORRY AI GODS I CANT MAKE THOSE MYSELF I TRIED TRUST Me

---

# Screenshots

The main dashboard:

<p align="center">
  <img src="assets/demo.gif" width="900" alt="Sushupti demo">
</p>

---

# Built With

Sushupti is written in **Go**.

The TUI is built using:

* **Go** — main language
* **Bubble Tea** — TUI framework
* **Lipgloss** — terminal styling
* **Go-Figure** — large terminal text
* **Hackatime API** — coding activity data

---

# Installation

You only need:

* Hackatime (REQUIRED IF YOU DONT HAVE Hackatime CLI DONT USE THIS)
* a terminal

Sushupti reads your Hackatime/WakaTime configuration from:

```text
~/.wakatime.cfg
```

Your Hackatime API configuration needs to be available there before running Sushupti.

## Linux

Run:

```bash
curl -fsSL https://raw.githubusercontent.com/Dat-one-dev/Sushupti/main/install.sh | sh
```

Then:

```bash
sushupti
```

The installer automatically detects whether you're using **amd64** or **arm64**.

## macOS

Run:

```bash
curl -fsSL https://raw.githubusercontent.com/Dat-one-dev/Sushupti/main/install.sh | sh
```

Then:

```bash
sushupti
```

## Termux

Run:

```bash
curl -fsSL https://raw.githubusercontent.com/Dat-one-dev/Sushupti/main/install.sh | sh
```

Then:

```bash
sushupti
```

> Make sure `curl` is installed in Termux first.

## Windows

Download the appropriate `.tar.gz` file from the latest GitHub release:

* **Windows amd64** — `Sushupti_windows_amd64.tar.gz`
* **Windows arm64** — `Sushupti_windows_arm64.tar.gz`

Extract the archive and run:

```text
Sushupti.exe
```

### Latest Release

[Download Sushupti from GitHub Releases](https://github.com/Dat-one-dev/Sushupti/releases/latest)

---

# What I Learned

This project started because I wanted a better graph.

It ended up teaching me much more about:

* Go
* Bubble Tea
* terminal rendering
* TUI architecture
* state-driven interfaces
* API data handling
* statistics
* terminal layouts
* UI design
* code organization

The biggest lesson was probably that **scope creep is real**.

I changed the design more times than I probably should have, but that process taught me a lot about building interfaces instead of just writing code that technically works.

---

# Why "Sushupti"?

Sushupti is a Sanskrit word associated with **deep sleep**.

(3rd quater od human consiousness acc to Madnukya Upanishad)

---

# Contributing

If you find a bug, have an idea, or want to improve something, feel free to open an issue or pull request.

Ideas for new statistics and dashboard components are especially welcome.

---

# Author

Built by **Dat-One-Dev**.

```text
┌─────────────────────────────────────────────┐
│                                             │
│                  SUSHUPTI                   │
│                                             │
│       Hackatime, but in your terminal.     │
│                                             │
└─────────────────────────────────────────────┘
```

**Source:**
https://github.com/Dat-one-dev/Sushupti
