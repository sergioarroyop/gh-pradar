# gh-pradar

A terminal app to monitor open GitHub pull requests across multiple repositories.

## Features

- Lists open PRs from all configured repositories.
- Auto-refreshes every minute.
- Opens selected PR in browser (`Enter`).
- Optional notification sound when PRs change.
- Draft-aware notifications: if the update is only on draft PRs, no sound is played.
- Responsive flex table in terminal.

## Requirements

- Go `1.24.3` (or compatible `1.24.x`).
- A GitHub Personal Access Token (PAT) with permission to read your repositories/PRs.
- Linux environment with `xdg-open` available (used to open PR links from terminal).

## Configuration

The app reads config from:

- `~/.config/gh-pradar/config.json`

Create the file with content like:

```json
{
  "personal_access_token": "ghp_xxxxxxxxxxxxxxxxxxxx",
  "sound": "msn",
  "repository_list": [
    "owner/repo-one",
    "owner/repo-two"
  ]
}
```

### Config fields

- `personal_access_token`: your GitHub PAT.
- `sound`: optional. Leave empty (`""`) to disable sound.
- `repository_list`: list of repositories in `owner/repo` format.

### Available sounds

- `coin`
- `flashbang`
- `fart`
- `wow`
- `pew`
- `msn`

## Run

From project root:

```bash
go run .
```

## Build

Build binary:

```bash
go build -o gh-pradar .
```

Run binary:

```bash
./gh-pradar
```

## Controls

- `↑/↓`: move selection
- `Enter`: open selected PR in browser
- `q`, `esc`, `ctrl+c`: quit

## Notes

- The app clears screen and runs in terminal alternate screen mode.
- If `repository_list` is empty, startup fails with an error.

---

_This README was lovingly assembled by an AI agent intern that definitely asked for coffee breaks._
