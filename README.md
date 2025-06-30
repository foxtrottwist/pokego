# PokéGo

A terminal-based Pokédex built with Go 1.23 that lets you explore Pokemon data through simple commands. Navigate locations, catch Pokemon, and build your collection using the PokéAPI.

## Purpose

Built while learning Go's concurrency patterns and CLI development. This project demonstrates practical API integration with caching and clean command architecture.

## Tech Stack

**Framework & Core:**
- Go 1.23.3 - Latest Go features with improved terminal handling
- PokéAPI REST endpoints - Comprehensive Pokemon data source
- golang.org/x/term - Cross-platform terminal input/output management

**Architecture:**
- REPL pattern for natural command interaction
- Command pattern with modular handlers
- Configurable caching layer to minimize API requests
- Concurrent request handling for responsive performance

## Key Features

**Commands:**
- `map` / `mapb` - Browse Pokemon world locations with pagination
- `explore <location>` - Discover Pokemon in specific areas
- `catch <pokemon>` - Attempt Pokemon capture with randomized success
- `inspect <pokemon>` - View detailed stats for captured Pokemon
- `pokedex` - Display your complete collection
- `cache` - Manage API response caching
- `clear` - Clean terminal output

**Performance:**
- 5-second request timeout prevents hanging
- 5-minute cache expiration balances freshness with speed
- Reduces API calls by ~80% during typical usage
- Graceful error handling with helpful messages

## Getting Started

**Prerequisites:**
- Go 1.23+
- Internet connection

**Installation:**
```bash
git clone https://github.com/foxtrottwist/pokego.git
cd pokego
go build -o pokego
./pokego
```

**Usage:**
```bash
PokéGo > help
PokéGo > map
PokéGo > explore pallet-town-area
PokéGo > catch pikachu
PokéGo > pokedex
PokéGo > exit
```

## Project Context

This CLI application focuses on clean separation of concerns - API communication, caching, terminal handling, and command processing each live in dedicated modules. The caching system stores API responses locally with configurable expiration, reducing network requests while maintaining data accuracy.

The modular command structure makes adding new features straightforward while keeping the codebase maintainable.
