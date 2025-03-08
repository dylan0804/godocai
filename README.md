# GodocAI

A terminal UI application that provides AI-powered explanations for Go packages and types.

## Overview

GodocAI combines web scraping of pkg.go.dev with local AI to help developers quickly understand Go types and packages. The application runs entirely in your terminal and processes all AI requests locally using Ollama.

## Features

- Search for Go packages and types from pkg.go.dev
- View documentation in your terminal
- Get AI-generated practical explanations
- 100% local processing (no API keys needed)

## Requirements

- Go 1.18+
- Ollama with deepseek-r1:14b model

## Quick Start

```bash
# Install Ollama and pull the model
ollama pull deepseek-r1:14b

# Start Ollama in a separate terminal
ollama serve

# Run GodocAI
go run main.go
```

## Usage

1. Type a package or type name
2. Press Enter to search
3. Navigate results with arrow keys
4. Press Enter on a result to view AI explanation

## Tech Stack

- Go programming language
- Bubble Tea terminal UI framework
- Colly web scraping library
- Ollama for local AI processing