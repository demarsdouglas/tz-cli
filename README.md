# tz-cli

A simple CLI tool for viewing timezones in real time.

![Go](https://img.shields.io/badge/Go-1.22-blue)
![Release](https://img.shields.io/github/v/release/demarsdouglas/tz-cli)

"_Time is hard_" - one of my previous managers, often

I built this tool as the current company I'm working for exists in multiple geographies and I'm admittedly horrible at quickly converting timezones in my head. Sometimes I just need to know "what time is it in Brisbane right now" and this tool solves that.

## ✨ Features

- 🌎 View the current time in any timezone, instantly
- 🔍 Customizable for the timezones you care about
- 🧼 Clean, minimal terminal output
- 🚀 Fast - built in Go

## Usage

#### Display configured timezones
```bash
$ tz
🕒 Timezone Overview:
---------------------------
Local                2025-07-01 13:21:03
America/New_York     2025-07-01 13:21:03
America/Los_Angeles  2025-07-01 10:21:03
Australia/Brisbane   2025-07-02 03:21:03
```

#### Flags
```
--help       Show help information
--version    Show version information
--live       Update the values in real time
```

#### Add a timezone
```bash
$ tz add los angeles
✅ Added timezone: America/Los_Angeles
```

#### Remove a timezone
```bash
$ tz remove brisbane
✅ Removed timezone: Australia/Brisbane
```


## 🛠️ Installation

### Homebrew (macOS/Linux)

```bash
brew tap demarsdouglas/homebrew-tap
brew install --cask tz
sudo xattr -rd com.apple.quarantine $(which tz)
```

### Manual (any OS)

Head to Releases and download the appropriate tarball for your OS and architecture.

Then:
```bash
tar -xzf tz-<OS>-<ARCH>.tar.gz
mv tz /usr/local/bin/  # or anywhere in your $PATH
```