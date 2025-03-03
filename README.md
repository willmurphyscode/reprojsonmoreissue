## Quick repro steps

Requirements: `go`, `mksquashfs`, `bash`, `wget`

1. Clone this repo
2. Run `./repro.sh`
3. See that even though the JSON file has a single document in it, the loop runs 100 times
