# BWM - Bing Webmaster Tools CLI

A command-line tool for managing [Bing Webmaster Tools](https://www.bing.com/webmasters/) from the terminal. Submit URLs, monitor crawl stats, analyze traffic, research keywords, manage sitemaps, and more — all without leaving the CLI.

Built for CLI-first workflows: Claude Code, CI/CD pipelines, and shell scripting.

## Installation

### Using Go (recommended)

Requires Go 1.25 or later.

```bash
go install github.com/Softorize/bwm@latest
```

Make sure `~/go/bin` is in your PATH:

```bash
# Add to ~/.zshrc, ~/.bashrc, or ~/.zshenv
export PATH="$HOME/go/bin:$PATH"
```

### From source

```bash
git clone https://github.com/Softorize/bwm.git
cd bwm
make build
# Binary is at ./bwm
```

### Verify installation

```bash
bwm version
```

## Getting Started

### 1. Get your Bing Webmaster API key

1. Go to [Bing Webmaster Tools](https://www.bing.com/webmasters/)
2. Sign in with your Microsoft account
3. Navigate to **Settings** (gear icon) > **API access** > **API key**
4. Copy your API key

### 2. Configure BWM

```bash
# Initialize config file (~/.bwm/config.yaml)
bwm config init

# Set your API key
bwm config set api_key YOUR_API_KEY

# Set a default site (optional, avoids passing --site every time)
bwm config set site https://yoursite.com
```

### 3. Verify it works

```bash
bwm site list
```

## Authentication

BWM uses the Bing Webmaster API key for authentication. The API key is resolved in this order:

| Priority | Source | Example |
|----------|--------|---------|
| 1 | `--api-key` flag | `bwm site list --api-key YOUR_KEY` |
| 2 | `BWM_API_KEY` env var | `export BWM_API_KEY=YOUR_KEY` |
| 3 | Config file | `bwm config set api_key YOUR_KEY` |

## Configuration

Config is stored at `~/.bwm/config.yaml` with `0600` permissions.

```bash
# Initialize config
bwm config init

# Set values
bwm config set api_key YOUR_API_KEY
bwm config set site https://yoursite.com
bwm config set output json          # default: table

# Read values
bwm config get api_key
bwm config get site
bwm config get output
```

Available keys: `api_key`, `site`, `output`

## Commands

### Global Flags

All commands support these flags:

```
--api-key string   Bing Webmaster API key
--config string    Config file path (default ~/.bwm/config.yaml)
-o, --output       Output format: json | table (default: table)
-s, --site         Site URL
```

### Site Management

```bash
# List all sites in your account
bwm site list

# Add a new site
bwm site add https://yoursite.com

# Verify a site
bwm site verify https://yoursite.com

# Remove a site
bwm site remove https://yoursite.com
```

### URL Submission

```bash
# Submit a single URL for indexing
bwm url submit https://yoursite.com/new-page -s https://yoursite.com

# Submit multiple URLs
bwm url submit-batch https://yoursite.com/page1 https://yoursite.com/page2 -s https://yoursite.com

# Submit URLs from a file (one URL per line)
bwm url submit-batch --file urls.txt -s https://yoursite.com

# Check your submission quota
bwm url quota -s https://yoursite.com
```

### Crawl Statistics

```bash
# Get crawl stats (pages crawled, errors, index count)
bwm crawl stats -s https://yoursite.com

# Get crawl issues
bwm crawl issues -s https://yoursite.com

# Get crawl rate settings
bwm crawl settings -s https://yoursite.com
```

### Traffic Analytics

```bash
# Get overall traffic stats (clicks, impressions, CTR, position)
bwm traffic stats -s https://yoursite.com

# Get query-level breakdown
bwm traffic query-stats -s https://yoursite.com

# Get stats for a specific page
bwm traffic page-stats -s https://yoursite.com --page https://yoursite.com/blog
```

### Link Analysis

```bash
# Get inbound link counts
bwm link counts -s https://yoursite.com

# Get inbound link details (paginated)
bwm link detail -s https://yoursite.com
bwm link detail -s https://yoursite.com --page 1
```

### Keyword Research

```bash
# Get keyword stats for your site
bwm keyword stats -s https://yoursite.com

# Research a keyword (volume, impressions)
bwm keyword research "your keyword"

# Get related keywords
bwm keyword related "your keyword"
```

### Sitemap Management

```bash
# List all sitemaps
bwm sitemap list -s https://yoursite.com

# Submit a sitemap
bwm sitemap submit https://yoursite.com/sitemap.xml -s https://yoursite.com

# Get sitemap details (URL count, errors, warnings)
bwm sitemap detail https://yoursite.com/sitemap.xml -s https://yoursite.com

# Remove a sitemap
bwm sitemap remove https://yoursite.com/sitemap.xml -s https://yoursite.com
```

## Output Formats

BWM supports two output formats:

### Table (default)

```bash
bwm site list
```

```
URL                      Verified
─                        ─
https://yoursite.com     true
https://other.com        false
```

### JSON

```bash
bwm site list -o json
```

```json
[
  {
    "Url": "https://yoursite.com",
    "IsVerified": true,
    "AuthenticationCode": ""
  }
]
```

Set the default format globally:

```bash
bwm config set output json
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `BWM_API_KEY` | Bing Webmaster API key (overrides config file) |

## CI/CD Usage

BWM works well in CI/CD pipelines. Use the environment variable for auth and JSON output for parsing:

```bash
export BWM_API_KEY="$BING_API_KEY"

# Submit URLs after deployment
bwm url submit https://yoursite.com/new-page -s https://yoursite.com

# Batch submit from a file
bwm url submit-batch --file deployed-urls.txt -s https://yoursite.com

# Check crawl health (parse JSON in scripts)
bwm crawl issues -s https://yoursite.com -o json
```

## Shell Completion

BWM supports shell completion for Bash, Zsh, Fish, and PowerShell:

```bash
# Zsh
bwm completion zsh > "${fpath[1]}/_bwm"

# Bash
bwm completion bash > /etc/bash_completion.d/bwm

# Fish
bwm completion fish > ~/.config/fish/completions/bwm.fish
```

## Command Reference

```
bwm version                          Print version
bwm config init                      Initialize config file
bwm config set <key> <value>         Set a config value
bwm config get <key>                 Get a config value
bwm site list                        List all sites
bwm site add <url>                   Add a site
bwm site remove <url>                Remove a site
bwm site verify <url>                Verify a site
bwm url submit <url>                 Submit URL for indexing
bwm url submit-batch [urls...]       Submit multiple URLs
bwm url quota                        Get submission quota
bwm crawl stats                      Get crawl statistics
bwm crawl issues                     Get crawl issues
bwm crawl settings                   Get crawl settings
bwm traffic stats                    Get traffic statistics
bwm traffic query-stats              Get query-level stats
bwm traffic page-stats               Get page-level stats
bwm link counts                      Get link counts
bwm link detail                      Get inbound link details
bwm keyword stats                    Get keyword statistics
bwm keyword research <query>         Research a keyword
bwm keyword related <query>          Get related keywords
bwm sitemap list                     List sitemaps
bwm sitemap submit <url>             Submit a sitemap
bwm sitemap remove <url>             Remove a sitemap
bwm sitemap detail <url>             Get sitemap details
```

## License

MIT
