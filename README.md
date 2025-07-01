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
🕒 Timezone Overview (relative):
────────────────────────────────────────────────────────────
🏠 Local                2025-07-01 15:34:57 (you are here)
🌍 Asia/Tokyo           2025-07-02 04:34:57 (+13h)    
🌍 Australia/Brisbane   2025-07-02 05:34:57 (+14h)    
🌍 Pacific/Auckland     2025-07-02 07:34:57 (+16h)    
🌍 America/Los_Angeles  2025-07-01 12:34:57 (-3h)     
🌍 America/Denver       2025-07-01 13:34:57 (-2h)     
🌍 America/Chicago      2025-07-01 14:34:57 (-1h)     
🟢 America/New_York     2025-07-01 15:34:57 (same)    
🌍 Europe/London        2025-07-01 20:34:57 (+5h)
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
```

If you're on OSX you may need to run the following command as well
```
sudo xattr -rd com.apple.quarantine $(which tz)
```

### Manual (any OS)

Head to Releases and download the appropriate tarball for your OS and architecture.

Then:
```bash
tar -xzf tz-<OS>-<ARCH>.tar.gz
mv tz /usr/local/bin/  # or anywhere in your $PATH
```