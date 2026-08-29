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

It pulls your coding activity from Hackatime and turns it into a dashboard that you can open directly in your terminal.

Instead of opening a browser just to check your coding statistics, Sushupti gives you things like:

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

Because an empty terminal corner looked wrong...

---

## Why did I make this?

The original idea was much smaller.

I wanted to look at my Hackatime activity and see **how much I coded each day in a graph**.

Hackatime already had the data, but I wanted something that felt natural to use from a terminal.

I'm a Linux user. If I'm already sitting in a terminal writing code, opening a browser just to check my coding statistics feels like an unnecessary context switch.

So I started building a tiny Go CLI that fetched Hackatime data and printed a daily graph.

Then I made a PNG exporter.

Then I added a proper TUI.

Then I kept adding things.

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

### Project Analytics

See where your coding time actually went.

Sushupti includes:

* project leaderboard
* project time distribution
* most-used project
* percentage of total coding time spent on projects

The project graph also uses the same project colours shown in the leaderboard, so you can tell what you're looking at without guessing.

### Live Dashboard

The dashboard isn't just static text.

It includes:

* live clock
* current date
* live status indicator
* animated graphs
* animated terminal elements
* keyboard-driven navigation

### The Cat

There is a cat.

It started because there was an empty space in the interface.

Then I spent way too much time animating it.

It now moves across the sidebar.

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

The project started as a small CLI and gradually grew into a full terminal application.

---

# Installation

Sushupti provides pre-built releases, so you **do not need Go or a compiler** to install it.

You only need:

* Hackatime
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

The installer automatically detects **Intel** and **Apple Silicon** Macs.

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

No Go installation or building from source is required.

---

# Controls

Sushupti is designed to be used entirely from the keyboard.

```text
q
    Quit

Ctrl+C
    Quit
```

---

# Project Structure

The project is split into a few parts instead of keeping everything inside one huge TUI file.

```text
Sushupti/
├── data/
├── graph/
├── styles/
├── utils/
├── tui/
├── assets/
├── main.go
└── README.md
```

The exact structure may change as development continues, but the goal is to keep the data processing, statistics, styling, and TUI logic separate.


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

The original idea was a tiny CLI.

The final idea became a full dashboard.

I changed the design more times than I probably should have, but that process taught me a lot about building interfaces instead of just writing code that technically works.

---

# What's Next?

Sushupti is close to the end of its current development cycle, but there are still things I would like to explore.

Possible future improvements include:

* more activity visualizations
* better session statistics
* additional dashboard views
* configuration options
* more keyboard controls
* improved installation/distribution

The goal isn't to turn Sushupti into a massive application.

I want it to stay a **small, fast, terminal-first way of checking Hackatime statistics.**

---

# Why "Sushupti"?

Sushupti is a Sanskrit word associated with **deep sleep**.

The name was chosen because the project is meant to stay quietly in the terminal rather than constantly demanding attention.

You open it when you want your statistics.

You check them.

Then you get back to coding.

---

# Contributing

If you find a bug, have an idea, or want to improve something, feel free to open an issue or pull request.

Ideas for new statistics and dashboard components are especially welcome.

---

# Author

Built by **Dat-One-Dev**.

I originally wanted a simple graph.

I accidentally built a TUI.

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
