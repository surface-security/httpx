# httpx

This scanner performs several probes on targets.

Input expected: one hostname or IP per line
Output provided: one json file per hostname/IP, with several fields (results for each probe)

## NOTE

This scanner has been purposedly made with same output as [baseline](.../baseline) so it can be used as a drop-in replacement for it.

## Run locally
- Customize `test/input/input.txt` and run `test/run.sh`
